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
	"strconv"
	"time"
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
	mux.HandleFunc("/", middleware.Logger(app.IdleHandle))
	mux.HandleFunc("/update_event", middleware.Logger(app.UpdateEvent))
	mux.HandleFunc("/delete_event", middleware.Logger(app.DeleteEvent))
	mux.HandleFunc("/events_for_day", middleware.Logger(app.GetForDay))
	mux.HandleFunc("/events_for_week", middleware.Logger(app.GetForWeek))
	mux.HandleFunc("/events_for_month", middleware.Logger(app.GetForMonth))

	return mux
}

// IdleHandle [GET] - test
func (app *Application) IdleHandle(w http.ResponseWriter, r *http.Request) {
	jRender.JSON(w, r, http.StatusOK, "im a http-server for WB L2 :dev03")
}

// Create [POST] - create an event
func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be POST")
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant decode json")
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

// UpdateEvent [POST] update information an event (title or date)
func (app *Application) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be POST")
		return
	}

	var event models.Event
	id := r.URL.Query().Get("id")
	nID, err := strconv.Atoi(id)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "ca")
	}
	event.ID = int64(nID)
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant decode json")
		return
	}

	err = app.db.Update(event.ID, event.Title, event.Date)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant update event in db")
		return
	}
	jRender.JSON(w, r, http.StatusOK, event)
}

// DeleteEvent [POST] delete an event from db
func (app *Application) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be POST")
		return
	}
	id := r.URL.Query().Get("id")
	nID, err := strconv.Atoi(id)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "ca")
	}
	if err := app.db.Delete(nID); err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant delete event from db")
		return
	}
	jRender.JSON(w, r, http.StatusOK, "event successfully deleted!")
}

// GetForDay [GET] events for day
func (app *Application) GetForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be GET")
		return
	}
	dateString := r.URL.Query().Get("date")
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant convert param to date")
		return
	}
	result, err := app.db.GetForDay(date)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant get data from db")
		return
	}
	fmt.Println("HERRE")
	if len(result) == 0 {
		jRender.NoContentJSON(w, r)
		return
	}
	jRender.JSON(w, r, http.StatusOK, result)
}

// GetForWeek [GET] events for week
func (app *Application) GetForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be GET")
		return
	}
	dateString := r.URL.Query().Get("date")
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant convert param to date")
		return
	}

	result, err := app.db.GetForWeek(date)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant get data from db")
		return
	}

	if len(result) == 0 {
		jRender.NoContentJSON(w, r)
		return
	}
	jRender.JSON(w, r, http.StatusOK, result)
}

// GetForMonth [GET] events for month
func (app *Application) GetForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jRender.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be GET")
		return
	}
	dateString := r.URL.Query().Get("date")
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant convert param to date")
		return
	}

	result, err := app.db.GetForMonth(date)
	if err != nil {
		jRender.ErrorJSON(w, r, http.StatusInternalServerError, err, "cant get data from db")
		return
	}

	if len(result) == 0 {
		jRender.NoContentJSON(w, r)
		return
	}
	jRender.JSON(w, r, http.StatusOK, result)
}
