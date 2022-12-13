package swaggerui

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func SwaggerHandle(r *chi.Mux, conf *Config) {
	fs := http.FileServer(http.Dir(conf.StaticRoot))
	r.Handle(fmt.Sprintf("/%s", conf.URLPatch), http.RedirectHandler(fmt.Sprintf("/%s/", conf.URLPatch), http.StatusMovedPermanently))
	r.Handle(fmt.Sprintf("/%s/*", conf.URLPatch), http.StripPrefix(fmt.Sprintf("/%s", conf.URLPatch), fs))
}
