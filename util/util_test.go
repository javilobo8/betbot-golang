package util

import (
	"fmt"
	"testing"
)

func TestGetCommandMessages(t *testing.T) {
	var tests = []struct {
		got  string
		want []string
	}{
		{"!bet 123 456", []string{"bet", "123", "456"}},
		{"  !bet 123 456", []string{"bet", "123", "456"}},
		{"!bet 123 456   ", []string{"bet", "123", "456"}},
		{"!bet 123  456   ", []string{"bet", "123", "", "456"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("for %s", tt.got)
		t.Run(testname, func(t *testing.T) {
			ans := GetCommandMessages(tt.got)
			for i := range ans {
				if ans[i] != tt.want[i] {
					t.Errorf("got %s, want %s", ans[i], tt.want[i])
				}
			}
		})
	}
}
