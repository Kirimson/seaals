package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"seaals/controller"
	"seaals/magick"
	"seaals/models"
	"strings"

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
	return Seal{
		Id:        strings.Split(filepath.Base(seal.Path), ".")[0],
		Permalink: "TODO",
		MimeType:  seal.MimeType,
		CreatedAt: seal.CreatedAt.String(),
		Tags:      tagsToString(seal.Tags),
	}
}

func tagsToString(tags []*models.Tag) []string {
	var tagsString []string
	for _, t := range tags {
		tagsString = append(tagsString, t.Name)
	}
	return tagsString
}

func (s Server) GetApiSeal(ctx *gin.Context, params GetApiSealParams) {
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}
	response := newSealAPIResponse(*seal)
	ctx.JSON(http.StatusOK, response)
	return
}

// Despite advertising this as an ID, we use the path, to make the seal ID look more unique
func (s Server) GetApiSealId(ctx *gin.Context, id string) {
	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}
	response := newSealAPIResponse(*seal)
	ctx.JSON(http.StatusOK, response)
	return
}

func (s Server) GetSeal(ctx *gin.Context, params GetSealParams) {
	// Convert the API Params to what the controller accepts (magick.Opts)
	mo := &magick.Opts{}
	if params.Filter != nil {
		mo.Filter = string(*params.Filter)
	}
	// Set any default options that have not been set by the User
	mo = magick.DefaultOpts(mo)

	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// Create the Seal image
	sealResult, err := s.controller.GetSealImage(seal, mo)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, sealResult.MimeType, sealResult.Image)
}

func (s Server) GetSealId(ctx *gin.Context, id string, params GetSealIdParams) {
	// Convert the API Params to what the controller accepts (magick.Opts)
	mo := &magick.Opts{}
	if params.Filter != nil {
		mo.Filter = string(*params.Filter)
	}
	// Set any default options that have not been set by the User
	mo = magick.DefaultOpts(mo)

	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	fmt.Printf("%+v\n", seal)

	// Create the Seal image
	sealResult, err := s.controller.GetSealImage(seal, mo)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
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

	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// Get the Seal image
	sealResult, err := s.controller.GetSealSaying(seal, text, mo)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, sealResult.MimeType, sealResult.Image)
}

func (s Server) GetSealIdSaysText(ctx *gin.Context, id string, text string, params GetSealIdSaysTextParams) {
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

	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// Get the Seal image
	sealResult, err := s.controller.GetSealSaying(seal, text, mo)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, sealResult.MimeType, sealResult.Image)
}
