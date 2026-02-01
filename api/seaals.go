package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"seaals/controller"
	"seaals/image"
	"seaals/models"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:generate go tool oapi-codegen --config models.yaml  ../openapi.yaml
//go:generate go tool oapi-codegen --config server.yaml ../openapi.yaml

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

func (s Server) GetApiStats(ctx *gin.Context) {
	// Get the Seal count
	sealCount, err := s.controller.CountSeals()
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal count: %s", err))
		return
	}
	popularTags, err := s.controller.GetPopularTags()
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal count: %s", err))
		return
	}

	var ts []TagStats
	for _, pt := range popularTags {
		ts = append(ts, TagStats{
			Name:  pt.Name,
			Count: int(pt.Count),
		})
	}
	response := Stats{
		Count: int(sealCount),
		Tags:  ts,
	}
	ctx.JSON(http.StatusOK, response)
}

func (s Server) GetApiSeal(ctx *gin.Context, params GetApiSealParams) {
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}
	response := newSealAPIResponse(*seal)
	response.Permalink = fmt.Sprintf("%s/seal/%s", ctx.Request.Host, seal.GetApiID())
	ctx.JSON(http.StatusOK, response)
}

// Despite advertising this as an ID, we use the path, to make the seal ID look more unique
func (s Server) GetApiSealId(ctx *gin.Context, id string) {
	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}
	response := newSealAPIResponse(*seal)
	response.Permalink = fmt.Sprintf("%s/seal/%s", ctx.Request.Host, seal.GetApiID())
	ctx.JSON(http.StatusOK, response)
}

// http://localhost:8080/seal?filter=monochrome&permalink=true
func (s Server) GetSeal(ctx *gin.Context, params GetSealParams) {
	// Get a random seal
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// If permalink is set 302 to to GetSealIdSaysText
	if params.Permalink != nil && *params.Permalink {
		// Remove permalink from query params
		q := ctx.Request.URL.Query()
		q.Del("permalink")
		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/seal/%s?%s", seal.GetApiID(), q.Encode()))
		return
	}

	// Convert the API Params to what the controller accepts (image.Opts)
	opts := &image.Opts{}
	if params.Filter != nil {
		opts.Filter = string(*params.Filter)
	}
	// Set any default options that have not been set by the User
	opts = image.DefaultOpts(opts)

	// Create the Seal image
	sealResult, err := s.controller.GetSealImage(seal, opts)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, fmt.Sprintf("image/%s", sealResult.MimeType), sealResult.Image)
}

func (s Server) GetSealId(ctx *gin.Context, id string, params GetSealIdParams) {
	// Convert the API Params to what the controller accepts (image.Opts)
	opts := &image.Opts{}
	if params.Filter != nil {
		opts.Filter = string(*params.Filter)
	}
	// Set any default options that have not been set by the User
	opts = image.DefaultOpts(opts)

	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// Create the Seal image
	sealResult, err := s.controller.GetSealImage(seal, opts)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, fmt.Sprintf("image/%s", sealResult.MimeType), sealResult.Image)
}

func (s Server) GetSealSaysText(ctx *gin.Context, text string, params GetSealSaysTextParams) {
	// Get a random seal
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// If permalink is set 302 to to GetSealIdSaysText
	if params.Permalink != nil && *params.Permalink {
		// Remove permalink from query params
		q := ctx.Request.URL.Query()
		q.Del("permalink")
		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/seal/%s/says/%s?%s", seal.GetApiID(), text, q.Encode()))
		return
	}

	// Convert the API Params to what the controller accepts (image.Opts)
	opts := &image.Opts{}
	if params.Filter != nil {
		opts.Filter = string(*params.Filter)
	}
	if params.Position != nil {
		opts.Position = image.TextPosition(*params.Position)
	}
	if params.FontSize != nil {
		opts.Size = int(*params.FontSize)
	}
	if params.FontColour != nil {
		opts.Colour = string(*params.FontColour)
	}
	if params.BorderSize != nil {
		opts.StrokeSize = int(*params.BorderSize)
	}
	if params.BorderColour != nil {
		opts.Stroke = string(*params.BorderColour)
	}
	// Set any defaults not set by the user
	opts = image.DefaultOpts(opts)

	// Get the Seal image
	sealResult, err := s.controller.GetSealSaying(seal, text, opts)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, fmt.Sprintf("image/%s", sealResult.MimeType), sealResult.Image)
}

func (s Server) GetSealIdSaysText(ctx *gin.Context, id string, text string, params GetSealIdSaysTextParams) {
	// Convert the API Params to what the controller accepts (image.Opts)
	opts := &image.Opts{}
	if params.Filter != nil {
		opts.Filter = string(*params.Filter)
	}
	if params.Position != nil {
		opts.Position = image.TextPosition(*params.Position)
	}
	if params.FontSize != nil {
		opts.Size = int(*params.FontSize)
	}
	if params.FontColour != nil {
		opts.Colour = string(*params.FontColour)
	}
	if params.BorderSize != nil {
		opts.StrokeSize = int(*params.BorderSize)
	}
	if params.BorderColour != nil {
		opts.Stroke = string(*params.BorderColour)
	}
	// Set any defaults not set by the user
	opts = image.DefaultOpts(opts)

	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get Seal record: %s", err))
		return
	}

	// Get the Seal image
	sealResult, err := s.controller.GetSealSaying(seal, text, opts)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to load Seal image: %s", err))
		return
	}

	ctx.Data(http.StatusOK, fmt.Sprintf("image/%s", sealResult.MimeType), sealResult.Image)
}
