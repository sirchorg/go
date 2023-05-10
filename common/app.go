package common

import (
	"encoding/hex"
	"os"
	"sync"
	"time"

	"github.com/fxamacker/cbor/v2"
)

type App struct {
	Clients
	GCPClients
	jwtSigningKey []byte
	cbor          cbor.EncMode
	routes        []Route
	debugMode     bool
	sync.RWMutex
}

func NewApp() *App {

	app := &App{}

	return app
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
