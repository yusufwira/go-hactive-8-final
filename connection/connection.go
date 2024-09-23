package connection

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	DB *gorm.DB
	Tx *gorm.DB // handle transaction
}

func NewConnection() (*GormDB, error) {
	conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		"postgres", "", "localhost", "5432", "golang")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	log.Info().Msg("connected successfully to the database with gorm!")

	return &GormDB{
		DB: db,
	}, nil
}

func (g *GormDB) BeginTransaction() {
	g.Tx = g.DB.Begin()
}

func (g *GormDB) CommitTransaction() {
	g.Tx.Commit()
	g.Tx = nil
}

func (g *GormDB) RollbackTransaction() {
	g.Tx.Rollback()
	g.Tx = nil
}

func (g *GormDB) GetDB() *gorm.DB {
	if g.Tx != nil {
		return g.Tx
	}
	return g.DB
}
