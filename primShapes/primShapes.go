package primShapes

import (
	"math"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)
type Triangular struct {
	TargetPos mgl32.Vec2
	CurrentPos mgl32.Vec2
	Color mgl32.Vec4
	Width float32
	Height float32
	Alpha float32
	Speed float32
	JumpHeight float32
}
func (t *Triangular) CreateTriangular(program uint32) ([]float32, int32){
	colorLoc := gl.GetUniformLocation(program, gl.Str("uColor\x00"))
  gl.Uniform4f(colorLoc, t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W())
	return []float32 {
		t.CurrentPos.X(), t.CurrentPos.Y(), 0.0, 
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
		
		t.CurrentPos.X()+t.Width/2, t.CurrentPos.Y()+t.Height, 0.0, 
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
		
		t.CurrentPos.X()+t.Width, t.CurrentPos.Y(), 0.0, 
		t.Color.X(), t.Color.Y(), t.Color.Z(), t.Color.W(),
	}, 3
}

type Quad struct {
	Pos [4]mgl32.Vec2
	Color []mgl32.Vec4
}
func (q *Quad) CreateQuad() ([]float32, []uint32){
	var vertices []float32
	for i := 0; i < 4; i++ {
		vertices = append(vertices,
			q.Pos[i].X(), q.Pos[i].Y(), 0.0,
			q.Color[i].X(), q.Color[i].Y(), q.Color[i].Z(), q.Color[i].W())
	}
	indices := []uint32{0, 1, 2, 2, 3, 0}
	return vertices, indices
}

type RoundedQuad struct {
	Pos [4]mgl32.Vec2
	Color []mgl32.Vec4
	Radius float32
	Segments int
}
func (rq *RoundedQuad) CreateRoundedQuad() ([]float32, []uint32){
	vertices := make([]float32, 0, (rq.Segments+1)*4*7)
	indices := make([]uint32, 0, rq.Segments*4*6)
	const Pi = math.Pi

	for corner := 0; corner < 4; corner++ {
		center := rq.Pos[corner]
		nextCorner := (corner +1 ) % 4

		prevDir := rq.Pos[(corner+3)%4].Sub(center).Normalize()
		nextDir := rq.Pos[nextCorner].Sub(center).Normalize()
	
		for s := 0; s < rq.Segments; s++ {
			t := float32(s) / float32(rq.Segments-1)
			angle := t * Pi / 2
			cosAngle := float32(math.Cos(float64(angle)))
			sinAngle := float32(math.Sin(float64(angle)))
			dir := prevDir.Mul(cosAngle).Add(nextDir.Mul(sinAngle)).Normalize()
			point := center.Add(dir.Mul(rq.Radius))

			vertices = append(vertices, 
				point.X(), point.Y(), 0.0,
				rq.Color[corner].X(), rq.Color[corner].Y(), 
				rq.Color[corner].Z(), rq.Color[corner].W())
		}
	}
	
	for corner := 0; corner < 4; corner++ {
		start := corner * (rq.Segments+1)
		nextCornerStart := ( (corner+1)%4 * (rq.Segments+1) )

		for s := 0; s < rq.Segments; s++ {
			indices = append(indices, 
				uint32(start+s), uint32(start+s+1), uint32(nextCornerStart+s))
			indices = append(indices, 
				uint32(start+s+1), uint32(nextCornerStart+s+1), uint32(nextCornerStart+s))
		}
	}

	return vertices, indices
}
