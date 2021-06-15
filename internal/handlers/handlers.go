package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/lejzab/bookings/internal/config"
	"github.com/lejzab/bookings/internal/models"
	render2 "github.com/lejzab/bookings/internal/render"
	"log"
	"net/http"
)

//Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config2.AppConfig
}

// NewRepo creates new repository for handlers
func NewRepo(appConfig *config2.AppConfig) *Repository {
	return &Repository{
		appConfig,
	}
}

// NewHandlers set new repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render2.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, cześć"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render2.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Reservation renders the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

//Majors is majors suite page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

//Generals renders general's quarter page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

//Contact renderd contact pahe
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Availability checks availability of rooms
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles form
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s, end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends back JSON
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
