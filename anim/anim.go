package anim

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"

	"FirstHero/primShapes"
)

const (
	AnimDeath = iota
	AnimJump
	AnimWalk
	AnimIdle
)

type Animation struct {
	animType int
	progress float32
	speed float32
	loop bool
	active bool
	fadingOut bool
}

type Animator struct {
	animations map[int]*Animation
	currentAnim *Animation
	limbOffsets map[string]mgl32.Vec3
}

func NewAnimator() *Animator {
	a := &Animator{
		animations: make(map[int]*Animation),
		limbOffsets: make(map[string]mgl32.Vec3),
	}
	a.animations[AnimDeath] = &Animation{AnimDeath, 0, 0.008, false, false, false}
	a.animations[AnimJump] = &Animation{AnimJump, 0, 0.008, false, false, false}
	return a
}

func (a *Animator) Update(limbs []*primShapes.Limb) {
	a.currentAnim = nil
	for i := AnimDeath; i <= AnimIdle; i++ {
		if anim, exists := a.animations[i]; exists && (anim.active || anim.fadingOut) {
			a.currentAnim = anim
			break
		}
	}
	if a.currentAnim == nil {
		return
	}

	if a.currentAnim.active {
		if a.currentAnim.progress < 1.0 {
			a.currentAnim.progress += a.currentAnim.speed
			if a.currentAnim.progress > 1.0 {
				if a.currentAnim.loop {
					a.currentAnim.progress = 0
				} else {
					a.currentAnim.progress = 1.0
					a.currentAnim.active = false
				}
			}
		}
	} else if a.currentAnim.fadingOut {
		if a.currentAnim.progress > 0.0 {
			a.currentAnim.progress -= a.currentAnim.speed

		}
	}

	switch a.currentAnim.animType {
	case AnimJump:
		for _, limb := range limbs{
			if limb.Parent != nil && limb.Name != "Head" {
				a.animateJump(limb)
			}
		}
	}
}

func (a *Animator) StartJump(limbs []*primShapes.Limb) {
	a.animations[AnimJump].active = true
	a.animations[AnimJump].fadingOut = false
}

func (a *Animator) StopJump(limbs []*primShapes.Limb) {
	a.animations[AnimJump].active = false 
	a.animations[AnimJump].fadingOut = true
	a.animations[AnimJump].speed = 0.03
}

func (a *Animator) animateJump(limb *primShapes.Limb) {
	var yOffset float32
	a.currentAnim.speed = 0.008
	progress := a.currentAnim.progress
	if progress <= 0.3 {
		// Repulsion phase (0 -> +0.1)
		yOffset = 0.01 * (progress / 0.3)
	} else if progress <= 0.7 {
		// Fly phase (+0.1 -> -0.05)
		yOffset = 0.01 - 0.015*((progress-0.3)/0.4)
	} else {
		// Landing phase (-0.05 -> -0.15)
		yOffset = -0.005 - 0.01*((progress-0.7)/0.3)
	}
	for i := range limb.CurrentPos {
		basePos := limb.Parent.CurrentPos[i].Add(limb.TargetPos[i])
		limb.CurrentPos[i] = mgl32.Vec3{
			basePos.X(), basePos.Y() + yOffset, basePos.Z(),
		}
	}

	a.currentAnim.speed = 0.01
	if limb.Name == "LeftHand" || limb.Name == "RightHand" {
		angle := float32(0)
		if progress > 0.2 {
			targetAngle := float32(math.Pi/12)

			if limb.Name == "LeftHand" {
				angle = targetAngle * progress
			} else {
				angle = -targetAngle * progress
			}
		}
		rotationCenter := limb.CurrentPos[0]
		rotMat := mgl32.HomogRotate3DZ(angle)

		for i := range limb.CurrentPos {
			point := limb.CurrentPos[i].Sub(rotationCenter)
			rotated := rotMat.Mul4x1(mgl32.Vec4{point.X(), point.Y(), point.Z(), 1})
			limb.CurrentPos[i] = mgl32.Vec3{rotated.X(), rotated.Y(), rotated.Z()}.Add(rotationCenter)
		}
	}
}
