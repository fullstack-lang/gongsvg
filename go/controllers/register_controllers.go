
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongsvg/go/orm"
)

// genQuery return the name of the column
func genQuery( columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongsvg/go")
	{// insertion point for registrations
		v1.GET("/v1/rects", GetRects)
		v1.GET("/v1/rects/:id", GetRect)
		v1.POST("/v1/rects", PostRect)
		v1.PATCH("/v1/rects/:id", UpdateRect)
		v1.PUT("/v1/rects/:id", UpdateRect)
		v1.DELETE("/v1/rects/:id", DeleteRect)

		v1.GET("/v1/svgs", GetSVGs)
		v1.GET("/v1/svgs/:id", GetSVG)
		v1.POST("/v1/svgs", PostSVG)
		v1.PATCH("/v1/svgs/:id", UpdateSVG)
		v1.PUT("/v1/svgs/:id", UpdateSVG)
		v1.DELETE("/v1/svgs/:id", DeleteSVG)

		v1.GET("/v1/texts", GetTexts)
		v1.GET("/v1/texts/:id", GetText)
		v1.POST("/v1/texts", PostText)
		v1.PATCH("/v1/texts/:id", UpdateText)
		v1.PUT("/v1/texts/:id", UpdateText)
		v1.DELETE("/v1/texts/:id", DeleteText)


		v1.GET("/commitnb", GetLastCommitNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}

