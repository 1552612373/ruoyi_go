package utility

import "github.com/golang-jwt/jwt/v5"

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
var ClaimsName = "Claims"

type UserClaims struct {
	jwt.RegisteredClaims
	UserAgent string
	// 声明自己业务数据
	UserId   int64
	UserName string
}
