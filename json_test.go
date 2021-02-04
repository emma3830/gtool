package gtool

import (
	"reflect"
	"testing"
	"time"
)

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		t      time.Time
		args   args
		hasErr bool
	}{
		{
			"正常時間",
			time.Date(2020, time.October, 14, 12, 33, 10, 0, time.UTC),
			args{[]byte(`"2020-10-14 12:33:10"`)},
			false,
		},
		{
			"空字串",
			time.Time{},
			args{[]byte(`""`)},
			false,
		},
		{
			"錯誤時間",
			time.Time{},
			args{[]byte(`"test"`)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := JSONTime{}
			if err := test.UnmarshalJSON(tt.args.data); (err != nil) != tt.hasErr {
				t.Errorf("JSONTime.UnmarshalJSON() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if !reflect.DeepEqual(test.Time(), tt.t) {
				t.Errorf("JSONTime.UnmarshalJSON() JSONTime = %v, wantJSONTime %v", test.Time(), tt.t)
			}
		})
	}
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		t      time.Time
		want   string
		hasErr bool
	}{
		{
			"正常時間",
			time.Date(2020, time.October, 14, 12, 33, 10, 0, time.UTC),
			`"2020-10-14 12:33:10"`,
			false,
		},
		{
			"空字串",
			time.Time{},
			`""`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := JSONTime(tt.t)
			got, err := test.MarshalJSON()
			if (err != nil) != tt.hasErr {
				t.Errorf("JSONTime.MarshalJSON() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("JSONTime.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
