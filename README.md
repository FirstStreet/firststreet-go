# Go First Street Foundation

**WARNING**: This library is beta and subject to change. It's not ready to be used in production quite yet.

---

This is the official implementation of the First Street Foundation API in Go. Please use this client if you are using Go to interact with the First Street Foundation API.

For more in-depth guides, usage and API access, please see the documentation at docs.firststreet.dev.

## Installation

First install firststreet-go:

```
go get -u github.com/firststreet/firststreet-go
```

Next, import it into your project:

```go
import (
  "github.com/firststreet/firststreet-go
)
```

## Documentation

**In order to use the First Street Foundation API, you must register for an API key at https://firststreet.dev.**

### Initalizing the client

```go
// Create a new Flood iq API Client
fsf := firststreet.New("api-key")
```

### **Risk Summary**

The risk summary API provides metadata and risks summary for a given `location`.

`fsf.Summary.<method>`

**Property**

- GetPropertyByID(ID `string`) - Retreives a `ParcelProperty` by specific ID
- GetPropertyByLatLng(lat `float64`, lng `float64`) - Retreives a `ParcelProperty` by a coordinate

**City**

- GetCityByID(ID `string`) - Retreives a `ParcelProperty` by specific ID
- GetCityByLatLng(lat `float64`, lng `float64`) - Retreives a `ParcelProperty` by a coordinate


### Errors

All errors will have a `Code`, `Status` and `Message` attached to it.

### Rate Limits

Rate limit information is provided wit
