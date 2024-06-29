package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input   http.Header
		wantStr string
		wantErr error
	}{
		{
			input:   http.Header{"Authorization": {"ApiKey 12345"}},
			wantStr: "12345",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		gotStr, gotErr := GetAPIKey(tc.input)
		if gotStr != tc.wantStr {
			t.Fatalf("GetAPIKey(%v) = %v, want %v", tc.input, gotStr, tc.wantStr)
		}
		if gotErr != tc.wantErr {
			t.Fatalf("GetAPIKey(%v) = %v, want %v error", tc.input, gotErr, tc.wantErr)
		}
	}
}
