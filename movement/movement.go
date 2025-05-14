package movement

import (
	"FirstHero/player"
)

func UpdatePos(p *player.Player, l *player.Limb){
	if l.Parent == nil {
		for i := range l.LocalPos {
			l.WorldPos[i] = bone.LocalPos[i]
		}
	} else {
		for i := range l.LocalPos {
			l.WorldPos[i] = l.Parent.WorldPos[i].Add(l.LocalPos[i])
		}
	}

	for _, child := range l.Limbs {
		if child.Parent == l {
			p.UpdatePos(child)
		}
	}
}
func SetTargetPos(l *player.Limb, axis int, delta float32) {
	for i := 0; i < 4; i++ {
		l.LocalPos[i][axis] += delta

		if l.LocalPos[i][axis] < -1.0 {
			l.LocalPos[i][axis] = -1.0
		} else if l.LocalPos[i][0] > 0.9 {
			l.LocalPos[i][axis] = 0.9
		} else if l.LocalPos[i][1] > 0.7 {
			l.LocalPos[i][1] = 0.7
		}
	}
}
