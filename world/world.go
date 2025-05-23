package world 

import (
	"github.com/go-gl/mathgl/mgl32"
	
	"FirstHero/primShapes"
)

func CreateGround() []*primShapes.Quad {
	var ground []*primShapes.Quad
	var offset float32 = 0.0
	for i := 0; i < 11; i++ {
		q := primShapes.Quad{Name: "Ground",
			Pos: [4]mgl32.Vec3{
				{-0.9+offset, -0.7, -0.9}, {-0.9+offset, -0.9, -0.9},
				{-0.8+offset, -0.9, -0.9}, {-0.8+offset, -0.7, -0.9},
			},
			Color: []mgl32.Vec4{
				{0.0, 1.0, 0.0, 1.0},
      	{0.0, 0.0, 0.0, 1.0},
      	{0.0, 0.0, 0.0, 1.0},
      	{0.0, 1.0, 0.0, 1.0},
			},
		}
		ground = append(ground, &q)
		offset += 0.1
	}
	offset = 0.3
	    q := primShapes.Quad{Name: "Ground",
      Pos: [4]mgl32.Vec3{
        {-0.9+offset, -0.5, -0.9}, {-0.9+offset, -0.7, -0.9},
        {-0.8+offset, -0.7, -0.9}, {-0.8+offset, -0.5, -0.9},
      },
      Color: []mgl32.Vec4{
        {0.0, 1.0, 0.0, 1.0},
        {0.0, 0.0, 0.0, 1.0},
        {0.0, 0.0, 0.0, 1.0},
        {0.0, 1.0, 0.0, 1.0},
      },
    }
		ground = append(ground, &q)
	return ground
}
