package jsontype

import (
	"database/sql/driver"
	"encoding/gob"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

func init() {
	gob.Register(JsonTime(time.Now()))
}

type JsonTime time.Time

func (t *JsonTime) UnmarshalJSON(b []byte) error {
	var unix int64
	if err := json.Unmarshal(b, &unix); err != nil {
		return err
	}

	*t = JsonTime(time.Unix(unix, 0))
	return nil
}

func Now() JsonTime {
	return JsonTime(time.Now())
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}

func (t JsonTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t JsonTime) GormDataType(gorm.Dialect) string {
	return "TIMESTAMP NULL"
}

func (t *JsonTime) Scan(src interface{}) error {
	var (
		err error
		tm  time.Time
	)
	switch raw := src.(type) {
	case string:
		tm, err = time.Parse("2006-01-02 15:04:05", raw)
	case []byte:
		tm, err = time.Parse("2006-01-02 15:04:05", string(raw))
	case time.Time:
		tm = raw
	case *time.Time:
		tm = *raw
	default:
		err = errors.New("unsupported value for JsonTime to scan")
	}
	if err != nil {
		return err
	}
	*t = JsonTime(tm)
	return nil
}

func (t JsonTime) Value() (driver.Value, error) {
	return driver.Value(t.String()), nil
}

func (t JsonTime) MarshalBinary() ([]byte, error) {
	tm := time.Time(t)
	return tm.MarshalBinary()
}

func (t *JsonTime) UnmarshalBinary(data []byte) error {
	tm := (*time.Time)(t)
	return tm.UnmarshalBinary(data)
}

func (t JsonTime) MarshalText() ([]byte, error) {
	return time.Time(t).MarshalText()
}

func (t *JsonTime) UnmarshalText(data []byte) error {
	return (*time.Time)(t).UnmarshalText(data)
}
