package norepeatingsubstr

import "testing"

func TestNorepeat(tt *testing.T) {

	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},
	}

	for _, ttt := range tests {
		if findMaxSubStr(ttt.s) != ttt.ans {
			tt.Errorf("error")
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "qweqeasdsaxzchiuafhaoifuhaosd"
	ans := 11
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := findMaxSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d", actual, s, ans)
		}
	}
}
