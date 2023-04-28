package docs

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/sirchorg/go/common"
	"golang.org/x/crypto/sha3"
)

var app *common.App

func init() {
	app = &common.App{}
	app.UseCBOR()
}

func EmptyDocument() *Document {
	doc := &Document{}
	return doc
}

func NewDocument(bucketID string, class string, data map[string]interface{}) *Document {
	doc := &Document{
		Bucket: bucketID,
		Class:  class,
		Data:   data,
	}
	return doc
}

type Document struct {
	Bucket   string
	Lat, Lng float64
	Place    []string
	Parent   string
	Class    string
	Data     map[string]interface{}
}

func (self *Document) ID() string {
	serial, err := app.MarshalCBOR(self)
	if err != nil {
		panic(err)
	}
	h := sha3.New224()
	h.Write([]byte(serial))
	return hex.EncodeToString(h.Sum(nil))
}

func (self *Document) TimePrefix(t time.Time) string {

	return fmt.Sprintf("%d/%02d/%02d/%02d/%02d/%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}
