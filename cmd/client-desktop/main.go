package main

import (
	"context"
	"fmt"
	"image"
	"log"
	"os"
	"sort"
	"time"
	"tomb_mates/internal/game"
	"tomb_mates/internal/protos"
	"tomb_mates/internal/resources"

	"github.com/coder/websocket"
	e "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	title  string
	width  int
	height int
	scale  float64
}

type Sprite struct {
	Frames []*e.Image
	Frame  int
	X      float64
	Y      float64
	Side   protos.Direction
	Config image.Config
}

type Camera struct {
	X       float64
	Y       float64
	Padding float64
}

var config *Config
var world *game.Game
var camera *Camera
var frames map[string]resources.Frames
var frame int
var currentKey e.Key
var prevKey e.Key
var levelImage *e.Image

// Game implements ebiten.Game interface.
type Game struct {
	Conn *websocket.Conn
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
var lastUpdateTime = time.Now()
var dt float64 = 0.0
var maxDt float64 = 0.0
var avgDt float64 = 0.0

func (g *Game) Update() error {
	dt = time.Now().Sub(lastUpdateTime).Seconds()
	if dt > maxDt {
		maxDt = dt
	}

	avgDt = (dt + avgDt) / 2

	world.HandlePhysics(dt)
	lastUpdateTime = time.Now()

	// Write your game's logical update.
	if world.Units[world.MyID] == nil {
		return nil
	}

	err := handleInput(g.Conn)
	if err != nil {
		return err
	}

	return nil
}

var sprites = make([]Sprite, 10000)

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *e.Image) {
	l := len(world.Units)
	if l == 0 {
		return
	}

	// Write your game's rendering.
	handleCamera(screen)
	if camera == nil {
		return
	}

	frame++

	i := 0
	world.Mx.Lock()
	for _, unit := range world.Units {
		sprites[i] = Sprite{
			Frames: frames[unit.Skin+"_"+unit.Action].Frames,
			Frame:  int(unit.Frame),
			X:      unit.Position.X,
			Y:      unit.Position.Y,
			Side:   unit.Side,
			Config: frames[unit.Skin+"_"+unit.Action].Config,
		}
		i++
	}
	world.Mx.Unlock()

	sort.Slice(sprites, func(i, j int) bool {
		depth1 := sprites[i].Y + float64(sprites[i].Config.Height)
		depth2 := sprites[j].Y + float64(sprites[j].Config.Height)
		return depth1 < depth2
	})

	for _, sprite := range sprites {
		op := &e.DrawImageOptions{}

		if sprite.Side == protos.Direction_left {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(sprite.Config.Width), 0)
		}

		op.GeoM.Translate(sprite.X-camera.X, sprite.Y-camera.Y)

		screen.DrawImage(sprite.Frames[(frame/7+sprite.Frame)%4], op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f ; FPS: %0.2f ; dt: %0.3f ; maxdt: %0.3f ; avgdt: %0.3f ; players: %d", e.ActualTPS(), e.ActualFPS(), dt, maxDt, avgDt, len(world.Units)))
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func init() {
	config = &Config{
		title:  "Another Hero",
		width:  640,
		height: 480,
	}

	var err error
	frames, err = resources.Load()
	if err != nil {
		log.Fatal(err)
	}

	levelImage, err = prepareLevelImage()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	world = game.New(true, map[string]*protos.Unit{})

	url := getEnv("HOST", "ws://localhost:3000/ws")

	ws, _, err := websocket.Dial(context.TODO(), url, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func(c *websocket.Conn) {
		defer c.CloseNow()

		for {
			var _, message, err = c.Read(context.TODO())
			if err != nil {
				log.Fatal("Error reading message:", err)
			}

			event := &protos.Event{}
			err = proto.Unmarshal(message, event)
			if err != nil {
				log.Fatal(err)
			}

			world.HandleEvent(event)

			if event.Type == protos.EventType_connect {
				me := world.Units[world.MyID]
				camera = &Camera{
					X:       me.Position.X,
					Y:       me.Position.Y,
					Padding: 30,
				}
			}
		}
	}(ws)

	e.SetRunnableOnUnfocused(true)
	e.SetWindowSize(config.width, config.height)
	e.SetWindowTitle(config.title)
	game := &Game{Conn: ws}
	if err := e.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func prepareLevelImage() (*e.Image, error) {
	tileSize := 16
	level := resources.LoadLevel()
	width := len(level[0])
	height := len(level)
	levelImage := e.NewImage(width*tileSize, height*tileSize)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			op := &e.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize), float64(j*tileSize))

			img := e.NewImageFromImage(frames[level[j][i]].Frames[0])
			levelImage.DrawImage(img, op)
		}
	}

	return levelImage, nil
}

func handleCamera(screen *e.Image) {
	if camera == nil {
		return
	}

	player := world.Units[world.MyID]
	frame := frames[player.Skin+"_"+player.Action]
	camera.X = player.Position.X - float64(config.width-frame.Config.Width)/2
	camera.Y = player.Position.Y - float64(config.height-frame.Config.Height)/2

	op := &e.DrawImageOptions{}
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.GeoM.Scale(1, 1)

	screen.DrawImage(levelImage, op)
}

func handleInput(c *websocket.Conn) error {
	event := &protos.Event{}

	if e.IsKeyPressed(e.KeyA) || e.IsKeyPressed(e.KeyLeft) {
		event = &protos.Event{
			Type: protos.EventType_move,
			Data: &protos.Event_Move{
				Move: &protos.EventMove{
					PlayerId:  world.MyID,
					Direction: protos.Direction_left,
				},
			},
		}
		if currentKey != e.KeyA {
			currentKey = e.KeyA
		}
	}

	if e.IsKeyPressed(e.KeyD) || e.IsKeyPressed(e.KeyRight) {
		event = &protos.Event{
			Type: protos.EventType_move,
			Data: &protos.Event_Move{
				Move: &protos.EventMove{
					PlayerId:  world.MyID,
					Direction: protos.Direction_right,
				},
			},
		}
		if currentKey != e.KeyD {
			currentKey = e.KeyD
		}
	}

	if e.IsKeyPressed(e.KeyW) || e.IsKeyPressed(e.KeyUp) {
		event = &protos.Event{
			Type: protos.EventType_move,
			Data: &protos.Event_Move{
				Move: &protos.EventMove{
					PlayerId:  world.MyID,
					Direction: protos.Direction_up,
				},
			},
		}
		if currentKey != e.KeyW {
			currentKey = e.KeyW
		}
	}

	if e.IsKeyPressed(e.KeyS) || e.IsKeyPressed(e.KeyDown) {
		event = &protos.Event{
			Type: protos.EventType_move,
			Data: &protos.Event_Move{
				Move: &protos.EventMove{
					PlayerId:  world.MyID,
					Direction: protos.Direction_down,
				},
			},
		}
		if currentKey != e.KeyS {
			currentKey = e.KeyS
		}
	}

	unit := world.Units[world.MyID]

	if event.Type == protos.EventType_move {
		if prevKey != currentKey {
			message, err := proto.Marshal(event)
			if err != nil {
				return err
			}

			err = c.Write(context.Background(), websocket.MessageBinary, message)
			if err != nil {
				return err
			}
		}
	} else {
		if unit.Action != game.UnitActionIdle {
			event = &protos.Event{
				Type: protos.EventType_idle,
				Data: &protos.Event_Idle{
					Idle: &protos.EventIdle{PlayerId: world.MyID},
				},
			}

			message, err := proto.Marshal(event)
			if err != nil {
				return err
			}
			err = c.Write(context.Background(), websocket.MessageBinary, message)
			if err != nil {
				// ...
				return err
			}
			currentKey = -1
		}
	}

	prevKey = currentKey

	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	fmt.Println("Env not found: ", key, " - Using fallback: ", fallback)
	return fallback
}