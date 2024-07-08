package initializers

import (
	"log/slog"
	"time"

	"github.com/millionsmonitoring/millionsgocore/env"
	"github.com/millionsmonitoring/millionsgocore/utils/timing"
	"github.com/millionsmonitoring/millionsgocore/wrappers/asynqwrapper"

	"github.com/hibiken/asynq"
)

func InitAsynq() (*asynqwrapper.TaskClient, error) {
	conn := asynq.RedisClientOpt{
		// Addr:      fmt.Sprintf("%s:%d", config.Secrets().WorkerRedisHost, config.Secrets().WorkerRedisPort),
		// Password:  config.Secrets().WorkerRedisPassword,
		// PoolSize:  config.Secrets().WorkerRedisPoolSize,
		// DB:        config.Secrets().WorkerRedisDB,
		// TLSConfig: &tls.Config{},
	}
	asynqConfigs := asynq.Config{
		// Logger: slog.Default(),
		Queues: map[string]int{
			"high":    6,
			"default": 3,
			"low":     1,
		},
		LogLevel: asynq.InfoLevel,
	}
	opts := &asynq.SchedulerOpts{
		// Logger:   zap.S(),
		LogLevel: asynq.InfoLevel,
		Location: timing.IndiaTimeZone(),
		PreEnqueueFunc: func(task *asynq.Task, opts []asynq.Option) {
			slog.Info("Job %s with started at %v", task.Type(), time.Now())
		},
		PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
			slog.Info("Job %s with id %s ended at %v", info.Type, info.ID, "time", time.Now())
			if err != nil {
				slog.Error("Job %s with id %s ended at %v with error %v", info.Type, info.ID, "time", time.Now(), "error", err)
			}
		},
	}
	// configs
	if env.IsDevelopment() {
		asynqConfigs.LogLevel = asynq.DebugLevel
		opts.LogLevel = asynq.DebugLevel
		conn.TLSConfig = nil

	}
	task := asynqwrapper.NewClientAndServer(conn, asynqConfigs, opts)
	return task, nil
}
