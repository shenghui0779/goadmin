package helpers

import (
	"context"
	"fmt"
	"goadmin/pkg/logger"
	"math"
	"sync"
	"time"

	"go.uber.org/zap"
)

type TimingHandler func(ctx context.Context, taskID string) error

type TimingTask struct {
	ctx      context.Context
	round    int
	addedAt  time.Time
	callback TimingHandler
}

// TimingWheel 单时间轮
type TimingWheel struct {
	slot     int
	interval time.Duration
	size     int
	tasks    []sync.Map
	stop     chan struct{}
}

func NewTimingWheel(interval time.Duration, size int) (*TimingWheel, error) {
	if interval < time.Second {
		interval = time.Second

		logger.Warn(context.Background(), "TimingWheel minimum accuracy is 1 second")
	}

	tw := &TimingWheel{
		interval: interval,
		size:     size,
		tasks:    make([]sync.Map, size),
		stop:     make(chan struct{}),
	}

	go tw.scheduler()

	return tw, nil
}

func (tw *TimingWheel) AddTask(ctx context.Context, taskID string, callback TimingHandler, delay time.Duration) {
	select {
	case <-tw.stop:
		logger.Warn(ctx, "TimingWheel has stoped")

		return
	default:
	}

	if delay < tw.interval {
		logger.Warn(ctx, "TimingWheel minimum accuracy is 1 second")

		if delay <= 0 {
			if err := callback(ctx, taskID); err != nil {
				logger.Err(ctx, fmt.Sprintf("TimingWheel task [%v] run error", taskID), zap.Error(err))

				return
			}

			logger.Info(ctx, fmt.Sprintf("TimingWheel task [%v] run ok", taskID), zap.String("delay", delay.String()))

			return
		}
	}

	task := &TimingTask{
		ctx:      CtxCopyWithReqID(ctx),
		addedAt:  time.Now(),
		callback: callback,
	}

	slot := tw.calcSlot(task, delay)

	tw.tasks[slot].Store(taskID, task)
}

func (tw *TimingWheel) Stop() {
	select {
	case <-tw.stop:
		logger.Warn(context.Background(), "TimingWheel has stoped")

		return
	default:
		close(tw.stop)
	}
}

func (tw *TimingWheel) calcSlot(task *TimingTask, delay time.Duration) int {
	interval := int(tw.interval.Seconds())
	total := interval * tw.size
	duration := int(math.Ceil(delay.Seconds()))

	if duration > total {
		task.round = duration / total
		duration = duration % total

		if duration == 0 {
			task.round--
		}
	}

	return (tw.slot + duration/interval) % tw.size
}

func (tw *TimingWheel) scheduler() {
	ctx := context.Background()

	defer Recover(ctx)

	ticker := time.NewTicker(tw.interval)
	defer ticker.Stop()

	for {
		select {
		case <-tw.stop:
			logger.Info(ctx, fmt.Sprintf("TimingWheel stoped at: %s", time.Now().String()))

			return
		case <-ticker.C:
			tw.slot = (tw.slot + 1) % tw.size
			go tw.run(tw.slot)
		}
	}
}

func (tw *TimingWheel) run(slot int) {
	defer Recover(context.Background())

	taskM := tw.tasks[slot]

	taskM.Range(func(key, value interface{}) bool {
		taskID := key.(string)
		task := value.(*TimingTask)

		if task.round > 0 {
			task.round--

			return true
		}

		go func() {
			defer Recover(task.ctx)

			if err := task.callback(task.ctx, taskID); err != nil {
				logger.Err(task.ctx, fmt.Sprintf("TimingWheel task [%s] run error", taskID), zap.Error(err), zap.String("delay", time.Since(task.addedAt).String()))

				return
			}

			logger.Info(task.ctx, fmt.Sprintf("TimingWheel task [%s] run ok", taskID), zap.String("delay", time.Since(task.addedAt).String()))
		}()

		taskM.Delete(key)

		return true
	})
}
