package logging

import (
	"log"
	"os"
)

func openLogFile(prePath, filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir(prePath)
	case os.IsPermission(err):
		log.Println("Fail by Permission:", err.Error())
		panic(err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Fail by OpenFile:", err.Error())
		panic(err)
	}

	return handle
}

// mkDir 建立資料夾
func mkDir(path string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+path, os.ModePerm)
	if err != nil {
		log.Println("Fail by MkdirAll:", err.Error())
		panic(err)
	}
}

// checkFileExist 檢查檔案是否存在，如果不存在就重新開檔(會檢查兩次，第一次檢查檔案是否存在，第二次走原本流程)
func checkFileExist(path, filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		openLogFile(path, filePath)
	}
}
