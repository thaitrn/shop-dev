package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("/Users/jthaitran/Projects/shop-dev/internal/error/en.json")
	bundle.LoadMessageFile("/Users/jthaitran/Projects/shop-dev/internal/error/es.json")
	bundle.LoadMessageFile("/Users/jthaitran/Projects/shop-dev/internal/error/vi.json")
}
func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(LanguageMiddleware())
	r.Use(ErrorHandler())

	r.GET("/example", func(c *gin.Context) {
		err := &AppError{
			ErrorCode: "20003",
			MessageID: "not_found",
			Status:    http.StatusNotFound,
		}
		c.Error(err)
	})
	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

type AppError struct {
	ErrorCode string
	MessageID string
	Status    int
}

func (e *AppError) Error() string {
	return e.MessageID
}
func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "vi"
		}
		c.Set("lang", lang)
		c.Next()
	}
}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if appErr, ok := err.(*AppError); ok {
				lang := c.GetString("lang")
				localizer := i18n.NewLocalizer(bundle, lang)
				message, _ := localizer.Localize(&i18n.LocalizeConfig{
					MessageID: appErr.MessageID,
				})
				c.JSON(appErr.Status, gin.H{"errorCode": appErr.ErrorCode, "errorMessage": message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			c.Abort()
		}
	}
}
