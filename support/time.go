package support

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

var mocks = make(map[string]time.Time)

func Create(value interface{}, layout interface{}) (t time.Time) {
	switch v := value.(type) {
	case int:
		t := time.Unix(int64(v), 0)
		return t
	case string:
		local, _ := time.LoadLocation(Config["app"].GetString("system.timezone"))
		t, _ := time.ParseInLocation(layout.(string), value.(string), local)
		return t
	}
	return
}

func Mock(value interface{}, layout interface{}, names ...string) {
	name := "default"
	if len(names) > 0 {
		name = names[0]
	}
	if value == nil {
		delete(mocks, name)
		return
	}

	mocks[name] = Create(value, layout)
}

func Now(names ...string) time.Time {
	name := "default"
	if len(names) > 0 {
		name = names[0]
	}
	t, ok := mocks[name]
	if ok {
		return t
	}
	local, _ := time.LoadLocation(Config["app"].GetString("system.timezone"))
	return time.Now().In(local)
}

// 自定义time
type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = Time(t1)
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t Time) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *Time) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = Time(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *Time) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
