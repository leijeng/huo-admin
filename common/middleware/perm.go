package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"regexp"
	"strings"
)

func PermHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		rid := c.GetInt("a_rid")
		if rid == -1 {
			c.Next()
			return
		}

		curUri := c.Request.URL.Path
		curMethod := c.Request.Method
		//uid := c.GetInt("a_uid")

		//权限限制先不需要
		//fmt.Println(curUri, curMethod, uid, rid)
		var aid int
		apis := GetApis()
		for _, api := range apis {
			if strings.ToUpper(api.Method) == curMethod && KeyMatch2(curUri, api.Path) {
				aid = api.Id
				c.Set("a_aid", api.Id)
				c.Set("a_atitle", api.Title)
				break
			}
		}

		if aid < 1 {
			Fail(c, 403, "无权限,aid为空")
			return
		}

		if err := service.SerSysMenu.CanAccess(c, aid); err != nil {
			Fail(c, 403, "无权限,获取进入权限是")
			return
		}

		c.Next()
	}
}

var apis []models.SysApi

func GetApis() []models.SysApi {
	if len(apis) == 0 {
		service.SerSysApi.GetByType(0, &apis)
	}
	return apis
}

// KeyMatch2 determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*", "/resource1" matches "/:resource"
func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

// RegexMatch determines whether key1 matches the pattern of key2 in regular expression.
func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}
