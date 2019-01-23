package log

import (
	"bufio"
	"demo/conf"
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

type LoggerParamter struct {
	LogPath      string
	LogFileName  string
	MaxAge       time.Duration
	RotationTime time.Duration
}

func (p *LoggerParamter) ConfigLocalFilesystemLogger() {
	conf := conf.ReadConfFile()
	baseLogPath := path.Join(p.LogPath, p.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(p.MaxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(p.RotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", err)
	}
	switch level := conf.GetString("debug"); level {
	case "true":
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stderr)
	default:
		p.setNull()
	}
	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			log.DebugLevel: writer, // 为不同级别设置不同的输出目的
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
			log.PanicLevel: writer,
		}, &log.TextFormatter{})
	log.AddHook(lfHook)
}
func (p *LoggerParamter) setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

func GenerateOrOutPutLogger() {
	handler := LoggerParamter{
		LogPath:      "./log",
		LogFileName:  "lg",
		MaxAge:       time.Hour * 24 * 365,
		RotationTime: time.Hour * 24,
	}
	handler.ConfigLocalFilesystemLogger()
}
