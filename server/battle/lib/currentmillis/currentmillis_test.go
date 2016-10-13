package currentmillis

import (
	"testing"
	"time"
)

func TestStubNow(t *testing.T) {
	StubNow = func() int64 { return 1 }
	actual := Now()
	StubNow = nil

	if actual != 1 {
		t.Error("Failed to set stub")
		return
	}
}

func TestMillis(t *testing.T) {
	actual := Millis(time.Date(2016, time.January, 31, 14, 11, 54, 921*1000000, time.UTC))
	var expected int64 = 1454249514921

	if actual != expected {
		t.Errorf("Failed converting to ms: expected=%v, actual=%v", expected, actual)
	}
}

func TestTime(t *testing.T) {
	actualTime := Time(1454249514921)
	actual := actualTime.UTC().String()
	expected := "2016-01-31 14:11:54.921 +0000 UTC"

	if actual != expected {
		t.Errorf("Failed converting to time: expected=%v, actual=%v", expected, actual)
	}
}

func TestDuration(t *testing.T) {
	actual := Duration(1000)
	expected := time.Second

	if actual != expected {
		t.Errorf("Failed converting to time duration: expected=%v, actual=%v", expected, actual)
	}
}

func TestDurationMillis(t *testing.T) {
	actual := DurationMillis(1 * time.Second)
	var expected int64 = 1000

	if actual != expected {
		t.Errorf("Failed converting to milliseconds: expected=%v, actual=%v", expected, actual)
	}
}
