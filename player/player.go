package player

import (
	"FirstHero/primShapes"

	"github.com/go-gl/mathgl/mgl32"
)

type Limb struct {
	Name       string
	Parent     *Limb
	TargetPos  [4]mgl32.Vec2
	CurrentPos [4]mgl32.Vec2
	Color      []mgl32.Vec4
}

func (l *Limb) CreateLimb() ([]float32, []uint32) {
	q := primShapes.Quad{Pos: l.CurrentPos, Color: l.Color}
	return q.CreateQuad()
}

type Player struct {
	RootLimb *Limb
	Limbs    map[string]*Limb

	Alpha      float32
	Speed      float32
	JumpHeight float32
}

func NewPlayer() *Player {
	rt := Limb{Name: "Root", Parent: nil,
		CurrentPos: [4]mgl32.Vec2{
			{-0.8, -0.4}, {-0.8, -0.55},
			{-0.73, -0.55}, {-0.73, -0.4},
		},
		TargetPos: [4]mgl32.Vec2{
			{-0.8, -0.4}, {-0.8, -0.55},
			{-0.73, -0.55}, {-0.73, -0.4},
		},
		Color: []mgl32.Vec4{
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
		},
	}
	rh := Limb{Name: "RightHand", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{-0.05, 0.0}, {-0.05, 0.0},
			{-0.07, 0.0}, {-0.07, 0.0},
		},
		TargetPos: [4]mgl32.Vec2{
			{-0.05, 0.0}, {-0.05, 0.0},
			{-0.07, 0.0}, {-0.07, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
	lh := Limb{Name: "LeftHand", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{0.07, 0.0}, {0.07, 0.0},
			{0.05, 0.0}, {0.05, 0.0},
		},
		TargetPos: [4]mgl32.Vec2{
			{0.07, 0.0}, {0.07, 0.0},
			{0.05, 0.0}, {0.05, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
	rl := Limb{Name: "RightLeg", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{0.0, -0.15}, {0.0, -0.15},
			{-0.035, -0.15}, {-0.035, -0.15},
		},
		TargetPos: [4]mgl32.Vec2{
			{0.0, -0.15}, {0.0, -0.15},
			{-0.035, -0.15}, {-0.035, -0.15},
		},
		Color: []mgl32.Vec4{
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
		},
	}
	ll := Limb{Name: "LeftLeg", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{0.025, -0.15}, {0.025, -0.15},
			{0.0, -0.15}, {0.0, -0.15},
		},
		TargetPos: [4]mgl32.Vec2{
			{0.025, -0.15}, {0.025, -0.15},
			{0.0, -0.15}, {0.0, -0.15},
		},
		Color: []mgl32.Vec4{
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
		},
	}
	head := Limb{Name: "Head", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{0.0, 0.15}, {0.0, 0.0},
			{0.0, 0.0}, {0.05, 0.15},
		},
		TargetPos: [4]mgl32.Vec2{
			{0.0, 0.1}, {0.0, 0.0},
			{0.002, 0.0}, {0.002, 0.1},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
  allLimbs := make(map[string]*Limb)
  allLimbs["Head"] = &head
  allLimbs["LeftHand"] = &lh
  allLimbs["RightHand"] = &rh
  allLimbs["LeftLeg"] = &ll
  allLimbs["RightLeg"] = &rl
  
	pl := Player{Alpha: 0.1, Speed: 0.2, JumpHeight: 0.3,
    RootLimb: &rt, Limbs: allLimbs}
	return &pl
}
func (p *Player) SetTarget(axis int, delta float32) {
	for i := 0; i < 4; i++ {
		newPos := p.RootLimb.TargetPos[i][axis] + delta

		if newPos < -1.0 {
			delta = -1.0 - p.RootLimb.TargetPos[i][axis]
		}
		if axis == 0 && newPos > 0.9 {
			delta = 0.9 - p.RootLimb.TargetPos[i][axis]
		}
		if axis == 1 && newPos > 0.7 {
			delta = 0.7 - p.RootLimb.TargetPos[i][axis]
		}
	}

	for i := 0; i < 4; i++ {
		p.RootLimb.TargetPos[i][axis] += delta
	}
}
func (p *Player) UpdatePos(l *Limb) {
	if l.Parent == nil {
		for i := range l.TargetPos {
			l.CurrentPos[i] = l.CurrentPos[i].Mul(1 - p.Alpha).Add(l.TargetPos[i].Mul(p.Alpha))
		}
	} else {
		for i := range l.TargetPos {
			l.CurrentPos[i] = l.Parent.CurrentPos[i].Add(l.TargetPos[i])
		}
	}

	for _, child := range p.Limbs {
		if child.Parent == l {
			p.UpdatePos(child)
		}
	}
}
func (p *Player) GetAllLimbs() []*Limb {
	allLimbs := make([]*Limb, 0)
	allLimbs = append(allLimbs, p.RootLimb)

	for _, limb := range p.Limbs {
		allLimbs = append(allLimbs, limb)
	}

	return allLimbs
}
