package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AB-Rhman/simple-go/models"
	"github.com/gorilla/mux"
)

// DB interface defines the database operations
type DB interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(task models.Task) error
	DeleteTask(id string) error
}

// Handler struct holds dependencies
type Handler struct {
	db DB
}

// NewHandler creates a new handler with dependencies
func NewHandler(db DB) *Handler {
	return &Handler{db: db}
}

// GetTasks handles GET /tasks
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.db.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask handles POST /tasks
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// DeleteTask handles DELETE /tasks/{id}
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.db.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
