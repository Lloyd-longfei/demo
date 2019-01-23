package conf

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const filename = "core"

func ReadConfFile() *viper.Viper {
	path := CurrentPath()
	v := viper.New()
	v.SetConfigName(filename)
	if runtime.GOOS == "windows" {
		v.AddConfigPath("$GOPATH\\" + path + "\\")
	} else {
		v.AddConfigPath("$GOPATH/" + path + "/")
	}
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

func CurrentPath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	result := strings.Index(file, "src")
	prefix := []byte(file)[result:]
	arr := strings.Split(string(prefix), "/")
	str := arr[0] + "/" + arr[1] + "/"
	resultTwo := strings.LastIndex(str, "/")
	currentFile := []byte(file)[result : result+resultTwo]
	return string(currentFile)
}
