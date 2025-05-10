package primShapes

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)
type Triangular struct {
	StartPos mgl32.Vec2
	Color mgl32.Vec4
	Width float32
	Height float32
}
func (t *Triangular) CreateTriangular(program uint32) ([]float32, int32){
	colorLoc := gl.GetUniformLocation(program, gl.Str("uColor\x00"))
  gl.Uniform4f(colorLoc, t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W())
	return []float32 {
		t.StartPos.X(), t.StartPos.Y(), 0.0, 
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
		
		t.StartPos.X()+t.Width/2, t.StartPos.Y()+t.Height, 0.0, 
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
		
		t.StartPos.X()+t.Width, t.StartPos.Y(), 0.0, 
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
	}, 3
}
