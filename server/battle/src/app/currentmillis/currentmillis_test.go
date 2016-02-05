package currentmillis_test

import (
	"app/currentmillis"
	"testing"
	"time"
)

func TestStubNow(t *testing.T) {
	currentmillis.StubNow = func() int64 { return 1 }
	actual := currentmillis.Now()
	currentmillis.StubNow = nil

	if actual != 1 {
		t.Error("Failed to set stub")
		return
	}
}

func TestMilliseconds(t *testing.T) {
	actual := currentmillis.Milliseconds(time.Date(2016, time.January, 31, 14, 11, 54, 921*1000000, time.UTC))
	var expected int64 = 1454249514921

	if actual != expected {
		t.Errorf("Failed to convert the time: expected=%v, actual=%v", expected, actual)
	}
}

func TestTime(t *testing.T) {
	actualTime := currentmillis.Time(1454249514921)
	actual := actualTime.UTC().String()
	expected := "2016-01-31 14:11:54.921 +0000 UTC"

	if actual != expected {
		t.Errorf("Failed to convert the currentMillis: expected=%v, actual=%v", expected, actual)
	}
}
