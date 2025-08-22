package api

import (
	"fmt"
	"net/http"
	"seaals/controller"
	"seaals/magick"
	"seaals/models"
	"slices"
	"strconv"

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

func newSealAPIResponse(seal models.Seal) Seal {
	tags := []Tag{}
	for _, t := range seal.Tags {
		tags = append(tags, t.Name)
	}
	return Seal{
		CreatedAt: seal.CreatedAt.String(),
		Id:        strconv.Itoa(int(seal.ID)),
		MimeType:  seal.MimeType,
		Tags:      tags,
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

	seal, err := s.controller.RandomSeal(*params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal image"))
		return
	}
	if slices.Contains(ctx.Request.Header["Accept"], "application/json") {
		response := newSealAPIResponse(*seal)
		ctx.JSON(http.StatusOK, response)
		return
	}

	// Create the Seal image
	sealResult, err := s.controller.GetSeal(seal, mo)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal image \n%s", err))
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
	if params.Position != nil {
		mo.Position = string(*params.Position)
	}
	if params.FontSize != nil {
		mo.FontSize = float64(*params.FontSize)
	}
	if params.FontColour != nil {
		mo.FontColour = string(*params.FontColour)
	}
	if params.BorderSize != nil {
		mo.BorderSize = float64(*params.BorderSize)
	}
	if params.BorderColour != nil {
		mo.BorderColour = string(*params.BorderColour)
	}
	// Set any defaults not set by the user
	mo = magick.DefaultOpts(mo)

	seal, err := s.controller.RandomSeal(*params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal image"))
		return
	}
	if slices.Contains(ctx.Request.Header["Accept"], "application/json") {
		response := newSealAPIResponse(*seal)
		ctx.JSON(http.StatusOK, response)
		return
	}

	// Get the Seal image
	sealResult, err := s.controller.GetSealSaying(seal, text, mo)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal image \n%s", err))
		return
	}

	ctx.Data(http.StatusOK, sealResult.MimeType, sealResult.Image)
}
