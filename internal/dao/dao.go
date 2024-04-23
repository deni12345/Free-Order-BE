package dao

import (
	"fmt"
	"github/lambda-microservice/custom"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DNS = `%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local`
)

type Config struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	Option   map[string]string
}

type DAO struct {
	UserDAO IUserDAO
}

func NewDAO(conf Config) *DAO {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		DNS,
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)), &gorm.Config{Logger: &custom.GormLogrus{Logger: *logrus.New()}})
	if err != nil {
		log.Fatalf("Can not connect db on error: %s", err)
		return nil
	}

	return &DAO{
		UserDAO: &UserImpl{client: db},
	}
}
