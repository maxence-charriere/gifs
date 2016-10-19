# gifs
[![Build Status](https://travis-ci.org/maxence-charriere/gifs.svg?branch=master)](https://travis-ci.org/maxence-charriere/gifs)
[![Go Report Card](https://goreportcard.com/badge/github.com/maxence-charriere/gifs)](https://goreportcard.com/report/github.com/maxence-charriere/gifs)
[![Coverage Status](https://coveralls.io/repos/github/maxence-charriere/gifs/badge.svg?branch=master)](https://coveralls.io/github/maxence-charriere/gifs?branch=master)
[![GoDoc](https://godoc.org/github.com/maxence-charriere/gifs?status.svg)](https://godoc.org/github.com/maxence-charriere/gifs)

Golang package for interact with gifs.com API


## Install
```
go get -u github.com/maxence-charriere/gifs
```

## Documentation
- [godoc](https://godoc.org/github.com/maxence-charriere/gifs)
- [gifs](http://docs.gifs.com/v1.0/docs)

## Init with API Key
```go
gifs.Init("[Your API Key]")
```

## Examples
Import:
```go
func TestImport(t *testing.T) {
	var r gifs.Response
	var err error

	p := gifs.ImportPayload{
		Source: "https://vine.co/v/ibAU6OH2I0K",
		Title:  "2015 Craziness",
		Attribution: gifs.Attribution{
			Site: "vine",
			User: "Maxence",
		},
	}

	if r, err = gifs.Import(p); err != nil {
		t.Fatal(err)
	}

	t.Logf("response: %+v", r)
}
```

Upload:
```go
func TestUpload(t *testing.T) {
	var f *os.File
	var r gifs.Response
	var err error

	if f, err = os.Open("files/test.gif"); err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	p := gifs.UploadPayload{
		File:  f,
		Title: "A super gif",
		Tags: []string{
			"stylish",
			"beautiful",
			"flashy",
		},
	}

	if r, err = gifs.Upload(p); err != nil {
		t.Error(err)
	}

	t.Logf("success: %+v", r.Success)
}
```