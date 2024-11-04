/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package components

import (
	"gomp_game/pkgs/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

type RenderData struct {
	Sprite *ebiten.Image
}

var RenderComponent = engine.CreateComponent[RenderData]
