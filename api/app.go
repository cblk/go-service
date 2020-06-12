package api

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	logy "github.com/sirupsen/logrus"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"go_service/api/v1"
	responseV1 "go_service/api/v1/response"
	"go_service/config"
	"gopkg.in/go-playground/validator.v9"
	"os"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

var (
	validatorObj  *validator.Validate
	validatorOnce sync.Once
)

func GetHttpApplication() *gin.Engine {
	cfg := config.GetConfig()
	gin.SetMode(cfg.GetString("gin.mode"))

	engine := gin.New()
	engine.Use(cors.Default())
	engine.Use(gin.LoggerWithWriter(os.Stdout))
	engine.Use(gin.RecoveryWithWriter(os.Stdout))
	engine.Use(APIVersion())

	// Serve static files under static folder
	// for OpenAPI documentations
	engine.Use(static.Serve("/static", static.LocalFile("./static", false)))

	fizzEngine := fizz.NewFromEngine(engine)

	// Do not include package name in component names
	fizzEngine.Generator().UseFullSchemaNames(false)

	// Initialize our own handlers
	tonic.SetBindHook(TonicBindHook)
	tonic.SetErrorHook(TonicResponseErrorHook)
	tonic.SetRenderHook(TonicRenderHook, "")

	// v1 api
	v1.InitRoutes(fizzEngine)

	// Serve OpenAPI specifications
	infos := &openapi.Info{
		Title:       "Go service",
		Description: "A template for Golang API server",
		Version:     "1.0.0",
	}

	fizzEngine.GET("/openapi.json", nil, fizzEngine.OpenAPI(infos, "json"))
	fizzEngine.GET("/openapi.yml", nil, fizzEngine.OpenAPI(infos, "yaml"))

	if len(fizzEngine.Errors()) != 0 {

		for err := range fizzEngine.Errors() {
			logy.Error(err)
		}

		panic("fizz initialization error")
	}

	return engine
}

func APIVersion() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.FullPath()

		re := regexp.MustCompile(`^/v([0-9]+)/`)
		matches := re.FindStringSubmatch(path)

		if len(matches) > 1 {
			c.Set("api_version", matches[1])
		}

		c.Next()
	}
}

// Distribute binding & error handling & render handling to implementations in different API versions

func TonicBindHook(ctx *gin.Context, i interface{}) error {

	if err := ctx.ShouldBind(i); err != nil {
		return err
	}

	// User our own validator to customize messages
	validatorOnce.Do(func() {
		validatorObj = validator.New()
		validatorObj.SetTagName("rules")
		validatorObj.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	})

	if err := validatorObj.Struct(i); err != nil {

		// Handle validation error by ourselves
		// err is of type validator.v9.ValidationErrors
		code, resp := TonicResponseErrorHook(ctx, err)
		TonicRenderHook(ctx, code, resp)

		// We take care of the validation error handling by ourselves
		// So we need to ignore the upcoming handling result by tonic
		return errors.New("must_ignore")
	}

	return nil
}

func TonicResponseErrorHook(ctx *gin.Context, err error) (int, interface{}) {

	if err.Error() == "binding error: must_ignore" {
		// Must be ignored
		return 9999, nil
	}

	apiVersion := ctx.GetString("api_version")

	switch apiVersion {
	case "1":
		return responseV1.TonicErrorResponse(ctx, err)
	default:
		return tonic.DefaultErrorHook(ctx, err)
	}
}

func TonicRenderHook(ctx *gin.Context, statusCode int, payload interface{}) {

	if statusCode == 9999 {
		// Do not render
		return
	}

	apiVersion := ctx.GetString("api_version")

	switch apiVersion {
	case "1":
		responseV1.TonicRenderResponse(ctx, statusCode, payload)
	default:
		tonic.DefaultRenderHook(ctx, statusCode, payload)
	}
}
