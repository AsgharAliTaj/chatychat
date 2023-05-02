package web

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/asgharalitaj/chatychat"
	"github.com/asgharalitaj/chatychat/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store *postgres.ThreadStore) *handler {
  h  := &handler {
    Mux: chi.NewMux(),
    ThreadStore: store,
  }
  h.Use(middleware.Logger)
  h.Route("/threads", func(r chi.Router) {
    r.Get("/", h.ThreadsList())
    r.Get("/new", h.ThreadCreate())
    r.Post("/", h.ThreadsStore())
    r.Post("/delete/{id}", h.ThreadDelete())
  })
  return h
}

type handler struct {
  *chi.Mux
  *postgres.ThreadStore
}

const ThreadListHtml = `
	<h1>Threads</h1>
		{{range .Threads}}
  	<p>Title: {{.Title}}</p>
  	<p>Description: {{.Description}}</p>
  	<form action="/threads/delete/{{.ID}}" method="POST">
  		<button type="submit">Delete</button>
  	</form>
  		</hr>
		{{end}}
		<a href="/threads/new">Create Thread</a>
`
func(h *handler) ThreadsList() http.HandlerFunc {
  type data struct {
    Threads []chatychat.Thread
  }
  templ := template.Must(template.New("").Parse(ThreadListHtml))
  return func(w http.ResponseWriter, r *http.Request) {
    tt, err := h.GetThreads()
    if err != nil {
      http.Error(w, err.Error() , http.StatusInternalServerError)
      return
    }
    templ.Execute(w, data{Threads: tt})
  }
}

const threadForm = `
<h1>New Thread </h1>
<form action="/threads" method="POST">
<table>
	<tr>
		<td>Title</td>
		<td><input type="text" name="title" /></td>
	</tr>
	<tr>
		<td>Description</td>
		<td><input type="text" name="description" /></td>
	</tr>
</table>
	<button type="submit">Create Thread</button>
</form>
`

func (h *handler) ThreadCreate() http.HandlerFunc {
	tmpl := template.Must(template.New("").Parse(threadForm))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *handler) ThreadsStore() http.HandlerFunc {
	var th chatychat.Thread
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")
		th.Title = title
		th.Description = description
		err := h.CreateThread(th)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
			http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *handler) ThreadDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.DeleteThread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
