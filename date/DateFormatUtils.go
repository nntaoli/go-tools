package date

import (
	"strings"
	"time"
)

const (
	yyyy = "2006"
	yy   = "06"
	MMMM = "January"
	MMM  = "Jan"
	MM   = "01"
	dddd = "Monday"
	ddd  = "Mon"
	dd   = "02"

	HHT = "03"
	HH  = "15"
	mm  = "04"
	SS  = "05"
	ss  = "05"
	tt  = "PM"
	Z   = "MST"
	ZZZ = "MST"
)

//FormatLongTime
/*
@param timeLong 时间值
@param duration 时间颗粒度
@param format 格式表达式
*/
func FormatLongTime(timeLong int64, duration time.Duration, format string) string {
	var (
		sec, mills int64
	)

	r := int64(time.Second / duration)

	if r == 0 {
		sec = timeLong
		mills = 0
	} else {
		sec = timeLong / r
		mills = (timeLong - sec*r) * int64(duration)
	}

	layout := convertFormat(format)
	t := time.Unix(sec, mills)
	return t.Format(layout)
}

func FormatTime(t time.Time, format string) string {
	return t.Format(convertFormat(format))
}

func ParseDate(dateTimeStr string, format string) time.Time {
	t, _ := time.Parse(convertFormat(format), dateTimeStr)
	return t
}

func convertFormat(format string) string {
	var goFormat = format
	if strings.Contains(goFormat, "YYYY") {
		goFormat = strings.Replace(goFormat, "YYYY", yyyy, -1)
	} else if strings.Contains(goFormat, "yyyy") {
		goFormat = strings.Replace(goFormat, "yyyy", yyyy, -1)
	} else if strings.Contains(goFormat, "YY") {
		goFormat = strings.Replace(goFormat, "YY", yy, -1)
	} else if strings.Contains(goFormat, "yy") {
		goFormat = strings.Replace(goFormat, "yy", yy, -1)
	}

	//month
	if strings.Contains(goFormat, "MMMM") {
		goFormat = strings.Replace(goFormat, "MMMM", MMMM, -1)
	} else if strings.Contains(goFormat, "MMM") {
		goFormat = strings.Replace(goFormat, "MMM", MMM, -1)
	} else if strings.Contains(goFormat, "MM") {
		goFormat = strings.Replace(goFormat, "MM", MM, -1)
	}

	if strings.Contains(goFormat, "mm") { //minute
		goFormat = strings.Replace(goFormat, "mm", mm, -1)
	}

	//day
	if strings.Contains(goFormat, "dddd") {
		goFormat = strings.Replace(goFormat, "dddd", dddd, -1)
	} else if strings.Contains(goFormat, "ddd") {
		goFormat = strings.Replace(goFormat, "ddd", ddd, -1)
	} else if strings.Contains(goFormat, "dd") {
		goFormat = strings.Replace(goFormat, "dd", dd, -1)
	}

	if strings.Contains(goFormat, "tt") {
		if strings.Contains(goFormat, "HH") {
			goFormat = strings.Replace(goFormat, "HH", HHT, -1)
		} else if strings.Contains(goFormat, "hh") {
			goFormat = strings.Replace(goFormat, "hh", HHT, -1)
		}
		goFormat = strings.Replace(goFormat, "tt", tt, -1)
	} else {
		if strings.Contains(goFormat, "HH") {
			goFormat = strings.Replace(goFormat, "HH", HH, -1)
		} else if strings.Contains(goFormat, "hh") {
			goFormat = strings.Replace(goFormat, "hh", HH, -1)
		}
		goFormat = strings.Replace(goFormat, "tt", "", -1)
	}

	//second
	if strings.Contains(goFormat, "SS") {
		goFormat = strings.Replace(goFormat, "SS", SS, -1)
	} else if strings.Contains(goFormat, "ss") {
		goFormat = strings.Replace(goFormat, "ss", SS, -1)
	}

	if strings.Contains(goFormat, "ZZZ") {
		goFormat = strings.Replace(goFormat, "ZZZ", ZZZ, -1)
	} else if strings.Contains(goFormat, "zzz") {
		goFormat = strings.Replace(goFormat, "zzz", ZZZ, -1)
	} else if strings.Contains(goFormat, "Z") {
		goFormat = strings.Replace(goFormat, "Z", Z, -1)
	} else if strings.Contains(goFormat, "z") {
		goFormat = strings.Replace(goFormat, "z", Z, -1)
	}

	if strings.Contains(goFormat, "tt") {
		goFormat = strings.Replace(goFormat, "tt", tt, -1)
	}
	return goFormat
}
