package player

import (
	"github.com/go-gl/mathgl/mgl32"
	
	"FirstHero/primShapes"
	"FirstHero/collision"
	"FirstHero/anim"
)

type Player struct {
	RootLimb *primShapes.Limb
	Limbs map[string]*primShapes.Limb

	Alpha      float32
	Speed      float32
	JumpHeight float32

	CanJump bool
	CanMoveLeft bool
	CanMoveRight bool
}
func NewPlayer() *Player {
	rt := primShapes.Limb{Name: "Root", Parent: nil,
		CurrentPos: [4]mgl32.Vec3{
			{-0.8, -0.4, 0.0}, {-0.8, -0.55, 0.0},
			{-0.73, -0.55, 0.0}, {-0.73, -0.4, 0.0},
		},
		TargetPos: [4]mgl32.Vec3{
			{-0.8, -0.4, 0.0}, {-0.8, -0.55, 0.0},
			{-0.73, -0.55, 0.0}, {-0.73, -0.4, 0.0},
		},
		Color: []mgl32.Vec4{
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
		},
	}
	rh := primShapes.Limb{Name: "RightHand", Parent: &rt,
		CurrentPos: [4]mgl32.Vec3{
			{-0.05, 0.0, 0.0}, {-0.05, 0.0, 0.0},
			{-0.07, 0.0, 0.0}, {-0.07, 0.0, 0.0},
		},
		TargetPos: [4]mgl32.Vec3{
			{-0.05, 0.0, 0.0}, {-0.05, 0.0, 0.0},
			{-0.07, 0.0, 0.0}, {-0.07, 0.0, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
	lh := primShapes.Limb{Name: "LeftHand", Parent: &rt,
		CurrentPos: [4]mgl32.Vec3{
			{0.07, 0.0, 0.0}, {0.07, 0.0, 0.0},
			{0.05, 0.0, 0.0}, {0.05, 0.0, 0.0},
		},
		TargetPos: [4]mgl32.Vec3{
			{0.07, 0.0, 0.0}, {0.07, 0.0, 0.0},
			{0.05, 0.0, 0.0}, {0.05, 0.0, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
	rl := primShapes.Limb{Name: "RightLeg", Parent: &rt,
		CurrentPos: [4]mgl32.Vec3{
			{0.0, -0.15, 0.0}, {0.0, -0.15, 0.0},
			{-0.035, -0.15, 0.0}, {-0.035, -0.15, 0.0},
		},
		TargetPos: [4]mgl32.Vec3{
			{0.0, -0.15, 0.0}, {0.0, -0.15, 0.0},
			{-0.035, -0.15, 0.0}, {-0.035, -0.15, 0.0},
		},
		Color: []mgl32.Vec4{
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
		},
	}
	ll := primShapes.Limb{Name: "LeftLeg", Parent: &rt,
		CurrentPos: [4]mgl32.Vec3{
			{0.025, -0.15, 0.0}, {0.025, -0.15, 0.0},
			{0.0, -0.15, 0.0}, {0.0, -0.15, 0.0},
		},
		TargetPos: [4]mgl32.Vec3{
			{0.025, -0.15, 0.0}, {0.025, -0.15, 0.0},
			{0.0, -0.15, 0.0}, {0.0, -0.15, 0.0},
		},
		Color: []mgl32.Vec4{
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
		},
	}
	head := primShapes.Limb{Name: "Head", Parent: &rt,
		CurrentPos: [4]mgl32.Vec3{
			{0.0, 0.15, 0.0}, {0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0}, {0.05, 0.15, 0.0},
		},
		TargetPos: [4]mgl32.Vec3{
			{0.0, 0.1, 0.0}, {0.0, 0.0, 0.0},
			{0.002, 0.0, 0.0}, {0.002, 0.1, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
  allLimbs := make(map[string]*primShapes.Limb)
  allLimbs["Head"] = &head
  allLimbs["LeftHand"] = &lh
  allLimbs["RightHand"] = &rh
  allLimbs["LeftLeg"] = &ll
  allLimbs["RightLeg"] = &rl
  
	pl := Player{Alpha: 0.1, Speed: 0.01, JumpHeight: 0.5, CanJump: true, CanMoveLeft: true, CanMoveRight: true,
    RootLimb: &rt, Limbs: allLimbs}
	return &pl
}
func (p *Player) SetTarget(axis int, delta float32) {
	for i := 0; i < 4; i++ {
		newPos := p.RootLimb.TargetPos[i][axis] + delta

		if newPos < -0.95 {
			delta = -0.95 - p.RootLimb.TargetPos[i][axis]
		}
		if axis == 0 && newPos > 0.95 {
			delta = 0.95 - p.RootLimb.TargetPos[i][axis]
		}
		if axis == 1 && newPos > 0.9 {
			delta = 0.9 - p.RootLimb.TargetPos[i][axis]
		}
	}

	for i := 0; i < 4; i++ {
		p.RootLimb.TargetPos[i][axis] += delta
	}
}
func (p *Player) UpdatePos(l *primShapes.Limb, gd []*primShapes.Quad, anim8 *anim.Animator) {
	if l.Parent == nil {
		for i := range l.TargetPos {
			l.CurrentPos[i] = l.CurrentPos[i].Mul(1 - p.Alpha).Add(l.TargetPos[i].Mul(p.Alpha))
		}
	} else {
		for i := range l.TargetPos {
			l.CurrentPos[i] = l.Parent.CurrentPos[i].Add(l.TargetPos[i])
		}
	}

	p.startAnim(anim8)

	for _, child := range p.Limbs {
		if child.Parent == l {
			p.UpdatePos(child, gd, anim8)
		}
	}

	if l.Name == "RightLeg" || l.Name == "LeftLeg" {
		collision.IsGrounded(l, gd, &p.CanJump)
	}

	if l.Parent == nil {
		p.CanMoveLeft = true
		p.CanMoveRight = true
		for _, limb := range p.GetAllLimbs() {
			collision.CheckWallCollision(limb, gd, &p.CanMoveLeft, &p.CanMoveRight)
			if !p.CanMoveLeft || !p.CanMoveRight {
				break
			}
		}
	}
}
func (p *Player) startAnim(anim8 *anim.Animator) {
	if !p.CanJump {
		anim8.StartJump(p.GetAllLimbs())
	} else {
		anim8.StopJump(p.GetAllLimbs())
	}
}
func (p *Player) GetAllLimbs() []*primShapes.Limb {
	allLimbs := make([]*primShapes.Limb, 0)
	allLimbs = append(allLimbs, p.RootLimb)

	for _, limb := range p.Limbs {
		allLimbs = append(allLimbs, limb)
	}

	return allLimbs 
}
