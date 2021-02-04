package gtool

import "time"

// JSONTime 用於解析 Json 的時間格式
type JSONTime time.Time

// jsonTimeFormat json 時間解析格式
var jsonTimeFormat = "2006-01-02 15:04:05"

// UnmarshalJSON 指定時間反序列化格式
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	var now time.Time
	// 不為空字串時，在解析時間資料
	if string(data) != `""` {
		now, err = time.ParseInLocation(`"`+jsonTimeFormat+`"`, string(data), time.UTC)
	}
	*t = JSONTime(now)
	return
}

// MarshalJSON 指定時間序列化格式
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(jsonTimeFormat)+2)
	b = append(b, '"')
	// 時間不為空時 才轉換為時間格式字串
	if !time.Time(t).IsZero() {
		b = time.Time(t).AppendFormat(b, jsonTimeFormat)
	}
	b = append(b, '"')
	return b, nil
}

// String 指定時間自傳格式
func (t JSONTime) String() string {
	return time.Time(t).Format(jsonTimeFormat)
}

// Time 取 time.Time
func (t JSONTime) Time() time.Time {
	return time.Time(t)
}
