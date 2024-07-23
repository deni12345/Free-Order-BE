module github/lambda-microservice

go 1.18

require (
	gorm.io/driver/mysql v1.5.6
	gorm.io/gorm v1.25.9
)

require golang.org/x/sys v0.19.0 // indirect

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/gorilla/mux v1.8.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.22.0
)
