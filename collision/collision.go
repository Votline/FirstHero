package collision

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"

	"FirstHero/primShapes"
)

const eps = 0.001
const gravity = 0.01

func CheckCollision(l *primShapes.Limb, gd []*primShapes.Quad, canJump *bool) {
	lLeft := l.CurrentPos[1]
	lRight := l.CurrentPos[2]
	*canJump = false

	for _, block := range gd {
		gdLeftU := block.Pos[0]
		gdRightU := block.Pos[3]
		gdLeftD := block.Pos[1]
		gdRightD := block.Pos[2]
	  log.Printf("\nlLeft: %v | lRight: %v\n gdLeftU: %v | gdRightU: %v\n gdLeftD: %v | gdRightD: %v\n", lLeft, lRight, gdLeftU, gdRightU, gdLeftD, gdRightD)
		
		if (isPointInQuad(lLeft, gdLeftU, gdRightU, gdLeftD, gdRightD) ||
		isPointInQuad(lRight, gdLeftU, gdRightU, gdLeftD, gdRightD)) {
			log.Println("\n\n\n\n\nMATCH\n\n\n\n\n")
			*canJump = true
		
			if l.Parent.TargetPos[0][1] < l.Parent.CurrentPos[0][1] {
				deltaY := lLeft.Y() - gdLeftU.Y() - eps
				for i := 0; i < 4; i++ {
					l.Parent.TargetPos[i][1] -= deltaY
				}
			}
			return
		}
	}
	for i := 0; i < 4; i++ {
		if !*canJump{
			l.Parent.TargetPos[i][1] -= gravity 
		}
  }
}
func isPointInQuad(point, qLeftU, qRightU, qLeftD, qRightD mgl32.Vec2) bool {
	return point.X() >= qLeftU.X()-eps && point.X() <= qRightU.X()-eps &&
				 point.Y() >= qLeftD.Y()-eps && point.Y() <= qRightU.Y()+eps
}
