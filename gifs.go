// Package gifs is a package to interact with Gifs.com API.
package gifs

var (
	key string
)

// Init inits the package.
func Init(apiKey string) {
	key = apiKey
}
