package orm

import (
	"context"
	"tiktokPlus/user/rpc/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

type DB struct {
	*gorm.DB
}

type ormLog struct {
	LogLevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Warn(ctx context.Context, fromat string, v ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(fromat, v...)
}

func (l *ormLog) Error(ctx context.Context, fromat string, v ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Infof(fromat, v...)
}

func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Infof("[%.3f ms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

func NewMysql(conf *Config) (*DB, error) {
	if conf.ConnMaxLifetime == 0 {
		conf.ConnMaxLifetime = 3600
	}

	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}

	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 100
	}

	db := InitGorm(config.Config.Mysql.Datasource)

}
