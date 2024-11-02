package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leijeng/huo-admin/common/consts"
	"github.com/leijeng/huo-admin/common/utils"
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-core/core"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	//生成测试 token
	token, err := GenerateJWT(0, "B64McOpQVrDuQaIJcFqVPiYcbYw5DhhkaNJUo4aCDaw3Cg2l1cy3bSDxrfKfdPE2")
	log.Println("token Bearer", token, err)
}

func GenerateJWT(sub int, secretKey string) (string, error) {
	claims := jwt.MapClaims{"sub": sub, "exp": time.Now().Add(time.Hour * 7 * 24).Unix()}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

func GenerateJWT2(sub int, secretKey string) (string, error) {
	claims := jwt.MapClaims{"exp": time.Now().Add(time.Hour * 7 * 24).Unix()}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

func AdminJwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		//检查是否退出登录
		isLogout := utils.CheckLogout(c, authorization)
		if isLogout != "" {
			Fail(c, 403, "you logout")
			return
		}

		core.Log.Debug("JwtHandler", zap.Any("Authorization", authorization))
		accessToken, err := GetAccessToken(authorization)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}

		// 解析Token
		claims := jwt.MapClaims{}
		//err = ParseStr(accessToken, core.Cfg.JWT.SignKey)
		if core.Cfg.Server.Mode == "test" {
			err = Parse(accessToken, claims, core.Cfg.JWT.SignKey, jwt.WithSubject(core.Cfg.JWT.Subject), jwt.WithoutClaimsValidation())
		} else {
			err = Parse(accessToken, claims, core.Cfg.JWT.SignKey, jwt.WithSubject(core.Cfg.JWT.Subject))
		}
		if err != nil {
			core.Log.Error("JwtHandler.Parse", zap.Any("err", err))
			Fail(c, 401, err.Error())
			return
		}

		fmt.Printf("claims ==== %+v", claims)
		sub, ok := claims["uid"]
		if ok {
			uid, _ := sub.(float64)
			c.Set("uid", int(uid))
			c.Set("a_uid", int(uid))
		}

		c.Next()
	}
	//return JwtHandler()
}

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		accessToken, err := GetAccessToken(authorization)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}

		typJWT := utils.TypJWT{
			SigningKey: []byte(core.Cfg.JWT.SignKey),
		}

		customClaims, err := typJWT.ParseTymonToken(accessToken)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}

		if customClaims == nil {
			Fail(c, 401, "unknow error")
			return
		}
		if customClaims.UserId <= 0 {
			customClaims.UserId = int(customClaims.Subject)
		}

		if customClaims.UserId != 0 {
			isLogin := GetAdminUser(customClaims.UserId)
			if isLogin != 1 {
				Fail(c, 401, "noAuth")
				return
			}
		}

		// 刷新时间大于0则判断剩余时间小于刷新时间时刷新Token并在Response header中返回
		if core.Cfg.JWT.Refresh > 0 {
			now := time.Now()
			diff := time.Unix(customClaims.ExpiresAt, 0).Sub(now)
			refreshTTL := time.Duration(core.Cfg.JWT.Refresh) * time.Minute
			//fmt.Println(diff.Seconds(), refreshTTL)
			if diff < refreshTTL {
				newToken, err := typJWT.RefreshTymonToken(accessToken)
				if err == nil {
					c.Writer.Header().Set("refresh-access-token", newToken)
					//c.Writer.Header().Set("refresh-exp", strconv.FormatInt(exp.Unix(), 10))
				}
			}
		}

		//core.Log.Info("loginMsg", zap.Any("info", utils.StructToJsonString(customClaims)))

		c.Set("a_uid", customClaims.UserId)
		c.Set("a_rid", customClaims.RoleId)
		c.Set("a_userid", customClaims.UserId)
		c.Set("a_roleid", customClaims.RoleId)
		c.Set("a_mobile", customClaims.Phone)
		c.Set("a_nickname", customClaims.Nickname)
		c.Set("a_username", customClaims.Username)
		log.Println("------------jwt------------", fmt.Sprintf("%#v", customClaims))
		//c.Set("jwt_data", customClaims.JwtData)
		c.Next()
	}
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// Refresh 刷新JWT Token
func Refresh(claims jwt.Claims, secretKey string) (string, error) {
	return utils.Generate(claims, secretKey)
}

// Parse 解析token
func Parse(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), err
	}, options...)
	if err != nil {
		return err
	}

	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return nil
	}

	return errors.New("Invalid Token")
}

// GetAccessToken 获取jwt的Token
func GetAccessToken(authorization string) (accessToken string, err error) {
	if authorization == "" {
		return "", errors.New("authorization header is missing")
	}

	// 检查 Authorization 头的格式
	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}

	// 提取 Token 的值
	accessToken = strings.TrimPrefix(authorization, "Bearer ")
	// 去除空格
	accessToken = strings.TrimSpace(accessToken)
	return
}

func GetAdminUser(id int) int {

	cacheKey := consts.AdminLoginUserKey + "_" + strconv.Itoa(id)
	client := utils.InitRedis()

	isLogin := client.Get(context.Background(), cacheKey)

	if isLogin.Val() != "" {

		num, _ := strconv.Atoi(isLogin.Val())

		return num
	}
	var userData models.SysUser
	//_ = service.SerSysUser.DB().Model(&models.SysUser{}).Select("is_login").Where("id = ?", id).First(&userData).Error
	_ = service.SerSysUser.DB().Model(&models.SysUser{}).Select("is_login").Where("id = ?", id).First(&userData).Error

	_ = client.Set(context.Background(), cacheKey, userData.IsLogin, 10*time.Minute)
	return userData.IsLogin
}
