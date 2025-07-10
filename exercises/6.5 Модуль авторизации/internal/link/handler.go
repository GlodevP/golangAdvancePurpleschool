package link

import (
	"log"
	"net/http"
	"strconv"
	"temp/pkg/middleware"
	"temp/pkg/request"
	"temp/pkg/response"

	"gorm.io/gorm"
)

type LinkRepositoryDeps struct {
	Router     *http.ServeMux
	Repository *LinkRepository
}

type linkHandler struct {
	router     *http.ServeMux
	repository *LinkRepository
}

func NewLinkHandler(dep *LinkRepositoryDeps) {
	l := linkHandler{
		router:     dep.Router,
		repository: dep.Repository,
	}
	l.router.HandleFunc("POST /link", l.addLink())
	l.router.HandleFunc("GET /link/{hash}", l.goToLink())
	l.router.Handle("PATCH /link/{id}", middleware.IsAuthed(l.updateLink()))
	l.router.HandleFunc("DELET /link/{id}", l.dellLink())

}

func (h *linkHandler) addLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[CreateRequest](&w, r)
		if err != nil {
			log.Println(err)
			return
		}
		link := NewLink(body.Url)
		for {
			existLink, _ := h.repository.GetByHash(link.Hash)
			if existLink == nil {
				break
			}
			link.GenerateHash()
		}
		l, err := h.repository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, l, http.StatusCreated)
	}
}
func (h *linkHandler) goToLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == "" {
			http.Error(w, "not valid id", http.StatusBadRequest)
			return
		}
		l, err := h.repository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, l.Url, http.StatusTemporaryRedirect)
	}
}
func (h *linkHandler) dellLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		if idString == "" {
			http.Error(w, "not valid id", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "not valid id", http.StatusBadRequest)
			return
		}
		_, err = h.repository.GetByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.repository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, http.StatusOK)
	}
}
func (h *linkHandler) updateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := request.HandleBody[UpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		if idString == "" {
			http.Error(w, "not valid id", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "not valid id", http.StatusBadRequest)
			return
		}
		link, err := h.repository.Update(&Link{
			Model: gorm.Model{
				ID: uint(id),
			},
			Url:  b.Url,
			Hash: b.Hash,
		})
		if err != nil {
			http.Error(w, "not valid id", http.StatusBadRequest)
			return
		}
		response.Json(w, link, http.StatusCreated)
	}
}
