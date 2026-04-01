package main

import (
	"net/http"
	"testing"

	"github.com/steveg22/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// Create a new instance of our application struct. For now, this just
	// contains a structured logger (which uses the slog.DiscardHandler handler
	// and will discard anything written to it with no action).
	app := newTestApplication(t)
	// We then use the httptest.NewTLSServer() function to create a new test
	// server, passing in the value returned by our app.routes() method as the
	// handler for the server. This starts up an HTTPS server which listens on a
	// randomly-chosen port of your local machine for the duration of the test.
	// Notice that we defer a call to ts.Close() so that the server is shut down
	// when the test finishes.
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	// The network address that the test server is listening on is contained in
	// the ts.URL field. We can use this to construct a new HTTP request for the
	// GET /ping route.

	// Use the ts.Client().Do() method to execute the request against the test
	// server. This returns an http.Response struct containing the response.
	res := ts.get(t, "/ping")
	// We can then check the value of the response status code and body using
	// the same pattern as before.
	assert.Equal(t, res.status, http.StatusOK)
	assert.Equal(t, res.body, "OK")
}
