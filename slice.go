package gtool

import (
	"math"
	"strings"
)

// IntInSlice 判斷數值是否在 slice 中
func IntInSlice(i int, list []int) bool {
	for _, data := range list {
		if data == i {
			return true
		}
	}
	return false
}

// StringInSlice 判斷字串是否在 slice 中
func StringInSlice(s string, list []string) bool {
	for _, data := range list {
		if data == s {
			return true
		}
	}
	return false
}

// UniqueIntList 排除重複數字
func UniqueIntList(list []int) (uniqueList []int) {
	for _, i := range list {
		if !IntInSlice(i, uniqueList) {
			uniqueList = append(uniqueList, i)
		}
	}
	return
}

// UniqueStrList 排除重複字串
func UniqueStrList(list []string) (uniqueList []string) {
	for _, s := range list {
		if !StringInSlice(s, uniqueList) {
			uniqueList = append(uniqueList, s)
		}
	}
	return
}

// StrInSliceWithContain 判斷字串是否在 slice 中有部份相同的存在
func StrInSliceWithContain(s string, list []string) bool {
	for _, data := range list {
		if strings.Contains(data, s) {
			return true
		}
	}
	return false
}

// SplitIntSlice 指定每組數量，將 []int 切分成多組
func SplitIntSlice(list []int, cutNum int) (listSplit [][]int) {
	if cutNum == 0 {
		// 每組數量 0，直接回傳空
		return
	}

	chap := math.Ceil(float64(len(list)) / float64(cutNum))
	for i := 0; i < int(chap); i++ {
		head := i * cutNum
		tail := (i + 1) * cutNum

		if tail > len(list) {
			tail = len(list)
		}

		newSlice := list[head:tail]
		listSplit = append(listSplit, newSlice)
	}
	return
}

// SplitStringSlice 指定每組數量，將 []string 切分成多組
func SplitStringSlice(list []string, cutNumber int) (listSplit [][]string) {
	if cutNumber == 0 {
		// 每組數量 0，直接回傳空
		return
	}

	chap := math.Ceil(float64(len(list)) / float64(cutNumber))
	for i := 0; i < int(chap); i++ {
		head := i * cutNumber
		tail := (i + 1) * cutNumber

		if tail > len(list) {
			tail = len(list)
		}

		newSlice := list[head:tail]
		listSplit = append(listSplit, newSlice)
	}
	return
}
