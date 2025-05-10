package primShapes

import "github.com/go-gl/mathgl/mgl32"

type Triangular struct {
	StartPos mgl32.Vec2
	Width float32
	Height float32
}
func (t *Triangular) CreateTriangular() ([]float32, int32){
	return []float32 {
		t.StartPos.X(), t.StartPos.Y(), 0.0,
		t.StartPos.X()+t.Width/2, t.StartPos.Y()+t.Height, 0.0,
		t.StartPos.X()+t.Width, t.StartPos.Y(), 0.0,
	}, 3
}
