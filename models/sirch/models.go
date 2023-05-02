package models

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

type User struct {
	Email string
	Role  string
}
