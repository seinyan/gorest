package migrations

import (
	"fmt"
	"github.com/seinyan/gorest/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Println("Migrate ...")

	//db.Migrator().DropTable(&models.Call{}, models.CallHistory{})

	db.AutoMigrate(models.Call{}, models.CallHistory{})



	//db.Model(&models.User{}).AddForeignKey("person_id", "api_person(id)", "CASCADE", "CASCADE")
	//db.Model(&models.User{}).AddUniqueIndex("api_person_id_unique", "person_id")


	fmt.Println("Migrate end ...")
}