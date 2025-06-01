package primShapes

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Limb struct {
	Name string
	Parent *Limb
	TargetPos [4]mgl32.Vec3
	CurrentPos [4]mgl32.Vec3
	Color []mgl32.Vec4
}
func (l *Limb) CreateLimb() ([]float32, []uint32) {
	q := Quad{Pos: l.CurrentPos, Color: l.Color}
	vertices, indices := q.CreateQuad()
	return vertices, indices
}

type Quad struct {
	Name string
	Pos   [4]mgl32.Vec3
	Color []mgl32.Vec4
}

func (q *Quad) CreateQuad() ([]float32, []uint32) {
	var vertices []float32
	for i := 0; i < 4; i++ {
		vertices = append(vertices,
			q.Pos[i].X(), q.Pos[i].Y(), q.Pos[i].Z(),
			q.Color[i].X(), q.Color[i].Y(), q.Color[i].Z(), q.Color[i].W())
	}
	indices := []uint32{0, 1, 2, 2, 3, 0}
	return vertices, indices
}

