# Go Flood iQ

**WARNING**: This library is beta and subject to change. It's not ready to be used in production quite yet.

---

This is the official implementation of the Flood iQ API in Go. Please use this client if you are using Go to interact with the Flood iQ API.

For more in-depth guides, usage and API access, please see the documentation at docs.floodiq.dev.

## Installation

First install floodiq-go:

```
go get -u github.com/firststreet/floodiq-go
```

Next, import it into your project:

```go
import (
  "github.com/firststreet/floodiq-go
)
```

## Documentation

**In order to use the Flood iQ API, you must register for an API key at https://floodiq.dev.**

### Initalizing the client

```go
// Create a new Flood iq API Client
fiq := floodiq.New("api-key")
```

### **Parcel**

`fiq.Parcel.<method>`

**Property Methods**

- GetPropertyByID(ID `string`) - Retreives a `ParcelProperty` by specific ID
- GetPropertyByLatLng(lat `float64`, lng `float64`) - Retreives a `ParcelProperty` by a coordinate

**City Methods**

- GetCityByID(ID `string`) - Retreives a `ParcelProperty` by specific ID
- GetCityByLatLng(lat `float64`, lng `float64`) - Retreives a `ParcelProperty` by a coordinate

---

## Tiles

Generally, while loading tiles, you can provide an X/Y/Z URL to access the tile service directly (such as api.floodiq.com/tile/hurricane/z/y/x.png?key=abc123&type=c3&year=2032). This portion of the client is included if you need the raw bytes returned from the tile service.

**Tile**

Resolves requested tile

```
fiq.Tile.Hurricane("z", "y", "x", "type", "year")
fiq.Tile.Tidal("z", "y", "x", "type", "year")
```

### Errors

All errors will have a `Code`, `Status` and `Message` attached to it.

### Rate Limits

Rate limit information is provided wit
