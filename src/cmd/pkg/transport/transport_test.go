package transport

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

func TestCustomHeaderTransport_RoundTrip_Actor(t *testing.T) {
	tests := []struct {
		name      string
		actor     string
		wantActor string
	}{
		{
			name:      "actor set",
			actor:     "my-actor",
			wantActor: "my-actor",
		},
		{
			name:  "actor empty",
			actor: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var captured http.Header
			transport := CustomHeaderTransport{
				actor: tt.actor,
				agent: "test-agent",
				Rt: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
					captured = r.Header.Clone()
					return &http.Response{StatusCode: http.StatusOK}, nil
				}),
			}

			req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
			require.NoError(t, err)

			_, err = transport.RoundTrip(req)
			require.NoError(t, err)

			assert.Equal(t, tt.wantActor, captured.Get(ActorHeader))
		})
	}
}
