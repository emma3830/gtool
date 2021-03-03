package gtool

import (
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// Time 自訂時間套件
type Time struct {
	t time.Time
}

// GetTime 取得 Gettime.Time 型態
func (o *Time) GetTime() time.Time {
	return o.t
}

// GetString 回傳格式 yyyy-mm-dd hh:ii:ss
func (o *Time) GetString() string {
	return o.t.Format(TimeFormat)
}

// GetDateString 回傳格式 yyyy-mm-dd
func (o *Time) GetDateString() string {
	return o.t.Format(DateFormat)
}

// Yesterday 取得昨天時間
func (o *Time) Yesterday() *Time {
	o.t = o.t.AddDate(0, 0, -1)
	return o
}

// TimeZero 取得零點時間
func (o *Time) TimeZero() *Time {
	o.t = time.Date(o.t.Year(), o.t.Month(), o.t.Day(), 0, 0, 0, 0, o.t.Location())
	return o
}

// TimeOnHour 取得整點時間
func (o *Time) TimeOnHour() *Time {
	o.t = time.Date(o.t.Year(), o.t.Month(), o.t.Day(), o.t.Hour(), 0, 0, 0, o.t.Location())
	return o
}

// TimeEndHour 取得該小時結束時間 (分、秒、毫秒為59)
func (o *Time) TimeEndHour() *Time {
	o.t = time.Date(o.t.Year(), o.t.Month(), o.t.Day(), o.t.Hour(), 59, 59, 59, o.t.Location())
	return o
}

// NewTime 取得 Time 自訂時間套件
func NewTime(time time.Time) *Time {
	return &Time{t: time}
}

// ESTNow 取得目前美東時間
func ESTNow() *Time {
	now := time.Now().UTC().Add(-4 * time.Hour)
	return &Time{t: now}
}

// TwNow 取得目前台灣時間
func TwNow() *Time {
	now := time.Now().UTC().Add(8 * time.Hour)
	return &Time{t: now}
}
