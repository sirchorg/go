package models

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
)

type Organization struct {
	Name             string
	DefaultFrequency int
}

type Feed struct {
	ID        string
	Type      string
	Title     string
	URL       string
	Frequency int
	Checked   int64
}

type FeedItem struct {
	Org    string
	Title  string
	Link   string
	Time   int64
	Scores map[string]float64
}

func (self *FeedItem) ID() string {

	b := [][]byte{
		[]byte(self.Org),
		[]byte(self.Title),
	}

	h := sha1.New()
	h.Write(bytes.Join(b, []byte("-------------")))
	return hex.EncodeToString(h.Sum(nil))
}

type User struct {
	Email string
	Role  string
}
