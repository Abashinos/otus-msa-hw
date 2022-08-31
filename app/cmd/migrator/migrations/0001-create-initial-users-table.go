package migrations

import "gorm.io/gorm"

func init() {

	Add("0001",
		func(tx *gorm.DB) error {
			type User struct {
				gorm.Model
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			}

			return tx.AutoMigrate(&User{})
		},
		func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		})
}
