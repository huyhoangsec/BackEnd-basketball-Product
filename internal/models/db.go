package models

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"backend-ocean-basketball/config"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	var dialector gorm.Dialector

	driver := strings.ToLower(cfg.DBDriver)
	isPostgresUrl := strings.HasPrefix(cfg.DBUrl, "postgres://") || strings.HasPrefix(cfg.DBUrl, "postgresql://")

	if isPostgresUrl && !strings.Contains(cfg.DBUrl, "${{") {
		dialector = postgres.Open(cfg.DBUrl)
		log.Println("Connecting to PostgreSQL Database via DATABASE_URL DSN...")
	} else if cfg.DBHost != "" && (driver == "postgres" || driver == "postgresql" || cfg.DBPort == "5432") {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
		dialector = postgres.Open(dsn)
		log.Printf("Connecting to PostgreSQL Database at host=%s port=%s dbname=%s...", cfg.DBHost, cfg.DBPort, cfg.DBName)
	} else if cfg.DBHost != "" && (driver == "mysql" || cfg.DBPort == "3306") {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dsn)
		log.Printf("Connecting to MySQL Database at host=%s port=%s...", cfg.DBHost, cfg.DBPort)
	} else if cfg.DBHost != "" {
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = sqlserver.Open(dsn)
		log.Printf("Connecting to SQL Server Database at host=%s...", cfg.DBHost)
	} else {
		log.Fatalf("Fatal: No database connection parameters provided! Please set DATABASE_URL or PGHOST in Environment Variables.")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 2 * time.Second, // 2s threshold for cloud latency
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Fatal: Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connection established successfully")

	err = db.AutoMigrate(
		&User{},
		&Coach{},
		&Court{},
		&Class{},
		&ClassSchedule{},
		&Student{},
		&TrialRegistration{},
		&AttendanceRecord{},
		&Tournament{},
		&TuitionPlan{},
		&FAQ{},
		&Review{},
		&Banner{},
		&Invoice{},
		&Payment{},
		&Contact{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	SeedDatabase(db)
}

func MockDB() sqlmock.Sqlmock {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to create mock db: %v", err)
	}
	
	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})
	
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open gorm mock db: %v", err)
	}

	DB = db
	return mock
}
