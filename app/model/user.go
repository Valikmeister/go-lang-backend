package userModel

import (

	"gorm.io/gorm"
)

// User Struct which contains a gorm model in JSON format

type User struct {

	gorm.Model

	Username string `gorm:"unique_index;not null" json:"username"`
	Email 	 string `gorm:"unique_index;not null" json:"email`
	Password string `gorm"not null" json:"password"`
	// Id		 int32	'gorm:"type:uuid;default:uuid_generate_v4()"'
	// Need to check proper uuid framework (is uuid-ossp for pgDB really needed for v4?)




}
