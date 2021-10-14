package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

var (
	Orm *xorm.Engine
)

func init() {
	godotenv.Load(".env")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var err error
	Orm, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	migrate()

	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
	Orm.SetDefaultCacher(cacher)
}
