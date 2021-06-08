module go_service

go 1.14

require (
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gormigrate/gormigrate/v2 v2.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/loopfz/gadgeto v0.9.0
	github.com/magiconair/properties v1.8.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.4.0
	github.com/wI2L/fizz v0.13.4
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.21.10
)

replace github.com/loopfz/gadgeto v0.9.0 => github.com/we-miks/gadgeto v0.10.2-0.20200628075840-6f87a797e771
