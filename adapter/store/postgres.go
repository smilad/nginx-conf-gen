package store

import (
	_ "github.com/jackc/pgx/v4/pgxpool"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"nginx/models"

	"gorm.io/gorm"
	"log"
	"nginx/config"
)

// NewGorm method job is connect to postgres database and check migration
func NewGorm() *gorm.DB {
	//var err error
	db := new(gorm.DB)
	var err error
	dsn := "postgresql://" + config.C().Postgres.Username + ":" + config.C().Postgres.Password + "@" + config.C().Postgres.Host + "/" + config.C().Postgres.Schema

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("error in database connection %v", err)
	}
	if config.C().Service.Debug {
		db = db.Debug()
	}
	if err := db.AutoMigrate(&models.DomainAddr{}); err != nil {
		log.Fatalf("error in migration, %v", err.Error())
	}
	log.Printf("postgres database loaded successfully \n")
	return db
}
