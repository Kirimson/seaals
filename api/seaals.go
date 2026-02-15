package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"seaals/controller"
	"seaals/image"
	"seaals/models"
	"strings"
)

//go:generate go tool oapi-codegen --config models.yaml  ../openapi.yaml
//go:generate go tool oapi-codegen --config server.yaml ../openapi.yaml

var _ ServerInterface = (*Server)(nil)

type Server struct {
	controller *controller.SealController
}

type SeaalsError struct {
	Message string `json:"msg"`
	Error   error  `json:"error"`
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

func (s Server) GetApiStats(w http.ResponseWriter, r *http.Request) {
	// Get the Seal count
	sealCount, err := s.controller.CountSeals()
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal count", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}
	popularTags, err := s.controller.GetPopularTags()
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal count", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
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
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (s Server) GetApiSeal(w http.ResponseWriter, r *http.Request, params GetApiSealParams) {
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal record", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}
	response := newSealAPIResponse(*seal)
	response.Permalink = fmt.Sprintf("%s/seal/%s", r.Host, seal.GetApiID())
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// Despite advertising this as an ID, we use the path, to make the seal ID look more unique
func (s Server) GetApiSealId(w http.ResponseWriter, r *http.Request, id string) {
	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal record", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}
	response := newSealAPIResponse(*seal)
	response.Permalink = fmt.Sprintf("%s/seal/%s", r.Host, seal.GetApiID())
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// http://localhost:8080/seal?filter=monochrome&permalink=true
func (s Server) GetSeal(w http.ResponseWriter, r *http.Request, params GetSealParams) {
	// Get a random seal
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal record", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	// If permalink is set 302 to to GetSealIdSaysText
	if params.Permalink != nil && *params.Permalink {
		// Remove permalink from query params
		q := r.URL.Query()
		q.Del("permalink")
		http.Redirect(w, r, fmt.Sprintf("/seal/%s?%s", seal.GetApiID(), q.Encode()), http.StatusMovedPermanently)
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
		e := SeaalsError{Message: "failed to load Seal image", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	// render.Data{Data: sealResult.Image, ContentType: fmt.Sprintf("image/%s", sealResult.MimeType)}.Render(w)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", fmt.Sprintf("image/%s", sealResult.MimeType))
	w.Write(sealResult.Image)
}

func (s Server) GetSealId(w http.ResponseWriter, r *http.Request, id string, params GetSealIdParams) {
	// Convert the API Params to what the controller accepts (image.Opts)
	opts := &image.Opts{}
	if params.Filter != nil {
		opts.Filter = string(*params.Filter)
	}
	// Set any default options that have not been set by the User
	opts = image.DefaultOpts(opts)

	seal, err := s.controller.GetSealByPath(id)
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal record", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	// Create the Seal image
	sealResult, err := s.controller.GetSealImage(seal, opts)
	if err != nil {
		e := SeaalsError{Message: "failed to load Seal image", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", fmt.Sprintf("image/%s", sealResult.MimeType))
	w.Write(sealResult.Image)
}

func (s Server) GetSealSaysText(w http.ResponseWriter, r *http.Request, text string, params GetSealSaysTextParams) {
	// Get a random seal
	seal, err := s.controller.RandomSeal(params.Tag)
	if err != nil {
		e := SeaalsError{Message: "failed to get Seal record", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	// If permalink is set 302 to to GetSealIdSaysText
	if params.Permalink != nil && *params.Permalink {
		// Remove permalink from query params
		q := r.URL.Query()
		q.Del("permalink")
		http.Redirect(w, r, fmt.Sprintf("/seal/%s?%s", seal.GetApiID(), q.Encode()), http.StatusMovedPermanently)
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
		e := SeaalsError{Message: "failed to load Seal image", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", fmt.Sprintf("image/%s", sealResult.MimeType))
	w.Write(sealResult.Image)
}

func (s Server) GetSealIdSaysText(w http.ResponseWriter, r *http.Request, id string, text string, params GetSealIdSaysTextParams) {
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
		e := SeaalsError{Message: "failed to get Seal record", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	// Get the Seal image
	sealResult, err := s.controller.GetSealSaying(seal, text, opts)
	if err != nil {
		e := SeaalsError{Message: "failed to load Seal image", Error: err}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", fmt.Sprintf("image/%s", sealResult.MimeType))
	w.Write(sealResult.Image)
}
