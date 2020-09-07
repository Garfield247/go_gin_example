package logging

import (
	"fmt"
	"github.com/Garfield247/go_gin_example/pkg/file"
	"github.com/Garfield247/go_gin_example/pkg/setting"
	"os"
	"time"
)

/*
获取日志文件路径
runtime/logs/
保存文件的根路径/日志文件路径/

*/
func getLogFilePath() string {
	return fmt.Sprintf("%s%s",setting.AppSetting.LogSavePath,setting.AppSetting.LogSaveName)
}

/*
获取日志文件名称
log20200722.log
日志文件名+格式化过后的日期.后缀名
*/
func  getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

//打开日志文件
func openLogFile(filename,filePath string) (*os.File,error) {
	dir,err := os.Getwd() //返回与当前目录对应的根路径名
	if err != nil{
		return nil, fmt.Errorf("os,Getwd err:%v",err)
	}
	src := dir + "/" +filePath //拼接文件名
	perm := file.CheckPermission(src) //检车文件权限
	if perm == true {
		return nil,fmt.Errorf("file.CheckPermission Permission denied src: %s",src)
	}

	err = file.IsNotExistMKDir(src) //如果不存在则新建文件夹
	if err != nil{
		return nil, fmt.Errorf("file。IsNotExistMkDir src: %s , err: %v",src,err)
	}

	f,err := file.Open(src + filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		return nil,fmt.Errorf("Fail to OpenFile： %v",err)
	}
	return f,nil
}