package mysql

import "github.com/jinzhu/gorm"

type UtilityMySQL struct {
	db *gorm.DB
}

func (u *UtilityMySQL) GetTableCreateSQL() {
	sql := "SHOW CREATE TABLE information_schema.TABLES"
	rows, err := u.db.Raw(sql).Rows()
	if err != nil {
		return
	}
}
