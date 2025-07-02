package link

import (
	"fmt"
	"net/http"
)
type LinkRepositoryDeps struct{
	router *http.ServeMux
	repository *LinkRepository
}

type linkHandler struct{
	router *http.ServeMux
	repository *LinkRepository
}

func NewLinkHandler(dep *LinkRepositoryDeps){
	l := linkHandler{
		router: dep.router,
		repository: dep.repository,
	}
	l.router.HandleFunc("POST /link",l.addLink())
	l.router.HandleFunc("GET /link/{hash}",l.getLink())
	l.router.HandleFunc("PATH /link/{id}",l.updateLink())
	l.router.HandleFunc("DELET /link/{id}",l.dellLink())
	

}


func (h *linkHandler) addLink() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		h.repository.Create(NewLink("test"))
	}
}
func (h *linkHandler) getLink() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
func (h *linkHandler) dellLink() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
func (h *linkHandler) updateLink() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}