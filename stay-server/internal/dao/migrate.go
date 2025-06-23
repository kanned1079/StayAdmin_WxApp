package dao

import "stay-server/internal/models"

func (this *DaoInstance) migrateTables() {
	this.DbDao.AutoMigrate(&models.User{})
}
