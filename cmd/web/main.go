package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ashkanamani/booking/pkg/config"
	"github.com/ashkanamani/booking/pkg/handlers"
	"github.com/ashkanamani/booking/pkg/render"
)

const portNumber = ":8000"

var app config.AppConfig

var session *scs.SessionManager
// main is the main applicatoin function
func main() {


	// change it to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	
	app.Session = session
	tc, err := render.CeateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache!", err)
	}
	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := http.Server{
		Handler: routes(&app),
		Addr: portNumber,
	}
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
