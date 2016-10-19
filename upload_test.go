package gifs

import (
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	var f *os.File
	var r Response
	var err error

	if f, err = os.Open("files/test.gif"); err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	p := UploadPayload{
		File:  f,
		Title: "A super gif",
		Tags: []string{
			"stylish",
			"beautiful",
			"flashy",
		},
	}

	if r, err = Upload(p); err != nil {
		t.Error(err)
	}

	t.Logf("success: %+v", r.Success)
}

func TestUploadError(t *testing.T) {
	var f *os.File
	var r Response
	var err error

	if f, err = os.Open("files/test.jpg"); err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	p := UploadPayload{
		File:  f,
		Title: "A drone",
	}

	if r, err = Upload(p); err == nil {
		t.Error("should error")
	}

	t.Logf("success: %+v", r.errors)
}
