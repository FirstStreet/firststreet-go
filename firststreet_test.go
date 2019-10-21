package firststreet

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNewBackend(t *testing.T) {
	backend := NewBackend("abc123")
	assert.Equal(t, "abc123", backend.Key)
}

// func TestIntegration(t *testing.T) {
// 	api := New("VpqgUoepulokFjdxZvE4iMjP8bTtN2PG")
// 	parcel, _ := api.Parcel.GetPropertyByID("100032470544")
// 	t.Log(parcel)
// }
