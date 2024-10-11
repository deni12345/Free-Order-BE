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
	SheetDAO ISheetDAO
	UserDAO  IUserDAO
}

func NewDAO(conf Config) *DAO {
	logger := custom.NewGormLogrus().SetLevel(logrus.InfoLevel)
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		DNS,
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)), &gorm.Config{Logger: logger})
	if err != nil {
		log.Fatalf("Can not connect db on error: %s", err)
		return nil
	}

	userImpl := &UserImpl{client: db}
	sheetImpl := &SheetImpl{client: db}
	return &DAO{
		UserDAO:  userImpl,
		SheetDAO: sheetImpl,
	}
}
