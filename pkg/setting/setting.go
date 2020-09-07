package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize   int
	JwtSecret string
	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}

type Server struct {
	HttpPort int
	RunMode string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}
var DatabaseSetting  = &Database{}

func init()  {
	log.Println("初始化配置中")
	Cfg,err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("获取ini文件失败：%v",err)
	}else{
		log.Print("加载文件成功")
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("AppSetting映射失败：%v",err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting)
		if err != nil {
		log.Fatalf("ServerSetting映射失败：%v",err)
	}
	ServerSetting.ReadTimeout = time.Duration(ServerSetting.ReadTimeout) * time.Second
	ServerSetting.WriteTimeout = time.Duration(ServerSetting.WriteTimeout) * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
		if err != nil {
		log.Fatalf("DatabaseSetting映射失败：%v",err)
	}
}