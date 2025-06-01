package anim

import (
	"github.com/go-gl/mathgl/mgl32"

	"FirstHero/primShapes"
)

type Animator struct {
	animProgress float32
	animSpeed float32
}

func NewAnimator() *Animator {
	return &Animator{
		animSpeed: 0.008,
	}
}

func (a *Animator) StartJump(limbs []*primShapes.Limb) {
	if a.animProgress < 1.0 {
		a.animProgress += a.animSpeed
		if a.animProgress > 1.0 {
			a.animProgress = 1.0
		}
	}

	for _, limb := range limbs {
		if limb.Parent != nil && limb.Name != "Head" {
			a.animateLeg(limb)
		}
	}
}

func (a *Animator) StopJump(limbs []*primShapes.Limb) {
	if a.animProgress > 0.0 {
		a.animProgress -= a.animSpeed
	}

	for _, limb := range limbs {
		if limb.Parent != nil && limb.Name != "Head" {
			a.animateLeg(limb)
		} 
	}
}

func (a *Animator) animateLeg(limb *primShapes.Limb) {
	var yOffset float32
	if a.animProgress <= 0.3 {
		// Фаза отталкивания (0 → +0.1)
		yOffset = 0.01 * (a.animProgress / 0.3)
	} else if a.animProgress <= 0.7 {
		// Фаза полета (+0.1 → -0.05)
		yOffset = 0.01 - 0.015*((a.animProgress-0.3)/0.4)
	} else {
		// Фаза приземления (-0.05 → -0.15)
		yOffset = -0.005 - 0.01*((a.animProgress-0.7)/0.3)
	}
	for i := range limb.CurrentPos {
		basePos := limb.Parent.CurrentPos[i].Add(limb.TargetPos[i])
		limb.CurrentPos[i] = mgl32.Vec3{
			basePos.X(), basePos.Y() + yOffset, basePos.Z(),
		}
	}
}

func (a *Animator) rotateLeg(limb *primShapes.Limb, angle float32) {
	rotationCenter := limb.CurrentPos[3]
	rotMat := mgl32.HomogRotate3DZ(angle)

	for i := range limb.CurrentPos {
		point := limb.CurrentPos[i].Sub(rotationCenter)
		rotated := rotMat.Mul4x1(mgl32.Vec4{point.X(), point.Y(), point.Z(), 1})
		limb.CurrentPos[i] = mgl32.Vec3{rotated.X(), rotated.Y(), rotated.Z()}.Add(rotationCenter)
	}
} 
