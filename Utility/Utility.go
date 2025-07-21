package utility

import (
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ThrowSysErrowIfneeded(ctx *gin.Context, err error) {
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrNotFound,
			"msg":  rescode.ErrNotFound.String(),
		})
		return
	}
	if strings.HasPrefix(err.Error(), "ZT") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rescode.ErrCustomize,
			"msg":  strings.TrimPrefix(err.Error(), "ZT"),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInternalServer,
			"msg":  rescode.ErrInternalServer.String(),
		})
	}
}

var DefaultTimeFormat = "2006-01-02 15:04:05"

// 13位时间戳转字符串
func FormatTimestamp(layout string, millis int64) string {
	if millis == 0 {
		return ""
	}
	return time.UnixMilli(millis).Format(layout)
}

// ParseToTimestamp 将指定日期时间字符串转为13位时间戳（毫秒）
func ParseToTimestamp(layout, datetime string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") // 设置时区
	t, err := time.ParseInLocation(layout, datetime, loc)
	if err != nil {
		return 0
	}
	return t.UnixMilli()
}

// IsChineseMobile 判断是否是国内手机号（不严格）
func IsChineseMobile(mobile string) bool {
	// 匹配 11 位手机号，以 13、14、15、17、18、19 开头
	regex := `^1[345789]\d{9}$`
	reg := regexp.MustCompile(regex)
	return reg.MatchString(mobile)
}

// IsValidEmail 判断邮箱是否合法
func IsValidEmail(email string) bool {
	// 简单匹配邮箱格式，支持常见格式即可
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	reg := regexp.MustCompile(regex)
	return reg.MatchString(email)
}

func ConvertIntArrToString(intArr []int64) string {
	// 转成字符串切片：["1", "2", "3"]
	var teacherStrs []string
	for _, id := range intArr {
		teacherStrs = append(teacherStrs, strconv.FormatInt(id, 10))
	}

	// 拼接成逗号分隔的字符串："1,2,3"
	value := strings.Join(teacherStrs, ",")
	return value
}
