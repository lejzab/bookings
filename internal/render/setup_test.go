package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/lejzab/bookings/internal/config"
	"github.com/lejzab/bookings/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// register datatypes stored in session
	gob.Register(models.Reservation{})
	// change this to true when in production
	testApp.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
