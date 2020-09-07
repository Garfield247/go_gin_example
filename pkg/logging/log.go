package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F * os.File   //文件
	DefaultPrefix = ""  //默认前缀
	DefaultCallDepth = 2 //调用深度

	logger *log.Logger  //打印
	logPrefix = "" //打印前缀
	levelFlags = []string{"DEBUG","INFO","WARN","ERROR","FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)
/*
iota是golang语言的常量计数器,只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
使用iota能简化定义，在定义枚举时很有用。
*/

func init()  {

	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F,err = openLogFile(fileName,filePath)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(F, DefaultPrefix ,log.LstdFlags)
}
/*
   log.New创建一个新的日志记录器。out定义要写入日志数据的IO句柄。prefix定义每个生成的日志行的开头。flag定义了日志记录属性
   log.LstdFlags：日志记录的格式属性之一，其余的选项如下
   const (
     Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
     Ltime                         // the time in the local time zone: 01:23:23
     Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
     Llongfile                     // full file name and line number: /a/b/c/d.go:23
     Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
     LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
     LstdFlags     = Ldate | Ltime // initial values for the standard logger
   )
*/

func setPrefix(level Level)  {
	_,file,line,ok := runtime.Caller(DefaultCallDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]",levelFlags[level],filepath.Base(file),line)
	}else{
		logPrefix = fmt.Sprintf("[%s]",levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}){
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{})  {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}){
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{})  {
	setPrefix(FATAL)
	logger.Println(v)
}