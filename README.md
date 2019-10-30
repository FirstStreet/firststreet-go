# Go First Street Foundation

**WARNING**: This library is beta and subject to change. It's not ready to be used in production quite yet.

---

This is the official implementation of the First Street Foundation API in Go. Please use this client if you are using Go to interact with the First Street Foundation API.

For more in-depth guides, usage and API access, please see the documentation at docs.firststreet.dev.

## Installation

First install `firststreet-go`:

```
go get -u github.com/firststreet/firststreet-go
```

Next, import `firststreet` and `client` to your project. `firststreet` is the backend that defines which version and server the client will use.

```go
import (
  "github.com/firststreet/firststreet-go"
  "github.com/firststreet/firststreet-go/client"
)
```

## Documentation

**In order to use the First Street Foundation API, you must register for an API key at https://firststreet.dev.**

### Initalizing the client

```go
// Create a new backend
backend := firststreet.NewBackend("api-key")

// Create the API client
fs := client.New(backend)
```

### **Risk Summary**

The risk summary API provides metadata and risks summary for a given `location`.

`fsf.DataSummary.<method>`

**Property**

- GetPropertyByFSID(FSID `string`) - Retreives a `ParcelProperty` by specific ID
- GetPropertyByLatLng(lat `float64`, lng `float64`) - Retreives a `ParcelProperty` by a coordinate
- GetPropertyByAddress(address `string`) - Retrieves a `ParcelProperty` by address lookup

**City**

- GetCityByID(ID `string`) - Retreives a `ParcelProperty` by specific ID
- GetCityByLatLng(lat `float64`, lng `float64`) - Retreives a `ParcelProperty` by a coordinate
- GetCityByAddress(address `string`) - Retrieves a `ParcelProperty` by address lookup

### **Hurricane Risk**

The hurricane risk API provides hurricane risk for a given `location`.

`fsf.Hurricane.<method>`

**Property**

- GetPropertyByFSID(FSID `string`) - Retreives Hurricane Risk for a property by specific ID
- GetPropertyByLatLng(lat `float64`, lng `float64`) - Retreives Hurricane Risk for a property by a coordinate
- GetPropertyByAddress(lat `float64`, lng `float64`) - Retreives Hurricane Risk for a property by address lookup

**City**

- GetCityByID(ID `string`) - Retreives Hurricane Risk for a city by specific ID
- GetCityByLatLng(lat `float64`, lng `float64`) - Retreives a Hurricane Risk for a city by a coordinate
- GetCityByAddress(lat `float64`, lng `float64`) - Retreives Hurricane Risk for a city by address lookup

### **Tidal Risk**

The Tidal risk API provides tidal risk for a given `location`.

`fsf.Tidal.<method>`

**Property**

- GetPropertyByFSID(FSID `string`) - Retreives Tidal Risk for a property by specific ID
- GetPropertyByLatLng(lat `float64`, lng `float64`) - Retreives Tidal Risk for a property by a coordinate
- GetPropertyByAddress(lat `float64`, lng `float64`) - Retreives Tidal Risk for a property by address lookup

**City**

- GetCityByID(ID `string`) - Retreives Tidal Risk for a city by specific ID
- GetCityByLatLng(lat `float64`, lng `float64`) - Retreives a Tidal Risk for a city by a coordinate
- GetCityByAddress(lat `float64`, lng `float64`) - Retreives Tidal Risk for a city by address lookup

### **Market Value Impact**

The Market Value Impact API provides tidal risk for a given `location`.

`fsf.MVI.<method>`

**Property**

- GetPropertyByFSID(FSID `string`) - Retreives Market Value Impact for a property by specific ID
- GetPropertyByLatLng(lat `float64`, lng `float64`) - Retreives Market Value Impact for a property by a coordinate
- GetPropertyByAddress(lat `float64`, lng `float64`) - Retreives Market Value Impact for a property by address lookup

**City**

- GetCityByID(ID `string`) - Retreives Market Value Impact for a city by specific ID
- GetCityByLatLng(lat `float64`, lng `float64`) - Retreives a Market Value Impact for a city by a coordinate
- GetCityByAddress(lat `float64`, lng `float64`) - Retreives Market Value Impact for a city by address lookup

### Errors

All errors will have a `Code`, `Status` and `Message` attached to it.

### Rate Limits

Rate limit information is provided wit
