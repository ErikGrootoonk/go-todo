package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/erikdev/go-todo/store"
)

type TodoHandler struct {
	store *store.TodoStore
	tmpl  *template.Template
}

func NewTodoHandler(s *store.TodoStore, tmpl *template.Template) *TodoHandler {
	return &TodoHandler{store: s, tmpl: tmpl}
}

func (h *TodoHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetAll()
	if err != nil {
		http.Error(w, "failed to load todos", http.StatusInternalServerError)
		log.Println("GetAll error:", err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	h.tmpl.Execute(w, todos)
}

func (h *TodoHandler) HandleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	title := strings.TrimSpace(r.FormValue("task"))
	if title != "" {
		if _, err := h.store.Create(title); err != nil {
			log.Println("Create error:", err)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *TodoHandler) HandleToggle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.store.Toggle(id); err != nil {
		log.Println("Toggle error:", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *TodoHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.store.Delete(id); err != nil {
		log.Println("Delete error:", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
