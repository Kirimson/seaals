package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(r *gin.Engine) {
	r.GET("/swagger/doc.json", openapiSpec)

	r.GET("swagger/index.html", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))
}

func openapiSpec(c *gin.Context) {
	swagger, err := GetSwagger()
	if err != nil {
		log.Fatalf("error loading swagger spec\n %s", err)
	}
	b, err := swagger.MarshalJSON()
	if err != nil {
		log.Fatalf("error marshalling swagger spec to JSON\n %s", err)
	}
	c.Data(http.StatusOK, "application/octet-stream", b)
}
