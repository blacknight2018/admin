package Utils

import (
	"admin/Config"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

const EmptyString = ""

func Integer(d int) *int {
	return &d
}
func GetJSONStringArray(jsonString string) []string {
	var ret []string
	json.Unmarshal([]byte(jsonString), &ret)
	return ret
}

func GetJSONIntArray(jsonString string) []int {
	var ret []int
	json.Unmarshal([]byte(jsonString), &ret)
	return ret
}

func StringArrayToJSON(strArray []string) string {
	if bytes, err := json.Marshal(strArray); err == nil {
		return string(bytes)
	}
	return EmptyString
}

func StrToInt(str string) int {
	var ret int
	ret = 0
	if data, err := strconv.Atoi(str); err == nil {
		ret = data
	}
	return ret
}

type Claims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

/**
 * @Description: 将数据解析成JWTToken
 * @param obj
 * @return string
 */
func GenerateJWT(obj interface{}) string {
	claims := Claims{
		obj,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(148 * time.Hour).Unix()},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString([]byte(Config.GetJWTSecret()))
	return token
}

/**
 * @Description: 将JWTToken解析为对象
 * @param token
 * @return interface{}
 */
func ParseJWT(token string) interface{} {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.GetJWTSecret()), nil
	})
	if tokenClaims != nil && err == nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.Data
		}
	}
	return nil
}

func ContextGetInt(context *gin.Context, key string) int {
	var ret int
	if v, ok := context.Get(key); ok {
		switch data := v.(type) {
		case int:
			ret = data
			break
		case int64:
			ret = int(data)
			break
		case float64:
			ret = int(data)
			break
		}
	}
	return ret
}
func ContextQueryInt(context *gin.Context, key string) int {
	data := context.Query(key)
	ret, _ := strconv.Atoi(data)
	return ret
}
