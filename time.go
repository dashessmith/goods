package goods

import "time"

// 一日开始时间 2021-5-26 19:20 ==> 2021-5-26 00:00
func TimeFloorDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, nil)
}

// 次日开始时间 2021-5-26 19:20 ==> 2021-5-27 00:00
func TimeCeilDay(t time.Time) time.Time {
	return TimeFloorDay(t).AddDate(0, 0, 1)
}

// 一周开始时间
func TimeFloorWeek(t time.Time) time.Time {
	res := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, nil)
	switch wd := int(t.Weekday()); wd {
	case 0: // 星期日
		res = res.AddDate(0, 0, -6)
	default: // 周 1 ~ 周 5
		res = res.AddDate(0, 0, -wd+1)
	}
	return res
}

// 下周开始时间
func TimeCeilWeek(t time.Time) time.Time {
	return TimeFloorWeek(t).AddDate(0, 0, 7)
}

// 一月开始时间
func TimeFloorMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, nil)
}

// 下月开始时间
func TimeCeilMonth(t time.Time) time.Time {
	return TimeFloorMonth(t).AddDate(0, 1, 0)
}
