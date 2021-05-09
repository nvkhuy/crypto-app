package handlers

import "testing"

func TestGinHandler_Handling(t *testing.T) {
	var tests []struct {
		name string
		g    *GinHandler
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Handling()
		})
	}
}
