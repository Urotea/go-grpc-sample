package postgres

import (
	"github.com/Urotea/go-grpc-sample/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	log = logger.GetLogger()
)

type PostgresDao struct {
	db *gorm.DB
}

func (db *PostgresDao) Create(user User) {
	db.db.Create(&user)
	log.Debugf("DBへの書き込みが完了しました。 data = ")
}

func New(dsn string) (*PostgresDao, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDao{db: db}, nil
}
