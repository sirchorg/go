package common

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/gin-gonic/gin"
)

func (app *App) newClients() Clients {
	return Clients{
		app: app,
	}
}

type Clients struct {
	app        *App
	gin        *gin.Engine
	httpClient *http.Client
	algolia    *search.Client
	sync.RWMutex
}

// UseGin enables a Gin instance
func (self *Clients) Gin() *gin.Engine {

	self.RLock()
	client := self.gin
	self.RUnlock()

	if client == nil {
		self.Lock()
		gin.SetMode(gin.ReleaseMode)
		self.gin = gin.Default()
		defer self.Unlock()
		return self.gin
	}

	return client
}

func (self *Clients) HTTP() *http.Client {

	self.RLock()
	client := self.httpClient
	self.RUnlock()

	if client == nil {
		self.Lock()
		self.httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
		defer self.Unlock()
		return self.httpClient
	}

	return client
}

func (self *Clients) Algolia() *search.Client {
	return self.algolia
}

/*
func (self *Clients) FirestoreGraph(dbNames ...string) *graph.GraphClient {
	if len(dbNames) > 0 {
		dbName := dbNames[0]
		self.Lock()
		defer self.Unlock()
		if self.graph[dbName] == nil {
			self.graph[dbName] = graph.NewClient(self.firestore, dbName)
		}
		return self.graph[dbName]
	}
	return self.graph["_"]
}
*/
