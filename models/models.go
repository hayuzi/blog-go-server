package models

import (
	"blog-go-server/pkg/setting"
	"blog-go-server/pkg/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	Id        int            `gorm:"primary_key" json:"id"`
	CreatedAt util.JSONTime  `json:"createdAt"`     // 创建时间 datetime
	UpdatedAt util.JSONTime  `json:"updatedAt"`     // 更新时间 datetime
	DeletedAt *util.JSONTime `sql:"index" json:"-"` // 软删除字段(可以为NULL)  datetime
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type,
		// %s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai  使用上海东八区，但是在scratch的docker镜像中不可用
		// %s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local， 服务器本地时区未设置，
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

	// 此处注释掉的内容是替换预置回调函数。 只是参考. 目前项目中不会用他
	// 可参考文档 http://gorm.io/docs/write_plugins.html#Register-a-new-callback
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	// scope.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("created_at", time.Now())
	// scope.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("updated_at", time.Now())
	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	// Gorm 会自动更新
	// scope.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
