package web

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/raynine/goreddit"
)

func NewHandler(store goreddit.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		Store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
	})

	return h
}

type Handler struct {
	*chi.Mux

	Store goreddit.Store
}

const ThreadsListHTML = `
	<h1> Threads </h1>
	<dl>
		{{ range .Threads }}
		<dt> <strong> {{ .Title }} </strong> </dt>
		<dd> {{ .Description }} </dd>
		{{ end }}
	</dl>
`

func (h *Handler) ThreadsList() http.HandlerFunc {

	type data struct {
		Threads []goreddit.Thread
	}

	tpl := template.Must(template.New("").Parse(ThreadsListHTML))

	return func(w http.ResponseWriter, r *http.Request) {
		threads, err := h.Store.GetAllThreads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, data{Threads: threads})
	}
}
