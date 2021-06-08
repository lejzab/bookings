package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig hold the configuration of the app
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
