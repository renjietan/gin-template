package logger

import (
	"io"
	"os"
	"path"
	"time"

	"example.com/t/types"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

//type  struct{}
//
//func (myClock) Now() time.Time {
//	now := time.Now()
//	fmt.Printf("[Clock.Now] %v, Location: %s\n", now, now.Location())
//	return now
//}

var GlobalLog *logrus.Logger // 全局logger实例，方便非DI场景使用

// NewLogger 创建并初始化一个logrus.Logger，支持按日期切割
func NewLogger(appConfig *types.AppConfig) (*logrus.Logger, error) {
	// 解析日志级别
	cfg := appConfig.LoggerConfig
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	// 配置rotatelogs（按日期切割）
	//cfg_fileName := path.Base(cfg.FilePath) // 返回 ”app.log“
	//prefix := cfg_fileName[:strings.LastIndex(cfg_fileName, ".")]    // 返回 ”app“
	//postfix := cfg_fileName[strings.LastIndex(cfg_fileName, ".")+1:] // 返回 ”log“
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//fmt.Printf("location is %v\n", loc)
	writer, err := rotatelogs.New(
		path.Join(path.Dir(cfg.FilePath), path.Base(cfg.FilePath)+".%Y-%m-%d.log"),
		rotatelogs.WithLinkName(cfg.FilePath+".log"),                  // 生成软链指向最新日志
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*24*time.Hour), // 保留 多长时间
		rotatelogs.WithRotationTime(cfg.RotationTime),                 // 设置日志轮转周期
		rotatelogs.WithClock(rotatelogs.Local),
	)
	if err != nil {
		return nil, err
	}

	log := logrus.New()
	log.SetLevel(level)
	log.SetOutput(io.MultiWriter(os.Stdout, writer))
	//log.SetFormatter(&logrus.TextFormatter{
	//	FullTimestamp:   true,
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	log.SetFormatter(&LinePerFieldFormatter{prefixed.TextFormatter{
		// 强制开启颜色
		ForceColors: true,
		// 强制禁用颜色（优先级高于 ForceColors）
		DisableColors: false,
		// 完整时间戳
		FullTimestamp: true,
		// 时间戳格式
		TimestampFormat: "2006-01-02 15:04:05",
		// 禁用排序
		DisableSorting: true,
	}})
	GlobalLog = log
	return log, nil
}

func L() *logrus.Logger {
	if GlobalLog == nil {
		panic("logger 暂未初始化")
	}
	return GlobalLog
}

// 暂时废弃
//func BeautifyJsonStr(obj map[string]interface{}) string {
//	rawJSON, e := utility.MapToJsonStr(obj)
//	if e != nil {
//		panic(e.Error())
//	}
//	// 美化
//	var prettyJSON bytes.Buffer
//	err := json.Indent(&prettyJSON, []byte(rawJSON), "", "  ") // 前缀为空，缩进为两个空格
//	if err != nil {
//		panic(err.Error())
//	}
//	return prettyJSON.String()
//}
