package middleware

import (
	utility "go_ruoyi_base/ruoyi_go/Utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJWTMiddlewareBuilder struct {
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if strings.Contains(path, "/login") || strings.Contains(path, "/signup") || strings.HasPrefix(path, "/swagger") {
			// 不用token
			return
		}
		// 根据约定，token 在 Authorization 头部
		// Bearer XXXX
		authCode := ctx.GetHeader("Authorization")
		// authCode := ctx.GetHeader("zt-jwt-token")
		if authCode == "" {
			// 没登录，没有 token, Authorization 这个头部都没有
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 && authCode != "" {
			// 没登录，Authorization 中的内容是乱传的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		var ac utility.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &ac, func(token *jwt.Token) (interface{}, error) {
			return utility.JWTKey, nil
		})
		if err != nil {
			// token 不对，token 是伪造的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// if ac.UserAgent != ctx.GetHeader("User-Agent") {
		// 	// 后期我们讲到了监控告警的时候，这个地方要埋点
		// 	// 能够进来这个分支的，大概率是攻击者
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// expireTime := ac.ExpiresAt
		// 不判定都可以
		//if expireTime.Before(time.Now()) {
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		// 剩余过期时间 < 50s 就要刷新
		// if expireTime.Sub(time.Now()) < time.Second*50 {
		// 	ac.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
		// 	tokenStr, err = token.SignedString(web.JWTKey)
		// 	ctx.Header("zt-jwt-token", tokenStr)
		// 	if err != nil {
		// 		// 这边不要中断，因为仅仅是过期时间没有刷新，但是用户是登录了的
		// 		log.Println(err)
		// 	}
		// }
		ctx.Set(utility.ClaimsName, ac)
	}
}
