package russiantime

import (
	. "math/big"
	. "strconv"
	"strings"
	"time"
)

//Struct contains embedded time.Time
type Time struct {
	time.Time
	year                                       *Int
	month, day, weekday, hour, minute, seconds int
}

//Format time to russian
func (t *Time) FormatRU(layout string) (rustime string) {
	t.year = NewInt(int64(t.Year()))
	t.month = int(t.Month())
	t.day = int(t.Day())
	t.weekday = int(t.Weekday())
	t.hour = int(t.Hour())
	t.minute = int(t.Minute())
	t.seconds = int(t.Second())
	rustime = layout
	if strings.Contains(rustime, "%YYYY") {
		rustime = strings.Replace(rustime, "%YYYY", strings.Repeat("0", 4-len(t.year.String()))+t.year.String(), -1)
	}
	if strings.Contains(rustime, "%YYY") {
		var years Int
		years.Mod(t.year, NewInt(1000))
		rustime = strings.Replace(rustime, "%YYY", strings.Repeat("0", 3-len(years.String()))+years.String(), -1)
	}
	if strings.Contains(rustime, "%YY") {
		var years Int
		years.Mod(t.year, NewInt(100))
		rustime = strings.Replace(rustime, "%YY", strings.Repeat("0", 2-len(years.String()))+years.String(), -1)
	}
	if strings.Contains(rustime, "%Y") {
		var years Int
		years.Mod(t.year, NewInt(10))
		rustime = strings.Replace(rustime, "%Y", strings.Repeat("0", 1-len(years.String()))+years.String(), -1)
	}
	//Parse month
	if strings.Contains(rustime, "%md") {
		rustime = strings.Replace(rustime, "%md", Itoa(int(t.month)), -1)
	}
	if strings.Contains(rustime, "%Md") {
		rustime = strings.Replace(rustime, "%Md", strings.Repeat("0", 2-len(Itoa(int(t.month))))+Itoa(int(t.month)), -1)
	}
	if strings.Contains(rustime, "%m") {
		rustime = strings.Replace(rustime, "%m", shortMonthNames[t.month], -1)
	}
	if strings.Contains(rustime, "%M") {
		rustime = strings.Replace(rustime, "%M", longMonthNames[t.month], -1)
	}
	//Parse Day
	if strings.Contains(rustime, "%D") {
		rustime = strings.Replace(rustime, "%D", strings.Repeat("0", 2-len(Itoa(t.day)))+Itoa(t.day), -1)
	}
	if strings.Contains(rustime, "%d") {
		rustime = strings.Replace(rustime, "%d", Itoa(t.day), -1)
	}
	//Parse week day
	if strings.Contains(rustime, "%Wd") {
		rustime = strings.Replace(rustime, "%Wd", Itoa(int(t.weekday)), -1)
	}
	if strings.Contains(rustime, "%w") {
		rustime = strings.Replace(rustime, "%w", shortDayNames[t.weekday], -1)
	}
	if strings.Contains(rustime, "%W") {
		rustime = strings.Replace(rustime, "%W", longDayNames[t.weekday], -1)
	}
	//Parse hour
	if strings.Contains(rustime, "%H") {
		rustime = strings.Replace(rustime, "%H", strings.Repeat("0", 2-len(Itoa(t.hour)))+Itoa(t.hour), -1)
	}
	if strings.Contains(rustime, "%h") {
		rustime = strings.Replace(rustime, "%h", Itoa(t.hour), -1)
	}
	//Parse minute
	if strings.Contains(rustime, "%N") {
		rustime = strings.Replace(rustime, "%N", strings.Repeat("0", 2-len(Itoa(t.minute)))+Itoa(t.minute), -1)
	}
	if strings.Contains(rustime, "%n") {
		rustime = strings.Replace(rustime, "%n", Itoa(t.minute), -1)
	}
	//Parse second
	if strings.Contains(rustime, "%S") {
		rustime = strings.Replace(rustime, "%S", strings.Repeat("0", 2-len(Itoa(t.seconds)))+Itoa(t.seconds), -1)
	}
	if strings.Contains(rustime, "%s") {
		rustime = strings.Replace(rustime, "%s", Itoa(t.seconds), -1)
	}
	return rustime
}

//DEPRECATED! Remove in future version
func Russian(t time.Time, layout string) (rustime string) {
	var ruTime Time
	ruTime.Time = t
	return ruTime.FormatRU(layout)
}
