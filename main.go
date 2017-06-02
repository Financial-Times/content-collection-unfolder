package main

import (
	fw "github.com/Financial-Times/content-collection-unfolder/forwarder"
	log "github.com/Sirupsen/logrus"
	"github.com/jawher/mow.cli"
	"net"
	"net/http"
	"os"
	"time"
)

const appDescription = "UPP Service that forwards mapped content collections to the content-collection-rw-neo4j. If a 200 answer is received from the writer, it retrieves the elements in the collection from the document-store-api and places them in Kafka on the Post Publication topic so that notifications will be created for them."

func main() {
	app := cli.App("content-collection-unfolder", appDescription)
	sc := createServiceConfiguration(app)

	log.SetLevel(log.InfoLevel)
	log.Infof("[Startup] content-collection-unfolder is starting ")

	app.Action = func() {
		log.Infof("System code: %s, App Name: %s, Port: %s", sc.appSystemCode, sc.appName, sc.appPort)

		client := setupHttpClient()
		config := &healthConfig{
			client:          client,
			port:            *sc.appPort,
			appSystemCode:   *sc.appSystemCode,
			appName:         *sc.appName,
			appDesc:         appDescription,
			writerHealthUri: *sc.writerHealthURI,
		}

		newRouting(
			newUnfolder(
				fw.NewForwarder(client, *sc.writerURI),
			),
			newHealthService(config),
		).listenAndServe(*sc.appPort)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("App could not start, error=[%v]\n", err)
	}
}

func setupHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConnsPerHost:   20,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}
