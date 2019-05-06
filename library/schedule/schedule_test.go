package schedule

import (
	"context"
	"log"
	"testing"
	"time"
)

func EventSecond(param string) {
	log.Printf("second event value:%v\n", param)

	//time.Sleep(12 * time.Second)
}

func EventMinute(param string) {
	log.Printf("minute event value:%v\n", param)
}

func EventHour(param string) {
	log.Printf("hour event value:%v\n", param)
}

func EventDay(param string) {
	log.Printf("day event value:%v\n", param)
}

func EventDayAtDatetime(param string) {
	log.Printf("EventDayAtDatetime event value:%v\n", param)
}

func EventAtDatetimeOnce(param string) {
	log.Printf("EventAtDatetimeOnce event value:%v\n", param)
}

func TestScheduler(t *testing.T) {
	/*err := EverySeconds(2).Do(EventSecond, "second")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err := EveryMinutes(1).Do(EventMinute, "minute")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}*/

	err := EveryHours(1).Do(EventHour, "hour") // minute event value:hours
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err = EveryDays(1).Do(EventDay, "dat")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err = AtDateTimeOnce(2019, time.April, 28, 19, 51, 10).Do(EventAtDatetimeOnce, "at_datetimeOnce")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	err = EveryDayAtDateTime(19, 51, 10).Do(EventDayAtDatetime, "at_datetimeEvery")
	if err != nil {
		t.Errorf("test schedule error:%v", err.Error())
		return
	}

	ctx, _ := context.WithCancel(context.Background())
	Start(ctx)

	select {}
}
