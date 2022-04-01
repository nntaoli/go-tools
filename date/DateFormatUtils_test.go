package date

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	t.Log(FormatLongTime(1511765558806649344, time.Nanosecond, "yyyy/MM/dd HH:mm:ss")) // output 2017/11/27 14:52:38
}

func TestDateFormat2(t *testing.T) {
	t.Log(FormatTime(time.Now(), "yyyy-MM-dd HH:mm"))
}

func TestParseDate(t *testing.T) {
	d := ParseDate("2017-12-26 00:10:00", "yyyy-MM-dd HH:mm:ss")
	t.Log(d.Month(), d.Day())
}
