package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/xlstudio/wxbizdatacrypt"
)

//6位随机数字
func RandomNumber() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

//解析用户信息
func EncryptDataUser(appid, sessionKey, EncryptedData, Iv string) (interface{}, error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: appid, SessionKey: sessionKey}
	result, err := pc.Decrypt(EncryptedData, Iv, true)
	return result, err
}

//MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//随机字符串
func GetRandomString(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//数组去重
func ArrayUniques(old []string) []string {
	var new_arr []string
	for i := 0; i < len(old); i++ {
		flag := true
		for j := range new_arr {
			if old[i] == new_arr[j] {
				flag = false
				break
			}
		}
		if flag == true {
			new_arr = append(new_arr, old[i])
		}
	}
	return new_arr
}

//截取指定位置字符串
func IndexSubstr(str, index string) string {
	result := strings.Index(str, index)
	prefix := []byte(str)[0:result]
	return string(prefix)
}

//指定获取多少位的float
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

//随机值获取
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
