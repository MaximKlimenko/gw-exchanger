package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/MaximKlimenko/gw-exchanger/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewConnection(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.SSLMode,
	)

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Логирование SQL-запросов
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Использовать имена таблиц в единственном числе
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Ошибка получения SQL-базы из GORM: %v", err)
		return nil, err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	//Проверка соединения
	err = sqlDB.Ping()
	if err != nil {
		log.Printf("Ошибка пинга к базе данных: %v", err)
		return nil, err
	}

	log.Println("\033[1;32mУспешно подключились к базе данных PostgreSQL\033[0m")
	return db, err
}
