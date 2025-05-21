package primShapes

import (
	"github.com/go-gl/gl/v4.1-core/gl"
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
	q := Quad{Pos: l.CurrentPos, Color: l.Color}
	vertices, indices := q.CreateQuad()
	return vertices, indices
}

type Triangular struct {
	TargetPos  mgl32.Vec2
	CurrentPos mgl32.Vec2
	Color      mgl32.Vec4
	Width      float32
	Height     float32
	Alpha      float32
	Speed      float32
	JumpHeight float32
}

func (t *Triangular) CreateTriangular(program uint32) ([]float32, int32) {
	colorLoc := gl.GetUniformLocation(program, gl.Str("uColor\x00"))
	gl.Uniform4f(colorLoc, t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W())
	return []float32{
		t.CurrentPos.X(), t.CurrentPos.Y(), 0.0,
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),

		t.CurrentPos.X() + t.Width/2, t.CurrentPos.Y() + t.Height, 0.0,
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),

		t.CurrentPos.X() + t.Width, t.CurrentPos.Y(), 0.0,
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
	}, 3
}

type Quad struct {
	Name string
	Pos   [4]mgl32.Vec2
	Color []mgl32.Vec4
}

func (q *Quad) CreateQuad() ([]float32, []uint32) {
	var vertices []float32
	for i := 0; i < 4; i++ {
		vertices = append(vertices,
			q.Pos[i].X(), q.Pos[i].Y(), 0.0,
			q.Color[i].X(), q.Color[i].Y(), q.Color[i].Z(), q.Color[i].W())
	}
	indices := []uint32{0, 1, 2, 2, 3, 0}
	return vertices, indices
}

