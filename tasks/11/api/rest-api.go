package api

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"http-calendar/internal/db"
	jRender "http-calendar/internal/json"
	"http-calendar/internal/middleware"
	"http-calendar/internal/models"
	"net/http"
)

type Application struct {
	logger *zap.Logger
	db     *db.PostgresRepository
}

func NewApplication(log *zap.Logger, db *db.PostgresRepository) *Application {
	return &Application{
		logger: log,
		db:     db,
	}
}

func (app *Application) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", middleware.Logger(app.Create))
	return mux
}

// Create [POST] - create an event
func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be POST")
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "can't decode json")
		return
	}

	id, err := app.db.Create(event)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "can't create an event")
		return
	}
	event.ID = id
	jRender.JSON(w, r, http.StatusCreated, event)
}

// [POST] update an event

// [POST] delete an event

// [GET] events for day

// [GET] events for week

// [GET] events for month
