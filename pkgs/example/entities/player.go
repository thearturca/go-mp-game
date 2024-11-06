/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package entities

import (
	"gomp_game/pkgs/gomp"
)

var Player = gomp.CreateEntity(
	gomp.PhysicsComponent,
	gomp.RenderComponent,
)