package schedule

import (
	"context"
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Time location, default set by the time.Local (*time.Location)
var timeLocal = time.Local

// Change the time location
func ChangeTimeLocation(newLocal *time.Location) {
	timeLocal = newLocal
}

type Task struct {
	isOnce   bool
	isFirst  bool
	interval time.Duration
	running  bool
	lastRun  time.Time
	gName    string
	gFunc    map[string]interface{}
	gParams  map[string]([]interface{})
}

type Scheduler struct {
	running bool
	time    *time.Ticker
	tasks   []*Task
	sync.RWMutex
}

var schedule *Scheduler

func newScheduler() *Scheduler {
	if schedule == nil {
		schedule = &Scheduler{
			running: false,
			tasks:   make([]*Task, 0),
		}
	}

	return schedule
}

func every(firstRun time.Time, interval uint64, once bool) *Task {
	if interval <= 0 {
		interval = 1
	}

	newTask := &Task{
		isOnce:   once,
		isFirst:  true,
		interval: time.Duration(interval),
		lastRun:  firstRun,
		gName:    "",
		gFunc:    make(map[string]interface{}, 0),
		gParams:  make(map[string][]interface{}, 0),
	}

	if once {
		newTask.lastRun = time.Unix(int64(interval), 0)
	}

	if schedule == nil {
		newScheduler()
	}

	schedule.Add(newTask)

	return newTask
}

func EverySeconds(interval uint64) *Task {
	return every(time.Now().Add(time.Duration(interval)*time.Second), interval, false)
}

func EveryMinutes(interval uint64) *Task {
	return every(time.Now().Add(time.Duration(interval)*time.Minute), interval*60, false)
}

func EveryHours(interval uint64) *Task {
	return every(time.Now().Add(time.Duration(interval)*time.Hour), interval*60*60, false)
}

func EveryDays(interval uint64) *Task {
	return every(time.Now().Add(time.Duration(interval*24)*time.Hour), interval*60*60*24, false)
}

// note: Execute once at a certain point in time
func AtDateTimeOnce(year int, month time.Month, day, hour, minute, second int) *Task {
	return every(time.Now(), uint64(time.Date(year, month, day, hour, minute, second, 0, timeLocal).Unix()), true)
}

func EveryDayAtDateTime(hour, minute, second int) *Task {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	locInterval := 60 * 60 * 24 // 24 hour

	return every(time.Date(year, month, day, hour, minute, second, 0, timeLocal), uint64(locInterval), false)
}

func (tk *Task) Do(taskFun interface{}, params ...interface{}) error {
	typ := reflect.TypeOf(taskFun)
	if typ.Kind() != reflect.Func {
		return errors.New("param taskFun type error")
	}

	funcName := runtime.FuncForPC(reflect.ValueOf(taskFun).Pointer()).Name() + strconv.Itoa(int(time.Now().Unix()))
	tk.gName = funcName
	tk.gFunc[funcName] = taskFun
	tk.gParams[funcName] = params

	return nil
}

func (tk *Task) run(locNow time.Time) (result []reflect.Value, err error) {
	if tk.isOnce && (locNow.Unix()-tk.lastRun.Unix() > 0) {
		return
	}

	if tk.running {
		return
	}

	tk.running = true
	defer func() {
		tk.running = false
		tk.isFirst = false
	}()

	if tk.isFirst && (locNow.Unix()-tk.lastRun.Unix()) > 0 {
		tk.lastRun = tk.lastRun.AddDate(0, 0, 1)
		tk.isFirst = false
		return
	}

	if (locNow.Unix() - tk.lastRun.Unix()) < 0 {
		return
	}

	f := reflect.ValueOf(tk.gFunc[tk.gName])
	params := tk.gParams[tk.gName]
	if len(params) != f.Type().NumIn() {
		err = errors.New(" param num not adapted ")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)

	if tk.isOnce {
		tk.lastRun = time.Unix(0, 0)
	}

	nextTime := locNow.Add(tk.interval * time.Second)
	tk.lastRun = nextTime

	return
}

func (sc *Scheduler) checkTaskStatus(isWait bool) bool {
retry:
	sc.RLock()
	for _, taskItem := range sc.tasks {
		locTask := taskItem

		if locTask.running {
			if isWait {
				time.Sleep(5 * time.Millisecond)
				goto retry
			}

			return false
		}
	}
	sc.RUnlock()

	return true
}

func (sc *Scheduler) Add(value *Task) *Scheduler {

	if value == nil {
		return sc
	}

	sc.Lock()
	sc.tasks = append(sc.tasks, value)
	sc.Unlock()

	return sc
}

func (sc *Scheduler) runAll(locNow time.Time) {

	sc.RLock()
	for _, taskItem := range sc.tasks {
		locTask := taskItem

		go func(task *Task) {
			_, _ = task.run(locNow)
		}(locTask)
	}
	sc.RUnlock()

	return
}

func Start(context context.Context) {
	if schedule == nil {
		newScheduler()
	}

	if schedule.running {
		return
	}

	schedule.Lock()
	schedule.running = true
	schedule.time = time.NewTicker(1 * time.Second)
	schedule.Unlock()

	go func() {
		for {
			select {
			case locNow := <-schedule.time.C:
				schedule.runAll(locNow)

			case <-context.Done():
				schedule.time.Stop()
				schedule.checkTaskStatus(true)
				return
			}
		}
	}()
}
