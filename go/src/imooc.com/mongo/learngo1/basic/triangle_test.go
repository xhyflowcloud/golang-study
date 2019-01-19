package basic

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 0},
		{12, 35, 37},
		{3000, 4000, 5000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d);"+
				" got %d; excepted %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
