package core

import (
	"fmt"
	"time"

	"github.com/andycai/wejoin/constant"
	"github.com/andycai/wejoin/library/utils"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

var (
	zoneUTC              = time.UTC
	zone                 = time.FixedZone("CST", 3600)
	validator            = utils.NewValidator()
	lang                 = "en"
	errorList   []string = make([]string, 0)
	messageList []string = make([]string, 0)
)

// IP get remote IP
func IP(c *fiber.Ctx) string {
	return c.IP()
}

func Str(c *fiber.Ctx, key string, defaultValue ...string) string {
	return c.Params(key, defaultValue...)
}

func Int(c *fiber.Ctx, key string, defaultValue ...string) int {
	return cast.ToInt(c.Params(key, defaultValue...))
}

func Uint(c *fiber.Ctx, key string, defaultValue ...string) uint {
	return cast.ToUint(c.Params(key, defaultValue...))
}

func U32(c *fiber.Ctx, key string, defaultValue ...string) uint32 {
	return cast.ToUint32(c.Params(key, defaultValue...))
}

func I32(c *fiber.Ctx, key string, defaultValue ...string) int32 {
	return cast.ToInt32(c.Params(key, defaultValue...))
}

func U64(c *fiber.Ctx, key string, defaultValue ...string) uint64 {
	return cast.ToUint64(c.Params(key, defaultValue...))
}

func I64(c *fiber.Ctx, key string, defaultValue ...string) int64 {
	return cast.ToInt64(c.Params(key, defaultValue...))
}

// Ok 正常响应
func Ok(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code": constant.Success,
		"data": data,
	})
}

// Push 推送响应
func Push(c *fiber.Ctx, code int) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  constant.CodeText(code),
	})
}

func Err(c *fiber.Ctx, err error) error {
	return c.JSON(fiber.Map{
		"error": err.Error(),
	})
}

// Msg push common response
func Msg(c *fiber.Ctx, code int, msg string) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  msg,
	})
}

func HashPassword(password string) string {
	h, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(h)
}

func CheckPassword(password, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(plain))
	return err == nil
}

func HTMXRedirectTo(HXURL string, HXGETURL string, c *fiber.Ctx) error {
	c.Append("HX-Replace-Url", HXURL)
	c.Append("HX-Reswap", "none")

	return Render(c, "components/redirect", fiber.Map{
		"HXGet":     HXGETURL,
		"HXTarget":  "#app-body",
		"HXTrigger": "load",
	}, "layouts/app-htmx")
}

func Render(c *fiber.Ctx, name string, bind interface{}, layouts ...string) error {
	return c.Render(fmt.Sprintf("%s", name), bind, layouts...)
}

func Validate(i interface{}) error {
	return validator.Validate(i)
}

//#region Date, Time, Zone etc

func ParseDate(date string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04", date, zoneUTC)
	if err == nil {
		return t.In(zoneUTC)
	}
	return time.Now().In(zoneUTC)
}

func SetZoneOffset(offset int) {
	zone = time.FixedZone("CST", offset*3600)
}

func DateFormat(t time.Time, layout string) string {
	return t.In(zone).Format(layout)
}

func Now() time.Time {
	return time.Now().In(zone)
}

//#endregion

//#region I18n

func SetLang(l string) {
	lang = l
}

func Lang() string {
	return lang
}

//#endregion

func PushError(err string) {
	errorList = append(errorList, err)
}

func GetErrors() []string {
	list := errorList[0:]
	errorList = []string{}

	return list
}

func PushMessages(msg string) {
	messageList = append(messageList, msg)
}

func GetMessages() []string {
	list := messageList[0:]
	messageList = []string{}

	return list
}
