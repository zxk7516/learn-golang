package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDb 初始化数据连接
func InitDb() (err error) {
	config := GetConfig()
	db_url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", config.PgHost, config.PgPort, config.PgUser, config.PgDb, config.PgPass)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: db_url,
	}), &gorm.Config{})
	return
}

func TestDb() {
	var result string
	var iresult int
	DB.Raw("select now() as res;").Scan(&result)
	fmt.Println(result)
	DB.Raw("select ?::int + ?::int as res;", 1, 1).Scan(&iresult)
	fmt.Println(iresult)

}
