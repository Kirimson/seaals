package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwagger(r *chi.Mux) {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), // The url pointing to API definition
		httpSwagger.URL("/doc.json"),
	))
	r.Get("/doc.json", func(w http.ResponseWriter, r *http.Request) {
		spec, err := openapiSpec()
		if err != nil {
			e := SeaalsError{Message: "failed to get API spec", Error: err}
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		w.Write(spec)
	})
}

func openapiSpec() ([]byte, error) {
	s, _ := GetSwagger()
	b, _ := s.MarshalJSON()
	return b, nil
}
