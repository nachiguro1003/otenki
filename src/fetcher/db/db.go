package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresEntity struct {
	db       *gorm.DB
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func InitDB(host, port, user, dbname, password string) *PostgresEntity {
	return &PostgresEntity{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   dbname,
		Password: password,
	}
}

func (p *PostgresEntity) Connect() (db *gorm.DB,err error) {
	p.db, err = gorm.Open(
		"postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		p.Host, p.Port, p.User, p.DBName, p.Password))
	if err != nil {
		return nil,err
	}
	return p.db,nil
}
