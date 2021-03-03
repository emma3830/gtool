package gtool

import "time"

// JSONTime 用於解析 json 的時間格 (yyyy-mm-dd hh:ii:ss)
type JSONTime time.Time

// UnmarshalJSON 指定時間反序列化格式
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	var now time.Time
	// 不為空字串時，在解析時間資料
	if string(data) != `""` {
		now, err = time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.UTC)
	}
	*t = JSONTime(now)
	return
}

// MarshalJSON 指定時間序列化格式
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	// 時間不為空時 才轉換為時間格式字串
	if !time.Time(t).IsZero() {
		b = time.Time(t).AppendFormat(b, TimeFormat)
	}
	b = append(b, '"')
	return b, nil
}

// String 指定時間自傳格式
func (t JSONTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// Time 取 time.Time
func (t JSONTime) Time() time.Time {
	return time.Time(t)
}
