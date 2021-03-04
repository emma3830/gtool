package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level log的層級編號
type Level int

// 層級列表
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// levelFlags 層級名稱
var levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

// 參數設定
const (
	LogFolder  = "logs/"
	LogExt     = "log"
	DateString = "{date}"   // 當下日期的替換文字
	dateFormat = "20060102" // 當下日期格式
)

var (
	rwLock         sync.RWMutex // 避免自訂 log 有名稱重複導致重複開檔，所以所有 log 共用一個鎖
	logFlag        = log.Ldate | log.Lmicroseconds
	CallerDepth    = 2
	ReportingLevel = DEBUG // 寫入何種層級以上的log
	DebugMode      = false // true:直接將log打印至終端機，不寫入檔案
)

// Logger log設定
type Logger struct {
	folder   string // log 資料夾路徑
	filename string // log名稱
	date     string // 目前 log 日期
	F        *os.File
	logger   *log.Logger
}

// GetLogger 取得logger使用
func (o *Logger) GetLogger() *log.Logger {
	return o.logger
}

// Debug 除錯層級
func (o *Logger) Debug(v ...interface{}) {
	o.println(DEBUG, v...)
}

// Info 一般層級
func (o *Logger) Info(v ...interface{}) {
	o.println(INFO, v...)
}

// Warn 警告層級
func (o *Logger) Warn(v ...interface{}) {
	o.println(WARNING, v...)
}

// Error 錯誤層級
func (o *Logger) Error(v ...interface{}) {
	o.println(ERROR, v...)
}

// Fatal 致命層級
func (o *Logger) Fatal(v ...interface{}) {
	o.println(FATAL, v...)
	os.Exit(1)
}

// println 寫入log
func (o *Logger) println(level Level, v ...interface{}) {
	o.Refresh() // 更新檔案名稱 or 路徑
	o.setPrefix(level)
	if level >= ReportingLevel {
		o.logger.Println(v...)
	}
}

// Write 實作 io.Writer 的方法
func (o *Logger) Write(p []byte) (n int, err error) {
	o.Refresh()
	o.logger.Print(string(p))
	return len(p), nil
}

// PrintMemUsage 查看記憶體使用量
func (o *Logger) PrintMemUsage(title string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	msg := fmt.Sprintf(
		"%s: Alloc = %v MiB  TotalAlloc = %v MiB  Sys = %v MiB  NumGC = %v",
		title,
		m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC,
	)
	o.Info(msg)
}

// Refresh 更新 log 檔案名稱 or 位置
func (o *Logger) Refresh() {
	// panic 處理
	defer func() {
		if e := recover(); e != nil {
			log.Println("Log Panic:", e)
			// 無法寫入 log 檔案，則改為直接輸出至畫面上
			o.logger = log.New(os.Stderr, "", logFlag)
		}
	}()

	if DebugMode {
		o.logger = log.New(os.Stderr, "", logFlag)
		return
	}

	// 加上Lock是為了防止同時寫入時重複開檔
	rwLock.Lock()
	defer rwLock.Unlock()

	nowDate := time.Now().Format(dateFormat)
	if o.date != nowDate {
		// 關閉舊log檔案
		o.F.Close()

		// 更新 log 日期
		o.date = nowDate
		// 更新 log 檔案
		prePath, filePath := GetLogFilePath(nowDate, o.folder, o.filename)
		o.F = openLogFile(prePath, filePath)
		o.logger = log.New(o.F, "", logFlag)
	}
}

// setPrefix 設定log前綴
func (o *Logger) setPrefix(level Level) {
	logPrefix := "[" + levelFlags[level] + "]"

	_, file, line, ok := runtime.Caller(CallerDepth)
	if ok {
		t, _ := filepath.Abs(file)
		logPrefix += fmt.Sprintf("[%s:%d]", t, line)
	}

	o.logger.SetPrefix(logPrefix)
}

// GetLogFilePath 取 log 完整路徑
func GetLogFilePath(nowDate, subFolder, filename string) (string, string) {
	// 替換當下日期
	srp := strings.NewReplacer(DateString, nowDate)
	filename = srp.Replace(filename)

	prefixPath := LogFolder
	if subFolder != "" {
		prefixPath += srp.Replace(subFolder) + "/"
	}

	return prefixPath, prefixPath + filename + "." + LogExt
}

// New 建立新的 Logger
func New(folder, filename string) *Logger {
	logger := Logger{folder: folder, filename: filename}
	// 載入檔案
	logger.Refresh()

	// 當 obj GC 時，執行檔案關閉的動作
	runtime.SetFinalizer(&logger, func(logger *Logger) {
		logger.F.Close()
	})

	return &logger
}
