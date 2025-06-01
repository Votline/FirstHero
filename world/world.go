package world

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"

	"FirstHero/primShapes"
)

const eps = 0.001 

type QuadTree struct {
	Bounds   primShapes.Quad
	Capacity int
	Objects  []*primShapes.Quad
	Nodes    [4]*QuadTree
}

func (qt *QuadTree) Contains(q *primShapes.Quad) bool {
	for _, point := range q.Pos {
		if point.X() < qt.Bounds.Pos[0].X()-eps ||
			point.X() > qt.Bounds.Pos[2].X()+eps ||
			point.Y() > qt.Bounds.Pos[0].Y()+eps ||
			point.Y() < qt.Bounds.Pos[2].Y()-eps {
			return false
		}
	}
	return true
}
func (qt *QuadTree) Subdivide() {
	if qt.Nodes[0] != nil {
		return
	}

	centerX := (qt.Bounds.Pos[1].X() + qt.Bounds.Pos[2].X()) / 2
	centerY := (qt.Bounds.Pos[0].Y() + qt.Bounds.Pos[1].Y()) / 2

	qt.Nodes[0] = &QuadTree{ //Northwest (upper left quadrant)
		Bounds: primShapes.Quad{
			Pos: [4]mgl32.Vec3{
				qt.Bounds.Pos[0],
				{qt.Bounds.Pos[1].X(), centerY, 0},
				{centerX, centerY, 0},
				{centerX, qt.Bounds.Pos[0].Y(), 0},
			},
		},
		Capacity: qt.Capacity,
	}
	qt.Nodes[1] = &QuadTree{ //Northeast (upper right)
		Bounds: primShapes.Quad{
			Pos: [4]mgl32.Vec3{
				{centerX, qt.Bounds.Pos[0].Y(), 0},
				{centerX, centerY, 0},
				qt.Bounds.Pos[2],
				{qt.Bounds.Pos[2].X(), qt.Bounds.Pos[0].Y(), 0},
			},
		},
		Capacity: qt.Capacity,
	}
	qt.Nodes[2] = &QuadTree{ //Southwest (lower left)
		Bounds: primShapes.Quad{
			Pos: [4]mgl32.Vec3{
				{qt.Bounds.Pos[1].X(), centerY, 0},
				qt.Bounds.Pos[1],
				{centerX, qt.Bounds.Pos[1].Y(), 0},
				{centerX, centerY, 0},
			},
		},
		Capacity: qt.Capacity,
	}
	qt.Nodes[3] = &QuadTree{ //South-east (lower right)
		Bounds: primShapes.Quad{
			Pos: [4]mgl32.Vec3{
				{centerX, centerY, 0},
				{centerX, qt.Bounds.Pos[1].Y(), 0},
				qt.Bounds.Pos[2],
				{qt.Bounds.Pos[2].X(), centerY, 0},
			},
		},
		Capacity: qt.Capacity,
	}
}
func (qt *QuadTree) Insert(q *primShapes.Quad) bool {
	if !qt.Contains(q) {
		return false
	}

	if len(qt.Objects) < qt.Capacity {
		qt.Objects = append(qt.Objects, q)
		return true
	}

	if qt.Nodes[0] == nil {
		qt.Subdivide()
	}

	for i := 0; i < 4; i++ {
		if qt.Nodes[i].Insert(q) {
			return true
		}
	}
	return false
}
func (qt *QuadTree) Query(searchRegion primShapes.Quad) []*primShapes.Quad {
	var result []*primShapes.Quad

	if !qt.intersects(searchRegion) {
		return result
	}

	result = append(result, qt.Objects...)

	if qt.Nodes[0] != nil {
		for i := 0; i < 4; i++ {
			result = append(result, qt.Nodes[i].Query(searchRegion)...)
		}
	}
	return result
}
func (qt *QuadTree) intersects(reg primShapes.Quad) bool {
	return qt.Bounds.Pos[2].X() >= reg.Pos[0].X() &&
		qt.Bounds.Pos[0].X() <= reg.Pos[2].X() &&
		qt.Bounds.Pos[0].Y() >= reg.Pos[2].Y() &&
		qt.Bounds.Pos[2].Y() <= reg.Pos[0].Y()
}

func CreateWorld() (*QuadTree, []*primShapes.Quad) {
	bounds := primShapes.Quad{
		Pos: [4]mgl32.Vec3{
			{-1, 1, 0},
			{-1, -1, 0},
			{1, -1, 0},
			{1, 1, 0},
		},
	}
	qt := QuadTree{Bounds: bounds, Capacity: 4}

	ground := createGround()
	for _, tile := range ground {
		if !qt.Insert(tile) {
			log.Printf("Failed to insert tile at %v", tile.Pos[0])
		}
	}
	return &qt, ground
}

func createGround() []*primShapes.Quad {
	var ground []*primShapes.Quad

	for i := 0; i < 11; i++ {
		startX := -0.9 + float32(i)*0.1

		tile := primShapes.Quad{Name: "Ground",
			Pos: [4]mgl32.Vec3{
				{startX, -0.7, -0.9}, {startX, -0.9, -0.9},
				{startX + 0.1, -0.9, -0.9}, {startX + 0.1, -0.7, -0.9},
			},
			Color: []mgl32.Vec4{
				{0.0, 1.0, 0.0, 1.0},
				{0.0, 0.0, 0.0, 1.0},
				{0.0, 0.0, 0.0, 1.0},
				{0.0, 1.0, 0.0, 1.0},
			},
		}
		ground = append(ground, &tile)
	}
	return ground
}

func GetViewBounds() primShapes.Quad {
	groundBounds := primShapes.Quad{
		Pos: [4]mgl32.Vec3{
			{-1, -0.6, 0},
			{-1, -1, 0},
			{1, -1, 0},
			{1, -0.6, 0},
		},
	}
	return groundBounds
}
