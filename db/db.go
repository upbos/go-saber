package db

import (
	"fmt"
	"github.com/upbos/go-base/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type DataSource struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Database        string        `yaml:"database"`
	MaxIdleConns    int           `yaml:"max-idle-conns"`
	MaxOpenConns    int           `yaml:"max-open-conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn-max-idle-time"`
	ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime"`
}

var DB *gorm.DB

type dbWriter struct{}

func (w *dbWriter) Printf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func Init(ds *DataSource) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		ds.Host, ds.Port, ds.User, ds.Database, ds.Password)

	slowLogger := logger.New(
		&dbWriter{},
		logger.Config{
			Colorful:                  false,
			SlowThreshold:             2 * time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: slowLogger,
	})

	if err != nil {
		log.Error(err, "initial the database error")
		os.Exit(1)
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Error(err, "initial the database connection pool error")
		os.Exit(1)
	}
	if ds.MaxIdleConns < 10 {
		ds.MaxIdleConns = 10
	}
	if ds.MaxOpenConns < 100 {
		ds.MaxOpenConns = 100
	}
	if ds.ConnMaxLifetime < 30*time.Minute {
		ds.ConnMaxLifetime = time.Hour
	}
	if ds.ConnMaxIdleTime < time.Minute {
		ds.ConnMaxIdleTime = time.Minute
	}
	sqldb.SetMaxIdleConns(ds.MaxIdleConns)
	sqldb.SetConnMaxIdleTime(ds.ConnMaxIdleTime)
	sqldb.SetMaxOpenConns(ds.MaxOpenConns)
	sqldb.SetConnMaxLifetime(ds.ConnMaxLifetime)
	DB = db

	log.Info("Connected to the database successfully")
}
