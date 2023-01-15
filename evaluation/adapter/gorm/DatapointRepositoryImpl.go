package gorm

import (
	"evaluation/adapter/gorm/entities"
	"evaluation/domain/model"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatapointRepositoryImpl struct {
	DB *gorm.DB
}

func NewDatapointRepositoryImpl() DatapointRepository {
	// Because GORM doesn't automatically create the database, it's done manually
	createDatabase()

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Berlin",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}
	err = db.AutoMigrate(&entities.Datapoint{})
	if err != nil {
		panic("Can't migrate database")
	}
	return &DatapointRepositoryImpl{
		DB: db,
	}
}

// createDatabase creates the database on the server if it doesn't already exist.
// Solution found here: https://stackoverflow.com/a/69812883
func createDatabase() {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Berlin",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		"postgres",
	)

	// connect to the postgres db just to be able to run the create db statement
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("Can't open database")
	}

	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", os.Getenv("POSTGRES_DATABASE"))
	rs := db.Raw(stmt)
	if rs.Error != nil {
		panic("Can't request existing databases")
	}

	// if not create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("POSTGRES_DATABASE"))
		if rs := db.Exec(stmt); rs.Error != nil {
			panic("Can't create database")
		}

		// close db connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			panic("Can't close database connection")
		}
	}
}

func (repository *DatapointRepositoryImpl) Save(datapoint *model.Datapoint) {
	datapointEntity := entities.Datapoint{
		Value: datapoint.Value,
		Date:  datapoint.Date,
	}

	repository.DB.Save(&datapointEntity)
}

func (repository *DatapointRepositoryImpl) FindForTime(start, end *time.Time) []*model.Datapoint {
	if start == nil {
		startValue := time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
		start = &startValue
	}
	if end == nil {
		endValue := time.Now()
		end = &endValue
	}

	var datapointEntities []*entities.Datapoint
	repository.DB.Where("date BETWEEN ? AND ?", start, end).Find(&datapointEntities)

	datapoints := make([]*model.Datapoint, len(datapointEntities))
	for i, entity := range datapointEntities {
		datapoints[i] = entity.ToDatapoint()
	}

	return datapoints
}
