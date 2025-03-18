package constVar

import (
	"batchLog/model"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	DBConst	model.Dbconfig
)


