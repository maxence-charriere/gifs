// Package gifs is a package for interact with gifs.com API
// http://docs.gifs.com/v1.0/docs
package gifs

var (
	key string
)

// Init inits the package.
func Init(apiKey string) {
	key = apiKey
}
