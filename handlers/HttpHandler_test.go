package handlers

import "testing"

func TestHttpHandler_Handling(t *testing.T) {
	var tests []struct {
		name string
		h    *HttpHandler
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Handling()
		})
	}
}
