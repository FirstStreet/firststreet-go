package firststreet

import (
	// "log"
	// "net/http/httptest"
	// "sync"
	"testing"
	assert "github.com/stretchr/testify/require"
)

// func TestNew(t *testing.T) {
// 	fiq := New("TEST_KEY")

// 	if fiq.Key != "TEST_KEY" {
// 		t.Fatalf("Key mismatched")
// 	}
// }

func TestAPIInit(t *testing.T) {
	a := &API{}
	a.InitAPI("abc123", nil)
	assert.Equal(t, "abc123", a.Parcel.Key)
}

func TestAPINew(t *testing.T) {
	api := newAPI("abc123", nil)
	assert.Equal(t, "abc123", api.Parcel.Key)
}

// func TestIntegration(t *testing.T) {
// 	api := New("VpqgUoepulokFjdxZvE4iMjP8bTtN2PG")
// 	parcel, _ := api.Parcel.GetPropertyByID("100032470544")
// 	t.Log(parcel)
// }
