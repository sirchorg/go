package models

type Organization struct {
	Name             string
	DefaultFrequency int
}

type Feed struct {
	ID        string
	Type      string
	URL       string
	Frequency int
	Checked   int64
}

type FeedItem struct {
	Title string
	Link  string
	Time  int64
}

type User struct {
	Email string
	Role  string
}
