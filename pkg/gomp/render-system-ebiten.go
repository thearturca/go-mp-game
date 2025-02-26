/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package gomp

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var EbitenRenderSystem = CreateSystem(new(ebitenRenderSystemController))

// ebitenRenderSystemController is a system that updates the physics of a game
type ebitenRenderSystemController struct {
	world donburi.World
}

func (c *ebitenRenderSystemController) Init(game *Game) {
	c.world = game.World
}

func (c *ebitenRenderSystemController) Update(dt float64) {
	SpriteComponent.Query.Each(c.world, func(e *donburi.Entry) {
		sprite := SpriteComponent.Query.Get(e)
		if !e.HasComponent(RenderComponent.Query) {
			donburi.Add(e, RenderComponent.Query, &RenderData{
				Sprite: ebiten.NewImageFromImage(sprite.Image),
			})
		}
	})
}
