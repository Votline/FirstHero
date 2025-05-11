package movement

import (
	"FirstHero/primShapes"
)

func UpdatePos(t *primShapes.Triangular){
	t.CurrentPos = t.CurrentPos.Mul(1-t.Alpha).Add(t.TargetPos.Mul(t.Alpha))
}
func SetTargetPos(t *primShapes.Triangular, axis int, delta float32) {
	t.TargetPos[axis] += delta
	if t.TargetPos[axis] < -1.0 {
		t.TargetPos[axis] = -1.0
	} else if t.TargetPos[0] > 0.9 {
		t.TargetPos[axis] = 0.9
	} else if t.TargetPos[1] > 0.7 {
		t.TargetPos[1] = 0.7
	}
}
