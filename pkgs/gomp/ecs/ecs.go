/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package ecs

type ECSID uint

type ECS struct {
	ID       ECSID
	Title    string
	Entities SparseSet[Entity, EntityID]

	nextEntityID    EntityID
	nextComponentID ComponentID
	entity          Entity
}

type AnyComponentPtr interface {
	register(ecs *ECS)
}

var nextId ECSID = 0

func generateECSID() ECSID {
	id := nextId
	nextId++
	return id
}

func New(title string) ECS {
	ecs := ECS{
		ID:       generateECSID(),
		Title:    title,
		Entities: NewSparseSet[Entity, EntityID](10000000),

		nextEntityID:    0,
		nextComponentID: 0,
	}

	return ecs
}

func (e *ECS) generateComponentID() ComponentID {
	id := e.nextComponentID
	e.nextComponentID++
	return id
}

func (e *ECS) generateEntityID() EntityID {
	id := e.nextEntityID
	e.nextEntityID++
	return id
}

func (e *ECS) RegisterComponents(component_ptr ...AnyComponentPtr) {
	for i := 0; i < len(component_ptr); i++ {
		component_ptr[i].register(e)
	}
}

func (e *ECS) CreateEntity(title string) *Entity {
	e.entity.ID = e.generateEntityID()
	e.entity.Title = title
	e.entity.ComponentsMask = NewBitArray(64)
	e.entity.ecs = e

	return e.Entities.Set(e.entity.ID, e.entity)
}
