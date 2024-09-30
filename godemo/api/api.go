// Package api provides router for HTTP API.
package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Demo
// @version 1.0
// @description Demo service
// @termsOfService http://swagger.io/terms/

// @contact.name Digital Services / SRE Team
// @contact.url http://www.elisa.fi
// @contact.email devops@elisa.fi

// @license.name Elisa Proprietary License
// @license.url http://www.elisa.fi

type Router struct {
	http.Handler
	spec *json.RawMessage
}

func NewRouter(spec *json.RawMessage) *Router {
	g := gin.New()
	r := &Router{
		Handler: g,
		spec:    spec,
	}

	g.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, "page not found") })
	g.GET("/api/hello", r.getHello)
	g.GET("/api/doc", r.getDoc)
	g.GET("/ready", r.ready)
	return r
}

type HelloResponse struct {
	Msg string `json:"msg" example:"hello"`
}

// getHello godoc
// @Description Get hello message
// @Produce json
// @Success 200 {object} HelloResponse "Hello message"
// @Router /api/hello [get]
func (r *Router) getHello(c *gin.Context) { c.JSON(200, HelloResponse{Msg: "hello"}) }

// getHello godoc
// @Description Get OpenAPI documentation
// @Produce json
// @Success 200
// @Router /api/doc [get]
func (r *Router) getDoc(c *gin.Context) { c.JSON(200, r.spec) }

// ready godoc
// @Description Readiness endpoint
// @Produce json
// @Success 200 {string} string "ready"
// @Router /ready [get]
func (r *Router) ready(c *gin.Context) { c.JSON(200, "ready") }
