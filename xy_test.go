package goods

import "testing"

func TestXYDistance(t *testing.T) {
	dis := XYDistance(112.937748, 28.223977, 112.936748, 28.223977)
	t.Logf("%v\n", dis)
}

func TestXYDistanceM(t *testing.T) {
	dis := XYDistanceMeter(112.937748, 28.223977, 112.936748, 28.223977)
	t.Logf("%v\n", dis)
}
