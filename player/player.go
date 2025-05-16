package player

import (
	"FirstHero/primShapes"

	"github.com/go-gl/mathgl/mgl32"
)
type Limb struct {
	Name string
	Parent *Limb
	TargetPos [4]mgl32.Vec2
	CurrentPos [4]mgl32.Vec2
	Color []mgl32.Vec4
}
func (l *Limb) CreateLimb() ([]float32, []uint32) {
	var vertices []float32
	var indices []uint32

	if l.Name != "Head"{
		q := primShapes.Quad{Pos: l.CurrentPos, Color: l.Color}
		vertices, indices = q.CreateQuad()
	} else {
		rq := primShapes.RoundedQuad{Pos: l.CurrentPos, Color: l.Color, 
			Radius: 0.1, Segments: 6}
		vertices, indices = rq.CreateRoundedQuad()
	}
	return vertices, indices
}

type Player struct {
	RootLimb *Limb
	Limbs map[string]*Limb 

	Alpha float32
	Speed float32
	JumpHeight float32
}
func (p *Player) SetTarget (axis int, delta float32) {
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
      l.CurrentPos[i] = l.CurrentPos[i].Mul(1-p.Alpha).Add(l.TargetPos[i].Mul(p.Alpha))
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
