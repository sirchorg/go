package common

import (
	"log"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/fxamacker/cbor/v2"
)

// UseGCP grants the conditions for the GCP services clients
func (app *App) UseGCP(projectID string) {
	if app.debugMode {
		log.Println("CONFIGURING >> GCP with project ", projectID)
	}
	app.GCPClients.Lock()
	defer app.GCPClients.Unlock()
	app.GCPClients.projectID = projectID
}

// UseCBOR is an efficient encoding package, check it out
func (app *App) UseCBOR() {
	// setup CBOR encoer
	cb, err := cbor.CanonicalEncOptions().EncMode()
	if err != nil {
		log.Fatalln(err)
	}
	app.cbor = cb
}

// UseAlgolia initialises the algolia client
func (app *App) UseAlgolia(appID, secretPath string) {
	secretBytes, err := os.ReadFile(secretPath)
	if err != nil {
		log.Fatal(err)
	}
	app.Clients.algolia = search.NewClient(appID, string(secretBytes))
}

// UseJWT caches a secret signing key in memory
func (app *App) UseJWT(signingKey string) {
	app.Lock()
	defer app.Unlock()
	app.jwtSigningKey = []byte(signingKey)
}
