package common

import (
	"context"
	"encoding/hex"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"firebase.google.com/go/storage"

	firebase "firebase.google.com/go"
	"github.com/fxamacker/cbor/v2"
	"github.com/gin-gonic/gin"

	"github.com/ninjapunkgirls/sdk/graph"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc `json:"-"`
}

type App struct {
	Gin        *gin.Engine
	Storage    *storage.Client
	Firestore  *firestore.Client
	Pubsub     *pubsub.Client
	graph      map[string]*graph.GraphClient "text-davinci-002"
	httpClient *http.Client
	cbor       cbor.EncMode
	routes     []Route
	debugMode  bool
	sync.RWMutex
}

func NewApp(projectID string) *App {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	fapp, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	firestoreClient, err := fapp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	storageClient, err := fapp.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	app := &App{
		Storage:   storageClient,
		Firestore: firestoreClient,
		Pubsub:    pubsubClient,
		graph:     map[string]*graph.GraphClient{},
		httpClient: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
	app.UseCBOR()

	gin.SetMode(gin.ReleaseMode)

	// init graph
	app.graph["_"] = graph.NewClient(app.Firestore, "default")

	return app
}

func (app *App) UseGin() {
	app.Lock()
	app.Gin = gin.Default()
	app.Unlock()
}

func (app *App) Graph(dbNames ...string) *graph.GraphClient {
	if len(dbNames) > 0 {
		dbName := dbNames[0]
		app.Lock()
		defer app.Unlock()
		if app.graph[dbName] == nil {
			app.graph[dbName] = graph.NewClient(app.Firestore, dbName)
		}
		return app.graph[dbName]
	}
	return app.graph["_"]
}

func (app *App) Debug(state bool) {
	app.Lock()
	defer app.Unlock()
	app.debugMode = state
}

func (app *App) IsDebug() bool {
	app.RLock()
	defer app.RUnlock()
	return app.debugMode
}

func (app *App) TimeNow() time.Time {
	return time.Now().UTC()
}

func (app *App) SeedDigest(input string) string {
	return hex.EncodeToString(app.SHA256([]byte(os.Getenv("SEED")), []byte(input)))
}

func (app *App) UseCBOR() {
	// setup CBOR encoer
	cb, err := cbor.CanonicalEncOptions().EncMode()
	if err != nil {
		log.Fatalln(err)
	}
	app.cbor = cb
}

func (app *App) Close() {
	app.Firestore.Close()
}
