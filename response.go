package gifs

// Response represents a response from the Gifs API
type Response struct {
	Success Success `json:"success"`
	Errors  Errors  `json:"errors"`
}

// Success represents a response success data.
type Success struct {
	Page   string `json:"page"`
	Files  Files  `json:"files"`
	Oembed string `json:"oembed"`
	Embed  string `json:"embed"`
	Meta   Meta   `json:"meta"`
}

// Files represents a success files data.
type Files struct {
	Gif  string `json:"gif"`
	Jpg  string `json:"jpg"`
	Mp4  string `json:"mp4"`
	Webm string `json:"webm"`
}

// Meta represents a success meta data.
type Meta struct {
	Duration string `json:"duration"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

// Errors represents a response errors data.
type Errors struct {
	Message string `json:"message"`
}
