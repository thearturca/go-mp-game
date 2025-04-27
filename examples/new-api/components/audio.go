/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

===-===-===-===-===-===-===-===-===-===
Donations during this file development:
-===-===-===-===-===-===-===-===-===-===

none :)

Thank you for your support!
*/

package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"gomp/examples/new-api/assets"
	"gomp/pkg/ecs"
)

type SoundEffect struct {
	Clip      *rl.Sound
	IsPlaying bool
	IsLooping bool
	// base is 1.0
	Volume float32
	// base is 1.0
	Pitch float32
	// base is 0.5. 1.0 is left, 0.0 is right
	Pan float32
}

type SoundEffectBuilder struct {
	soundEffect SoundEffect
}

// Creates a new sound effect builder. Use `Build()` to get the sound effect
//
// filename should be relative to assets folder.
// Example: "sound.wav"
func NewSoundEffect(
	filename string,
) *SoundEffectBuilder {
	return &SoundEffectBuilder{
		soundEffect: SoundEffect{
			Clip:      assets.Audio.Get(filename),
			IsPlaying: false,
			IsLooping: false,
			Pitch:     1.0,
			Volume:    1.0,
			Pan:       0.5,
		},
	}
}

// should sound be looped. Default is false
func (b *SoundEffectBuilder) WithLoop(loop bool) *SoundEffectBuilder {
	b.soundEffect.IsLooping = loop

	return b
}

// base is 1.0
func (b *SoundEffectBuilder) WithPitch(pitch float32) *SoundEffectBuilder {
	b.soundEffect.Pitch = pitch

	return b
}

// base is 1.0
func (b *SoundEffectBuilder) WithVolume(volume float32) *SoundEffectBuilder {
	b.soundEffect.Volume = volume

	return b
}

// base is 0.5. 1.0 is left, 0.0 is right
func (b *SoundEffectBuilder) WithPan(pan float32) *SoundEffectBuilder {
	b.soundEffect.Pan = pan

	return b
}

func (b *SoundEffectBuilder) Build() SoundEffect {
	return b.soundEffect
}

type SoundEffectsComponentManager = ecs.ComponentManager[SoundEffect]

func NewSoundEffectsComponentManager() SoundEffectsComponentManager {
	return ecs.NewComponentManager[SoundEffect](SoundEffectManagerComponentId)
}

type SpatialAudio struct {
	// follows raylib rules. Base is 1.0
	Volume float32
	// follows raylib rules. Base is 0.5, 1.0 is left, 0.0 is right
	Pan float32
}

type SpatialAudioComponentManager = ecs.ComponentManager[SpatialAudio]

func NewSpatialAudioComponentManager() SpatialAudioComponentManager {
	return ecs.NewComponentManager[SpatialAudio](SpatialAudioManagerComponentId)
}
