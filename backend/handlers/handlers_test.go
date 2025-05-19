package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AB-Rhman/simple-go/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock database
type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetAllTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockDB) CreateTask(task models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockDB) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetTasks(t *testing.T) {
	// Create mock DB
	mockDB := new(MockDB)
	mockDB.On("GetAllTasks").Return([]models.Task{}, nil)

	// Create handler with mock DB
	handler := NewHandler(mockDB)

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetTasks).ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateTask(t *testing.T) {
	// Create mock DB
	mockDB := new(MockDB)
	mockDB.On("CreateTask", mock.AnythingOfType("models.Task")).Return(nil)

	// Create handler with mock DB
	handler := NewHandler(mockDB)

	// Create a test task
	task := models.Task{
		Title: "Test Task",
	}
	body, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.CreateTask).ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestCreateTaskInvalidJSON(t *testing.T) {
	// Create mock DB
	mockDB := new(MockDB)

	// Create handler with mock DB
	handler := NewHandler(mockDB)

	// Create invalid JSON
	body := []byte(`{"invalid json"`)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.CreateTask).ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteTask(t *testing.T) {
	// Create mock DB
	mockDB := new(MockDB)
	mockDB.On("DeleteTask", "1").Return(nil)

	// Create handler with mock DB
	handler := NewHandler(mockDB)

	// Create a request with URL parameters
	req, err := http.NewRequest("DELETE", "/tasks/1", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Create a router and add the request context
	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
