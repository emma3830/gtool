package gtool

import (
	"reflect"
	"testing"
	"time"
)

func TestTime_GetTime(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			"取得自訂時間套件中的time.Time",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.GetTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_GetString(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"取得時間字串",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			"2021-01-01 12:11:22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.GetString(); got != tt.want {
				t.Errorf("Time.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_GetDateString(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"取得日期字串",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			"2021-01-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.GetDateString(); got != tt.want {
				t.Errorf("Time.GetDateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_Yesterday(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *Time
	}{
		{
			"取得昨天時間",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			&Time{t: time.Date(2020, time.December, 31, 12, 11, 22, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.Yesterday(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.Yesterday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_TimeZero(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *Time
	}{
		{
			"取得零點時間",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			&Time{t: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.TimeZero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.TimeZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_TimeOnHour(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *Time
	}{
		{
			"取得整點時間",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			&Time{t: time.Date(2021, time.January, 1, 12, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.TimeOnHour(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.TimeOnHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_TimeEndHour(t *testing.T) {
	type fields struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *Time
	}{
		{
			"取得小時最後時間",
			fields{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			&Time{t: time.Date(2021, time.January, 1, 12, 59, 59, 59, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Time{
				t: tt.fields.t,
			}
			if got := o.TimeEndHour(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.TimeEndHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTime(t *testing.T) {
	type args struct {
		time time.Time
	}
	tests := []struct {
		name string
		args args
		want *Time
	}{
		{
			"取得自定義時間套件",
			args{time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
			&Time{t: time.Date(2021, time.January, 1, 12, 11, 22, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTime(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestESTNow(t *testing.T) {
	t.Run("取得當下美東時間", func(t *testing.T) {
		t.Logf("ESTNow().GetString() = %+v", ESTNow().GetString())
	})
}

func TestTwNow(t *testing.T) {
	t.Run("取得當下台灣時間", func(t *testing.T) {
		t.Logf("TwNow().GetString() = %+v", TwNow().GetString())
	})
}
