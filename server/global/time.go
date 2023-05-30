package global

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time time.Time

func (t *Time) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf(`"%v"`, tTime.Format(DateFormat))), nil
}

func (t *Time) Value() (driver.Value, error) {
	var zTime time.Time
	tlt := time.Time(*t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *Time) Scan(v interface{}) error {
	if val, ok := v.(time.Time); ok {
		*t = Time(val)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
