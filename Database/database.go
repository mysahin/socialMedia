package Database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"socialMedia/Models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	var dsn = "root:password@tcp(db:3306)/db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected success")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(Models.User{}, Models.Login{}, Models.Post{}, Models.Follow{}, Models.Files{})
	DB = Dbinstance{
		Db: db,
	}

}
