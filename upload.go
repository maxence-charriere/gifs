package gifs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

const (
	uploadURL = "https://api.gifs.com/media/upload"
)

// UploadPayload represents the payload to upload a gif.
type UploadPayload struct {
	// The file being uploaded. Required.
	File io.Reader

	// The title of the media.
	Title string

	// The tags relating the the media.
	Tags []string

	// If the media is not safe for work.
	Nsfw bool

	// twitter, reddit, instagram, or a custom URL.
	AttributionSite string

	// The username for social media site.
	AttributionUser string

	// The url if site != social media.
	AttributionURL string
}

// Upload uploads a gif.
func Upload(p UploadPayload) (response Response, err error) {
	var b bytes.Buffer
	var req *http.Request
	var res *http.Response

	w := multipart.NewWriter(&b)

	if err = prepareMultipartFormData(w, p); err != nil {
		return
	}

	w.Close()

	if req, err = http.NewRequest("POST", uploadURL, &b); err != nil {
		return
	}

	req.Header.Add("Gifs-API-Key", key)
	req.Header.Add("Content-Type", w.FormDataContentType())

	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	defer res.Body.Close()

	d := json.NewDecoder(res.Body)

	if err = d.Decode(&response); err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New(response.errors.Message)
	}

	return
}

func prepareMultipartFormData(w *multipart.Writer, p UploadPayload) (err error) {
	var fileWriter io.Writer

	if fileWriter, err = w.CreateFormFile("file", "gif"); err != nil {
		return
	}

	if _, err = io.Copy(fileWriter, p.File); err != nil {
		return
	}

	if err = w.WriteField("title", p.Title); err != nil {
		return
	}

	if err = w.WriteField("tags", strings.Join(p.Tags, ", ")); err != nil {
		return
	}

	if err = w.WriteField("nsfw", strconv.FormatBool(p.Nsfw)); err != nil {
		return
	}

	if err = w.WriteField("attribution_site", p.AttributionSite); err != nil {
		return
	}

	if err = w.WriteField("attribution_user", p.AttributionUser); err != nil {
		return
	}

	err = w.WriteField("attribution_url", p.AttributionURL)
	return
}
