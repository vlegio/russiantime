package russiantime

import (
	. "math/big"
	. "strconv"
	"strings"
	"time"
)

func Russian(t time.Time, layout string) (rustime string) {
	year := NewInt(int64(t.Year()))
	month := t.Month()
	day := t.Day()
	weekday := t.Weekday()
	hour := t.Hour()
	minute := t.Minute()
	seconds := t.Second()
	//TODO
	//	timezoneStr, timezoneOffset := t.Zone()
	//	nanosecond := t.Nanosecond()
	rustime = layout
	//Parse year
	if strings.Contains(rustime, "%YYYY") {
		rustime = strings.Replace(rustime, "%YYYY", strings.Repeat("0", 4-len(year.String()))+year.String(), -1)
	}
	if strings.Contains(rustime, "%YYY") {
		var years Int
		years.Mod(year, NewInt(1000))
		rustime = strings.Replace(rustime, "%YYY", strings.Repeat("0", 3-len(years.String()))+years.String(), -1)
	}
	if strings.Contains(rustime, "%YY") {
		var years Int
		years.Mod(year, NewInt(100))
		rustime = strings.Replace(rustime, "%YY", strings.Repeat("0", 2-len(years.String()))+years.String(), -1)
	}
	if strings.Contains(rustime, "%Y") {
		var years Int
		years.Mod(year, NewInt(10))
		rustime = strings.Replace(rustime, "%Y", strings.Repeat("0", 1-len(years.String()))+years.String(), -1)
	}
	//Parse month
	if strings.Contains(rustime, "%md") {
		rustime = strings.Replace(rustime, "%md", Itoa(int(month)), -1)
	}
	if strings.Contains(rustime, "%Md") {
		rustime = strings.Replace(rustime, "%Md", strings.Repeat("0", 2-len(Itoa(int(month))))+Itoa(int(month)), -1)
	}
	if strings.Contains(rustime, "%m") {
		rustime = strings.Replace(rustime, "%m", shortMonthNames[month], -1)
	}
	if strings.Contains(rustime, "%M") {
		rustime = strings.Replace(rustime, "%M", longMonthNames[month], -1)
	}
	//Parse Day
	if strings.Contains(rustime, "%D") {
		rustime = strings.Replace(rustime, "%D", strings.Repeat("0", 2-len(Itoa(day)))+Itoa(day), -1)
	}
	if strings.Contains(rustime, "%d") {
		rustime = strings.Replace(rustime, "%d", Itoa(day), -1)
	}
	//Parse week day
	if strings.Contains(rustime, "%Wd") {
		rustime = strings.Replace(rustime, "%Wd", Itoa(int(weekday)), -1)
	}
	if strings.Contains(rustime, "%w") {
		rustime = strings.Replace(rustime, "%w", shortDayNames[weekday], -1)
	}
	if strings.Contains(rustime, "%W") {
		rustime = strings.Replace(rustime, "%W", longDayNames[weekday], -1)
	}
	//Parse hour
	if strings.Contains(rustime, "%H") {
		rustime = strings.Replace(rustime, "%H", strings.Repeat("0", 2-len(Itoa(hour)))+Itoa(hour), -1)
	}
	if strings.Contains(rustime, "%h") {
		rustime = strings.Replace(rustime, "%h", Itoa(hour), -1)
	}
	//Parse minute
	if strings.Contains(rustime, "%N") {
		rustime = strings.Replace(rustime, "%N", strings.Repeat("0", 2-len(Itoa(minute)))+Itoa(minute), -1)
	}
	if strings.Contains(rustime, "%n") {
		rustime = strings.Replace(rustime, "%n", Itoa(minute), -1)
	}
	//Parse second
	if strings.Contains(rustime, "%S") {
		rustime = strings.Replace(rustime, "%S", strings.Repeat("0", 2-len(Itoa(seconds)))+Itoa(seconds), -1)
	}
	if strings.Contains(rustime, "%s") {
		rustime = strings.Replace(rustime, "%s", Itoa(seconds), -1)
	}
	return rustime
}
