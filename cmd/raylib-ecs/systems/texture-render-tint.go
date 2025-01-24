/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package systems

import (
	"gomp_game/cmd/raylib-ecs/components"
	"gomp_game/pkgs/gomp/ecs"
)

// TextureRenderScale is a system that sets Scale of textureRender
type trTintController struct{}

func (s *trTintController) Init(world *ecs.World)        {}
func (s *trTintController) FixedUpdate(world *ecs.World) {}
func (s *trTintController) Update(world *ecs.World) {
	// Get component managers
	tints := components.TintService.GetManager(world)
	textureRenders := components.TextureRenderService.GetManager(world)

	// Update sprites and spriteRenders
	textureRenders.AllParallel(func(entity ecs.Entity, tr *components.TextureRender) bool {
		if tr == nil {
			return true
		}

		tint := tints.Get(entity)
		if tint == nil {
			return true
		}

		trTint := &tr.Tint
		trTint.A = tint.A
		trTint.R = tint.R
		trTint.G = tint.G
		trTint.B = tint.B

		return true
	})
}
func (s *trTintController) Destroy(world *ecs.World) {}
