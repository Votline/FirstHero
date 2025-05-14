package movement

import (
	"FirstHero/primShapes"
)

func UpdatePos(q *primShapes.Quad){
	for i := 0; i < 4; i++{
		q.CurrentPos[i] = q.CurrentPos[i].Mul(1-q.Alpha).Add(q.TargetPos[i].Mul(q.Alpha))
	}
}
func SetTargetPos(q *primShapes.Quad, axis int, delta float32) {
	for i := 0; i < 4; i++ {
		q.TargetPos[i][axis] += delta
		if q.TargetPos[i][axis] < -1.0 {
			q.TargetPos[i][axis] = -1.0
		} else if q.TargetPos[i][0] > 0.9 {
			q.TargetPos[i][axis] = 0.9
		} else if q.TargetPos[i][1] > 0.7 {
			q.TargetPos[i][1] = 0.7
		}
	}
}
