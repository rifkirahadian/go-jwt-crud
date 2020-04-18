package configs

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitGormDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "storage.db")
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()

	return db
}