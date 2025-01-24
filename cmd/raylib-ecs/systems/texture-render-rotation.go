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

// TextureRenderRotation is a system that sets Rotation of textureRender
type trRotationController struct{}

func (s *trRotationController) Init(world *ecs.World)        {}
func (s *trRotationController) FixedUpdate(world *ecs.World) {}
func (s *trRotationController) Update(world *ecs.World) {
	// Get component managers
	rotations := components.RotationService.GetManager(world)
	textureRenders := components.TextureRenderService.GetManager(world)

	// Update sprites and spriteRenders
	textureRenders.AllParallel(func(entity ecs.Entity, tr *components.TextureRender) bool {
		if tr == nil {
			return true
		}

		rotation := rotations.Get(entity)
		if rotation == nil {
			return true
		}

		tr.Rotation = rotation.Angle

		return true
	})
}
func (s *trRotationController) Destroy(world *ecs.World) {}
