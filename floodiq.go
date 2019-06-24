package floodiq

// DefaultHost is where all API requests are handled
const DefaultHost = "https://api.floodiq.com"

// Version is the current client library version
const Version = "0.1-beta"

// Client allows interactions between the Flood iQ API
type Client struct {
	// APIKey is `the` users API Key. Get one at https://floodiq.dev
	APIKey string
}