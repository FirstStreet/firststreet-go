# Go Flood iQ

---
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

```
import (
  "github.com/firststreet/floodiq-go
)
```

## Quick Start

**In order to use the Flood iQ API, you must register for an API key at https://floodiq.dev.**

### Initalizing Client

```

// Create a new Flood iq API Client
fiq := floodiq.Client{
  APIKey: "string-from-floodiq-dev-account",
}

```

### Data Services

**Parcel**

All parcel queries require a `type`.

```

fiq.Data.Parcel.ByID("id", "type")
fiq.Data.Parcel.ByLatLng(lat, lng, "type")
fiq.Data.Parcel.ByAddress("address", "type")
```

**Hurricane**

```
fiq.Data.Hurricane.ByID("id", "type")
fiq.Data.Hurricane.ByLatLng(lat, lng, "type")
fiq.Data.Hurricane.ByAddress("address", "type")
```

**Tidal**

```
fiq.Data.Tidal.ByID("id", "type")
fiq.Data.Tidal.ByLatLng(lat, lng, "type")
fiq.Data.Tidal.ByAddress("address", "type")
```

### Tile Service

Generally, while loading tiles, you can provide an X/Y/Z URL to access the tile service directly (such as api.floodiq.com/tile/hurricane/z/y/x.png?key=abc123&type=c3&year=2032). This portion of the client is included if you need the raw bytes returned from the tile service.

**Tile**

Resolves requested tile

```
fiq.Tile.Hurricane("z", "y", "x", "type", "year")
fiq.Tile.Tidal("z", "y", "x", "type", "year")
```
