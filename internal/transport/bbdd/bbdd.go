package bbdd

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewBBDD(dbURI string, maxIdelConns int, maxOpenConns int) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(maxIdelConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	return db, nil
}
