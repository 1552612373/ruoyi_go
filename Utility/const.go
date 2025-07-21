package utility

import "github.com/golang-jwt/jwt/v5"

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
var ClaimsName = "Claims"

type AdminClaims struct {
	jwt.RegisteredClaims
	UserAgent string
	// 声明自己业务数据
	AdminId   int64
	AdminName string
}
