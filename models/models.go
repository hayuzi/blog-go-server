package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"blog-go-server/pkg/setting"
)

var db *gorm.DB

type Model struct {
	Id        int    `gorm:"primary_key" json:"id"`
	CreatedAt string `json:"createdAt"` // 创建时间 datetime
	UpdatedAt string `json:"updatedAt"` // 更新时间 datetime
	deletedAt int                       // 软删除字段(可以为NULL)  datetime
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name,
			setting.DatabaseSetting.Charset))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func init() {
	var (
		err                                                        error
		dbType, dbName, user, password, host, charset, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("DB_TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	user = sec.Key("DB_USER").String()
	password = sec.Key("DB_PASSWORD").String()
	host = sec.Key("DB_HOST").String()
	charset = sec.Key("DB_CHARSET").String()
	tablePrefix = sec.Key("DB_TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
		charset))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	// Gorm 会自动更新
	// scope.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
