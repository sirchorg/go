package docs

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ninjapunkgirls/sdk/cloudfunc"
)

type Document struct {
	client   *Client
	Bucket   string
	Lat, Lng float64
	Place    Place
	Parent   string
	Class    string
	Data     map[string]interface{}
}

func (self *Document) Serialise() (string, error) {
	s, err := cloudfunc.CompactSerial(self)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (self *Document) ID() string {
	serial, err := self.Serialise()
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(Hash([]byte(serial)))
}

func (self *Document) TimePrefix(t time.Time) string {

	return fmt.Sprintf("%d/%02d/%02d/%02d/%02d/%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}
