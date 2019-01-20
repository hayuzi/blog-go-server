// JSONTime类型是自定义的时间类型， 提供给 Gorm 使用. 以用来实现我们自定义的时间格式。
// 我们只需定义一个内嵌time.Time的结构体，并重写MarshalJSON方法，然后在定义model的时候把time.Time类型替换为我们自己的类型即可。
// 但是在gorm中只重写MarshalJSON是不够的，只写这个方法会在写数据库的时候会提示delete_at字段不存在。
// 需要加上database/sql的Value和Scan方法 https://github.com/jinzhu/gorm/issues/1611#issuecomment-329654638。
package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// 在额外的类型上实现 json.Marshaler 以获取自定义 Time格式,
// 因为 go的time包中实现json.Marshaler接口时指定了使用RFC3339Nano这种格式， 但是go的方法不能重写, 所以只能自定义类型
// MarshalJSON implements the json.Marshaler interface.
// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
