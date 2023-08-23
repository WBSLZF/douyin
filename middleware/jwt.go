package middleware

import (
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserId             int64
	jwt.StandardClaims //jwt的预定义声明
}

var secret_key = []byte("YYDS")

func ReleaseToken(userLogin model.UserLogin) (string, error) {

	claims := Claims{
		UserId: userLogin.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(0 * time.Hour).Unix(), // token的有效时间为24小时
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt_token, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return jwt_token, nil
}

// ParseToken
func ParseToken(tokenString string) (*Claims, bool) {
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil { //解析失败
		return nil, false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { //断言其是否为声明，且token有效
		return claims, true
	}
	return nil, false
}

// 鉴权中间件

func JWTMiddleWare(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		token = ctx.PostForm("token")
	}

	if token == "" {
		ctx.JSON(http.StatusOK, model.Response{
			StatusCode: 401,
			StatusMsg:  "token不存在",
		})
		ctx.Abort()
		return
	}
	// 解析token，一个是解析是否成功，一个是token是否过期了
	parse_token, ok := ParseToken(token)
	if !ok {
		ctx.JSON(http.StatusOK, model.Response{
			StatusCode: 401,
			StatusMsg:  "token解析失败",
		})
		ctx.Abort()
		return
	}
	if time.Now().Unix() > parse_token.ExpiresAt {
		ctx.JSON(http.StatusOK, model.Response{
			StatusCode: 401,
			StatusMsg:  "token已经过期",
		})
		ctx.Abort()
		return
	}
	ctx.Set("user_id", parse_token.UserId)
	ctx.Next()

}
