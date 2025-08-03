package api

import (
	"errors"
	"net/http"
	"seaals-api/controller"
	"seaals-api/magick"

	"github.com/gin-gonic/gin"
)

//go:generate oapi-codegen --config models.yaml  ../openapi.yaml
//go:generate oapi-codegen --config server.yaml ../openapi.yaml

var _ ServerInterface = (*Server)(nil)

type Server struct {
	controller *controller.SealController
}

func NewSeaalsServer(controller *controller.SealController) *Server {
	return &Server{
		controller: controller,
	}
}

func (s Server) GetSeal(ctx *gin.Context, params GetSealParams) {
	// Convert the API Params to what the controller accepts (magick.Opts)
	mo := &magick.Opts{}
	if params.Filter != nil {
		mo.Filter = string(*params.Filter)
	}
	// Set any default options that have not been set by the User
	mo = magick.DefaultOpts(mo)

	sealResult, err := s.controller.GetSeal(mo)
	if err != nil {
		ctx.Error(errors.New("failed to get Seal image"))
		return
	}

	ctx.Data(http.StatusOK, sealResult.MimeType, sealResult.Image)
}

func (s Server) GetSealSaysText(ctx *gin.Context, text string, params GetSealSaysTextParams) {
	// Convert the API Params to what the controller accepts (magic.Opts)
	mo := &magick.Opts{}
	if params.Filter != nil {
		mo.Filter = string(*params.Filter)
	}
	// Set any defaults not set by the user
	mo = magick.DefaultOpts(mo)

	sealResult, err := s.controller.GetSealSaying(text, mo)
	if err != nil {
		ctx.Error(errors.New("failed to get Seal image"))
		return
	}

	ctx.Data(http.StatusOK, sealResult.MimeType, sealResult.Image)
}
