package main

import (
	health "github.com/Financial-Times/go-fthealth/v1_1"
	status "github.com/Financial-Times/service-status-go/httphandlers"
	log "github.com/Sirupsen/logrus"
	"github.com/jawher/mow.cli"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const appDescription = "UPP Service that forwards mapped content collections to the content-collection-rw-neo4j. If a 200 answer is received from the writer, it retrieves the elements in the collection from the document-store-api and places them in Kafka on the Post Publication topic so that notifications will be created for them."

func main() {
	app := cli.App("content-collection-unfolder", appDescription)

	appSystemCode := app.String(cli.StringOpt{
		Name:   "app-system-code",
		Value:  "content-collection-unfolder",
		Desc:   "System Code of the application",
		EnvVar: "APP_SYSTEM_CODE",
	})

	appName := app.String(cli.StringOpt{
		Name:   "app-name",
		Value:  "Content Collection Unfolder",
		Desc:   "Application name",
		EnvVar: "APP_NAME",
	})

	port := app.String(cli.StringOpt{
		Name:   "port",
		Value:  "8080",
		Desc:   "Port to listen on",
		EnvVar: "APP_PORT",
	})

	log.SetLevel(log.InfoLevel)
	log.Infof("[Startup] content-collection-unfolder is starting ")

	app.Action = func() {
		log.Infof("System code: %s, App Name: %s, Port: %s", *appSystemCode, *appName, *port)

		go func() {
			serveAdminEndpoints(*appSystemCode, *appName, *port)
		}()

		// todo: insert app code here

		waitForSignal()
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("App could not start, error=[%s]\n", err)
		return
	}
}

func serveAdminEndpoints(appSystemCode string, appName string, port string) {
	healthService := newHealthService(&healthConfig{appSystemCode: appSystemCode, appName: appName, port: port})

	serveMux := http.NewServeMux()

	hc := health.HealthCheck{SystemCode: appSystemCode, Name: appName, Description: appDescription, Checks: healthService.checks}
	serveMux.HandleFunc(healthPath, health.Handler(hc))
	serveMux.HandleFunc(status.GTGPath, status.NewGoodToGoHandler(healthService.gtgCheck))
	serveMux.HandleFunc(status.BuildInfoPath, status.BuildInfoHandler)

	if err := http.ListenAndServe(":"+port, serveMux); err != nil {
		log.Fatalf("Unable to start: %v", err)
	}
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
