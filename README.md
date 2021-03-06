# Go First Street Foundation [![CircleCI](https://circleci.com/gh/FirstStreet/firststreet-go.svg?style=svg&circle-token=1c06b4bacc08aeb1028dddc5c86d09b5c417d1cd)](https://circleci.com/gh/FirstStreet/firststreet-go) [![codecov](https://codecov.io/gh/FirstStreet/firststreet-go/branch/master/graph/badge.svg?token=KGfgBydmly)](https://codecov.io/gh/FirstStreet/firststreet-go) [![](https://godoc.org/github.com/firststreet/firststreet-go?status.svg)](https://godoc.org/github.com/FirstStreet/firststreet-go)

**Notice**: This library is beta and subject to change.

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

**In order to use the First Street Foundation API, you must register for an API key at https://firststreet.dev.** For more in-depth documentation and API reference, see https://docs.firststreet.dev

### Initalizing the client

```go
// Create a new backend
backend := firststreet.NewBackend("api-key")

// Create the API client
fsf := client.New(backend)
```

## Lookup

Provide a `Lookup` to the specific product. A [lookup](https://docs.firststreet.dev/docs/lookups) determines how the location will be queried.

### FSID

```go
lookup := &firststreet.Lookup{
  FSID: "123456",
}
```

### Coordinate

```go
lookup := &firststreet.Lookup{
  Lat: 51.4779,
  Lng: 0.0015,
}
```

### Address

```go
lookup := &firststreet.Lookup{
  Address: "247 Water Street, Brooklyn NY, 11201",
}
```

## **Risk Summary**

The Risk Summary API provides metadata and risks summary for a given `location`. Depending on the location type, a `PropertySummary` or `CitySummary` will be returned.

### Property

- `fsf.PropertySummary.Lookup(lookup)`

### City

- `fsf.CitySummary.Lookup(lookup)`

## Location Types

When doing a data request that is outside of Risk Summary, a `locationType` is required, depending on the type of location you are querying.

- `firststreet.PropertyLocationType` provides the `property` location type
- `firststreet.CityLocationType` provides the `city` location type

## **Hurricane Risk**

The Hurricane Risk API provides hurricane risk for a given `location`.

`fsf.Hurricane.Lookup(locationType, lookup)`

## **Tidal Risk**

The Tidal Risk API provides tidal risk for a given `location`.

`fsf.Tidal.Lookup(locationType, lookup)`

## **Market Value Impact**

The Market Value Impact API provides tidal risk for a given `location`.

`fsf.MVI.Lookup(locationType, lookup)`

## Errors

All errors will have a `Code`, `Status` and `Message` attached to it.

## Rate Limits

Rate limit information is provided wit
