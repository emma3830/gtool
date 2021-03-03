package gtool

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

// JSON jsoniter專用的比原生快的套件
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

// JsonRegisterFuzzyDecoders 啟用 json-iterator 的模糊模式
func JsonRegisterFuzzyDecoders() {
	// 模糊模式，可支援 PHP 空object 為 []、型態自動轉換等...
	extra.RegisterFuzzyDecoders()
}

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
