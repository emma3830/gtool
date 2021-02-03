package gtool

import (
	"reflect"
	"testing"
)

func TestIntInSlice(t *testing.T) {
	type args struct {
		a    int
		list []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"in slice",
			args{1, []int{1, 2, 3}},
			true,
		},
		{
			"not in slice",
			args{4, []int{1, 2, 3}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("IntInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"in slice",
			args{"a", []string{"a", "b", "c"}},
			true,
		},
		{
			"not in slice",
			args{"d", []string{"a", "b", "c"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueIntList(t *testing.T) {
	type args struct {
		intList []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"排除重複的一個數值",
			args{[]int{1, 2, 2, 3, 1}},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueIntList(tt.args.intList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueIntList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrList(t *testing.T) {
	type args struct {
		strList []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"排除重複的一個字串",
			args{[]string{"a", "b", "b", "c", "a"}},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStrList(tt.args.strList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrInSliceWithContain(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"contain in slice",
			args{"a", []string{"abc", "def"}},
			true,
		},
		{
			"not contain in slice",
			args{"z", []string{"abc", "def"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrInSliceWithContain(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("StrInSliceWithContain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitIntSlice(t *testing.T) {
	type args struct {
		list      []int
		cutNumber int
	}
	tests := []struct {
		name          string
		args          args
		wantListSplit [][]int
	}{
		{
			"slice數量整除切割數量",
			args{[]int{1, 2, 3, 4}, 2},
			[][]int{{1, 2}, {3, 4}},
		},
		{
			"slice數量非整除切割數量",
			args{[]int{1, 2, 3, 4, 5}, 2},
			[][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "切割數量0",
			args: args{[]int{1, 2, 3, 4, 5}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotListSplit := SplitIntSlice(tt.args.list, tt.args.cutNumber); !reflect.DeepEqual(gotListSplit, tt.wantListSplit) {
				t.Errorf("SplitIntSlice() = %v, want %v", gotListSplit, tt.wantListSplit)
			}
		})
	}
}

func TestSplitStringSlice(t *testing.T) {
	type args struct {
		list      []string
		cutNumber int
	}
	tests := []struct {
		name          string
		args          args
		wantListSplit [][]string
	}{
		{
			"slice數量整除切割數量",
			args{[]string{"a", "b", "c", "d"}, 2},
			[][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			"slice數量非整除切割數量",
			args{[]string{"a", "b", "c", "d", "e"}, 2},
			[][]string{{"a", "b"}, {"c", "d"}, {"e"}},
		},
		{
			name: "切割數量0",
			args: args{[]string{"a", "b", "c", "d", "e"}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotListSplit := SplitStringSlice(tt.args.list, tt.args.cutNumber); !reflect.DeepEqual(gotListSplit, tt.wantListSplit) {
				t.Errorf("SplitStringSlice() = %v, want %v", gotListSplit, tt.wantListSplit)
			}
		})
	}
}
