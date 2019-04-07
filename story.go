package cyoa

// Story is what we will be using in our code
type Story map[string]Chapter

// Chapter represents the struct that contains our json data
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option organizes the data for the options portion in our chapter
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
