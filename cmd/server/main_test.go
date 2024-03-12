package main

import (
	"encoding/json"
	"net/http"
	"server/api"
	"testing"

	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	var err error
	s := NewServer(api.NewAPIHandler(), "8080")
	r := s.Handler

	t.Run("Hello", func(t *testing.T) {
		rr := testutil.NewRequest().Get("/hello").GoWithHTTPHandler(t, r).Recorder
		assert.Equal(t, http.StatusOK, rr.Code)

		var body string
		err = json.NewDecoder(rr.Body).Decode(&body)
		assert.NoError(t, err, "error unmarshaling response")
		assert.Equal(t, body, "Hello, world!")
	})
}
