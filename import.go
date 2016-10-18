package gifs

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	importURL = "https://api.gifs.com/media/import"
)

// ImportPayload represent the payload for send when importing a gif.
type ImportPayload struct {
	// The source URL of the media to add. Required.
	Source string `json:"source"`

	// The title of the media.
	Title string `json:"title,omitempty"`

	// The tags relating the the media.
	Tags []string `json:"tags,omitempty"`

	// If the media is not safe for work.
	Nsfw string `json:"nsfw,omitempty"`

	// Any attribution you want displayed.
	Attribution Attribution `json:"attribution"`
}

// Attribution represents the attribution schema required in a ImportPayload.
type Attribution struct {
	// twitter, reddit, instagram, vine, or a custom URL. Required.
	Site string `json:"site"`

	// The username for social media sites.
	User string `json:"user,omitempty"`

	// The url if site != social media.
	URL string `json:"url,omitempty"`
}

// Import imports a gif.
func Import(p ImportPayload) (response Response, err error) {
	var payload []byte
	var req *http.Request
	var res *http.Response

	if payload, err = json.Marshal(p); err != nil {
		return
	}

	log.Print(string(payload))
	r := bytes.NewReader(payload)

	if req, err = http.NewRequest("POST", importURL, r); err != nil {
		return
	}

	req.Header.Add("Gifs-API-Key", key)
	req.Header.Add("Content-Type", "application/json")

	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	defer res.Body.Close()

	d := json.NewDecoder(res.Body)
	d.Decode(&response)

	if res.StatusCode != 200 {
		err = errors.New(response.Errors.Message)
		return
	}

	return
}
