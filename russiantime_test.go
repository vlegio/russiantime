package russiantime

import (
	"testing"
	"time"
)

const valid = "2009 009 09 9, 11 11 Ноя Ноябрь, 10 10, 2 Вт Вторник, 23 23, 00 0, 00 0"

func TestFormat(t *testing.T) {
	var testTime Time
	testTime.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	rus_time := testTime.FormatRU("%YYYY %YYY %YY %Y, %md %Md %m %M, %D %d, %Wd %w %W, %H %h, %N %n, %S %s")
	if rus_time == valid {
		t.Log("All ok!")
	} else {
		t.Fatal("Not pass!")
	}
}
