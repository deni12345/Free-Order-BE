package custom

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogrus struct {
	log.Logger

	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func NewGormLogrus() *GormLogrus {
	return &GormLogrus{
		Logger:                    *log.New(),
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
	}
}

// LogMode(LogLevel) Interface
// 	Info(context.Context, string, ...interface{})
// 	Warn(context.Context, string, ...interface{})

// 	Error(context.Context, string, ...interface{})
// 	Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)

func (l *GormLogrus) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogrus) Info(ctx context.Context, s string, args ...interface{}) {
	l.WithFields(log.Fields{"context": ctx, "data": args}).Info(s)
}

func (l *GormLogrus) Warn(ctx context.Context, s string, args ...interface{}) {
	l.WithFields(log.Fields{"context": ctx, "data": args}).Warn(s)
}

func (l *GormLogrus) Error(ctx context.Context, s string, args ...interface{}) {
	l.WithFields(log.Fields{"context": ctx, "data": args}).Error(s)
}

func (l *GormLogrus) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.GetLevel() >= log.ErrorLevel && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Error(ctx, err.Error(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Error(ctx, err.Error(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.GetLevel() >= log.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Warn(ctx, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Warn(ctx, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.GetLevel() == log.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Info(ctx, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Info(ctx, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}

}

func (l *GormLogrus) SetLevel(level log.Level) *GormLogrus {
	l.Logger.SetLevel(level)
	return l
}

func (l *GormLogrus) SetSlowThreshold(threshold time.Duration) {
	if threshold == 0 {
		l.SlowThreshold = time.Microsecond
		return
	}
	l.SlowThreshold = threshold
}

func (l *GormLogrus) SetIgnoreRecordNotFoundError(ignore *bool) {
	if ignore == nil {
		l.IgnoreRecordNotFoundError = true
		return
	}
	l.IgnoreRecordNotFoundError = *ignore
}
