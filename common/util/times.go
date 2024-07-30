// 时间工具类
package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var (
	formatTime = "2006-01-02 15:04:05"
)

func (t HTime) MarshalJSON() ([]byte, error) {
	formatted := t.Format(formatTime)
	return []byte(`"` + formatted + `"`), nil
}

func (t *HTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*t = HTime{Time: time.Time{}}
		return nil
	}
	tt, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = HTime{Time: tt}
	return nil
}

func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *HTime) Scan(v interface{}) error {
	switch v := v.(type) {
	case time.Time:
		*t = HTime{Time: v}
	case []byte:
		tt, err := time.Parse(formatTime, string(v))
		if err != nil {
			return err
		}
		*t = HTime{Time: tt}
	case string:
		tt, err := time.Parse(formatTime, v)
		if err != nil {
			return err
		}
		*t = HTime{Time: tt}
	default:
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	return nil
}
