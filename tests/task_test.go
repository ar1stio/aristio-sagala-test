package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"aristio-sagala-test/config"
	"aristio-sagala-test/routes"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.ConnectDB()
	code := m.Run()
	config.DB.Close()
	os.Exit(code)
}

func TestCreateTask(t *testing.T) {
	router := routes.SetupRouter()

	// Step 1: Create a new task
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"title":"Test Task","description":"Test Description","due_date":"2024-08-17T00:00:00Z"}`)
	req, err := http.NewRequest("POST", "/tasks", body)
	fmt.Println("err:", err)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Step 2: Extract the task ID from the response
	var createdTask map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &createdTask)
	fmt.Println(createdTask)
	taskID := int64(createdTask["id"].(float64))

	// Step 3: Update the task's title and description
	updatedBody := strings.NewReader(`{"title":"Updated Test Task","description":"Updated Description","due_date":"2024-08-18T00:00:00Z"}`)
	updateReq, _ := http.NewRequest("PUT", "/tasks/"+strconv.FormatInt(taskID, 10), updatedBody)
	updateReq.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, updateReq)

	assert.Equal(t, http.StatusOK, w.Code)

	// Step 4: Update the task's status to 'In Progress'
	statusUpdateBody := strings.NewReader(`{"title":"Updated Test Task","description":"Updated Description","due_date":"2024-08-18T00:00:00Z","status":"In Progress"}`)
	statusUpdateReq, _ := http.NewRequest("PUT", "/tasks/"+strconv.FormatInt(taskID, 10), statusUpdateBody)
	statusUpdateReq.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, statusUpdateReq)

	assert.Equal(t, http.StatusOK, w.Code)

	// Step 5: Verify the status update
	getReq, _ := http.NewRequest("GET", "/tasks/"+strconv.FormatInt(taskID, 10), nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, getReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedTask map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &updatedTask)

	assert.Equal(t, "In Progress", updatedTask["status"].(string))

	delReq, _ := http.NewRequest("DELETE", "/tasks/"+strconv.FormatInt(taskID, 10), nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, delReq)

	assert.Equal(t, http.StatusOK, w.Code)

	getDeleteReq, _ := http.NewRequest("GET", "/tasks/"+strconv.FormatInt(taskID, 10), nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, getDeleteReq)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
