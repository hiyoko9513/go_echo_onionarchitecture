package util

import (
	"time"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

// Now time.Nowはserver依存の為、UTC固定としている
func Now() time.Time {
	return time.Now().UTC()
}

func JstNow() time.Time {
	return time.Now().In(jst)
}

// ToRFC3339 SO 8601形式の文字列に変換
func ToRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func LoadTimezone() {
	time.Local = Timezone()
}

func Timezone() (tz *time.Location) {
	tz, err := time.LoadLocation(Env("TZ").GetString("Asia/Tokyo"))
	if err != nil {
		tz = jst
	}
	return
}
