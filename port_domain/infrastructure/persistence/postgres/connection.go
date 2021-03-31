package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	Url                string
	Host               string
	Port               int
	Username           string
	Password           string
	Database           string
	SSLMode            string
	MaxIdleConnections int
	MaxOpenConnections int
}

func NewPostgresConn(config Config) (db *gorm.DB, err error) {
	if config.Host == "" {
		config.Host = os.Getenv("POSTGRES_HOSTNAME")
	}

	if config.Port == 0 {
		config.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	}

	if config.Username == "" {
		config.Username = os.Getenv("POSTGRES_USERNAME")
	}

	if config.Password == "" {
		config.Password = os.Getenv("POSTGRES_PASSWORD")
	}

	if config.Database == "" {
		config.Database = os.Getenv("POSTGRES_DB_NAME")
	}

	if config.SSLMode == "" {
		config.SSLMode = os.Getenv("POSTGRES_SSLMODE")
		if config.SSLMode == "" {
			config.SSLMode = "disable"
		}
	}

	if config.MaxOpenConnections == 0 {
		config.MaxOpenConnections, _ = strconv.Atoi(os.Getenv("POSTGRES_MAX_CONNECTIONS"))
		if config.MaxOpenConnections == 0 {
			config.MaxOpenConnections = 50
		}
	}

	if config.MaxIdleConnections == 0 {
		config.MaxIdleConnections, _ = strconv.Atoi(os.Getenv("MaxIdleConnections"))
		if config.MaxIdleConnections == 0 {
			config.MaxIdleConnections = 50
		}
	}

	if config.Url != "" {
		reg := regexp.MustCompile("\\/\\/(.+):(.+)@(.+):(.+)\\/(.+)")
		parts := reg.FindStringSubmatch(config.Url)
		config.Username = parts[1]
		config.Password = parts[2]
		config.Host = parts[3]
		config.Port, _ = strconv.Atoi(parts[4])
		config.Database = parts[5]
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.SSLMode,
	)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		FullSaveAssociations:   false,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return
	}

	dbRef, err := db.DB()
	if err != nil {
		return
	}

	if err = dbRef.Ping(); err != nil {
		return
	}

	maxIdleConnections := 5
	maxOpenConnections := 10

	if config.MaxIdleConnections != 0 {
		maxIdleConnections = config.MaxIdleConnections
	}

	if config.MaxOpenConnections != 0 {
		maxOpenConnections = config.MaxOpenConnections
	}

	dbRef.SetMaxIdleConns(maxIdleConnections)
	dbRef.SetMaxOpenConns(maxOpenConnections)

	return
}
