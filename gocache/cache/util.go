// cache
package cache

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B  = 1 << (iota * 10) // 1
	KB                    // 1024
	MB                    // 1048576
	GB                    // 1073741824
	TB                    // ...
	PB                    // ...
)

func ParseSize(size string) (int64, string) {
	re, _ := regexp.Compile("[0-9]+")
	// 把数字部分替换成空字符串，这样就得到了单位部分
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)

	var byteNum int64 = 0

	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		byteNum = 0
	}

	if byteNum == 0 {
		log.Println("ParseSize 仅支持 B、KB、MB、GB、TB、PB")
		num = 100
		byteNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}
