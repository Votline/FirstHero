package collision

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"

	"FirstHero/primShapes"
)

const eps = 0.001
const gravity = 0.01

func IsGrounded(l *primShapes.Limb, gd []*primShapes.Quad, canJump *bool) {
	lLeft := l.CurrentPos[1]
	lRight := l.CurrentPos[2]
	*canJump = false

	for _, block := range gd {
		gdLeftU := block.Pos[0]
		gdRightU := block.Pos[3]
		gdLeftD := block.Pos[1]
		gdRightD := block.Pos[2]
		
		for _, point := range []mgl32.Vec3{lLeft, lRight} {
			if isInGround(point, gdLeftU, gdRightU, gdLeftD, gdRightD) {
				normal := getCollisionNormal(point, gdLeftU, gdRightU, gdLeftD, gdRightD)

				if normal.Y() > 0.5 {
					*canJump = true

					if l.Parent.TargetPos[0][1] < l.Parent.CurrentPos[0][1] {
						deltaY := point.Y() - gdLeftU.Y() - eps
						for i := range l.Parent.TargetPos {
							l.Parent.TargetPos[i][1] -= deltaY
						}
					}
					return
				}
			}
		} 
		
	}
	for i := 0; i < 4; i++ {
		if !*canJump{
			l.Parent.TargetPos[i][1] -= gravity 
		}
  }
}
func isInGround(point, qLeftU, qRightU, qLeftD, qRightD mgl32.Vec3) bool {
	return point.X() > qLeftU.X()-eps && point.X() < qRightU.X()-eps &&
		   point.Y() > qLeftD.Y()-eps && point.Y() < qRightU.Y()-eps
}
func getCollisionNormal(point, qLeftU, qRightU, qLeftD, qRightD mgl32.Vec3) mgl32.Vec3 {
	center := qLeftU.Add(qRightD).Mul(0.5)
	toPoint := point.Sub(center)

	if toPoint.Y() > 0 {
		return mgl32.Vec3{0, 1, 0}
	}
	return mgl32.Vec3{0, -1, 0}
}

func CheckWallCollision(l *primShapes.Limb, wl []*primShapes.Quad, canMoveLeft *bool, canMoveRight *bool) {
	lLeftU := l.CurrentPos[0]
	lRightU := l.CurrentPos[3]
	lLeftD := l.CurrentPos[1]
	lRightD := l.CurrentPos[2]
	
	for _, block := range wl {
		wlLeft := block.Pos[0]
		wlRight := block.Pos[3]
		
		if math.Abs(float64(lLeftU.Y()-wlLeft.Y())) < float64(0.1) {
			if (isInWall(lLeftU.X(), wlLeft.X(), wlRight.X()) ||
			isInWall(lLeftD.X(), wlLeft.X(), wlRight.X()) ||
			isInWall(lRightU.X(), wlLeft.X(), wlRight.X()) ||
			isInWall(lRightD.X(), wlLeft.X(), wlRight.X())) {
				plCenter := (lLeftU.X() + lRightU.X())/2
				qCenter := (wlLeft.X() + wlRight.X())/2
				if plCenter < qCenter {
					*canMoveRight = false
					deltaX := wlLeft.X() - lRightU.X() - eps
					for i := 0; i < 4; i++ {
						if l.Parent != nil {
							l.Parent.TargetPos[i][0] = l.Parent.CurrentPos[i][0] + deltaX
						} else {
							l.TargetPos[i][0] += deltaX
						}
					}
				} else {
					*canMoveLeft = false
					deltaX := wlRight.X() - lLeftU.X() - eps
					for i := 0; i < 4; i++ {
						if l.Parent != nil {
							l.Parent.TargetPos[i][0] = l.Parent.CurrentPos[i][0] + deltaX
						} else {
							l.TargetPos[i][0] += deltaX
						}
					}
				}
				return
			}
		}
	}
	*canMoveLeft = true
	*canMoveRight = true
}

func isInWall(point, qLeft, qRight float32) bool {
  return point >= qLeft-eps && point <= qRight-eps
}
