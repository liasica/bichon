package router

import (
    "encoding/json"
    "github.com/chatpuppy/puppychat/assets"
    "github.com/getkin/kin-openapi/openapi2"
    "github.com/getkin/kin-openapi/openapi2conv"
    jsoniter "github.com/json-iterator/go"
    "github.com/labstack/echo/v4"
    echoSwagger "github.com/swaggo/echo-swagger"
    "net/http"
)

// @title                ChatPuppy Api Doc
// @version              1.0
// @BasePath             /
// @description.markdown
func loadDocRoutes(r *echo.Echo) {
    g := r.Group("/docs")

    // docs.SwaggerInfo.Host = ar.Config.App.Host

    g.GET("", func(c echo.Context) error {
        return c.HTML(200, assets.SwaggerRedocUI)
    })

    g.GET("/swagger.json", func(c echo.Context) error {
        return c.JSONBlob(200, assets.SwaggerSpec)
    })

    g.GET("/swagger/*", echoSwagger.WrapHandler)

    g.GET("/oai3.json", func(c echo.Context) (err error) {
        var doc2 openapi2.T
        if err = json.Unmarshal(assets.SwaggerSpec, &doc2); err != nil {
            return
        }
        doc, err := openapi2conv.ToV3(&doc2)
        b, _ := jsoniter.Marshal(doc)
        return c.Blob(200, "application/json", b)
    })

    g.GET("/docs/octicons.css", func(c echo.Context) error {
        return c.Blob(http.StatusOK, "text/css; charset=utf-8", assets.OcticonsCss)
    })
}
