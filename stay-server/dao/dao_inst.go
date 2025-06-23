package dao

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"stay-server/routers"
)

type DaoInstance struct {
	Id       int32
	DbDao    *gorm.DB
	DbConfig routers.DbConfig
}

func NewDaoInstance(id int32) *DaoInstance {
	var err error = nil
	var daoInst *DaoInstance = &DaoInstance{Id: id}
	if err = daoInst.readDatabaseConfig(); err != nil {
		panic(fmt.Sprintf("failure read config: %v", err))
	}
	daoInst.DbDao, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			daoInst.DbConfig.MysqlConfig.Username,
			daoInst.DbConfig.MysqlConfig.Password,
			daoInst.DbConfig.MysqlConfig.Protocol,
			daoInst.DbConfig.MysqlConfig.Host,
			daoInst.DbConfig.MysqlConfig.Port,
			daoInst.DbConfig.MysqlConfig.Database),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %v", err))
	}
	if daoInst.DbDao.Exec(`SELECT 1+1;`).Error != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	log.Println("Server is connected and tested.")
	return daoInst
}

func (d *DaoInstance) readDatabaseConfig() error {
	file, err := os.ReadFile("config/params.yaml")
	if err != nil {
		return fmt.Errorf("read yaml file error: %w", err)
	}
	var cfg routers.DbConfig
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return fmt.Errorf("unmarshal yaml error: %w", err)
	}
	d.DbConfig = cfg
	return nil
}
