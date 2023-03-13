package shared

import (
	"time"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

// Now time.Nowはserver依存の為、UTC固定としている
func Now() {
	time.Now().UTC()
}

func JstNow() {
	time.Now().In(jst)
}

func Timezone() (tz *time.Location) {
	tz, err := time.LoadLocation(Env("TZ").GetString("Asia/Tokyo"))
	if err != nil {
		tz = jst
	}
	return
}
