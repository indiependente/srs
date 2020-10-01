package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Hello(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		timeout       time.Duration
		queryDuration string
		wantStatus    int
	}{
		{
			name:          "Happy path",
			timeout:       3 * time.Second,
			queryDuration: "1",
			wantStatus:    http.StatusOK,
		},

		{
			name:          "Sad path",
			timeout:       1 * time.Second,
			queryDuration: "2",
			wantStatus:    http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(hello(tt.timeout))
			srv := httptest.NewServer(handler)
			client := srv.Client()
			resp, err := client.Get(srv.URL + "/hello?t=" + tt.queryDuration)
			require.NoError(t, err)
			require.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}

func Test_Hello2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		timeout       time.Duration
		queryDuration string
		wantStatus    int
	}{
		{
			name:          "Happy path",
			timeout:       3 * time.Second,
			queryDuration: "1",
			wantStatus:    http.StatusOK,
		},

		{
			name:          "Sad path",
			timeout:       1 * time.Second,
			queryDuration: "2",
			wantStatus:    http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(hello2(tt.timeout))
			srv := httptest.NewServer(handler)
			client := srv.Client()
			resp, err := client.Get(srv.URL + "/hello2?t=" + tt.queryDuration)
			require.NoError(t, err)
			require.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}
