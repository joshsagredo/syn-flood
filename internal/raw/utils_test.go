package raw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveHost(t *testing.T) {
	cases := []struct {
		caseName string
		host     string
	}{
		{"case1", "example.com"},
		{"case2", "93.184.216.34"},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			res := resolveHost(tc.host)
			assert.NotNil(t, res)
		})
	}
}

func TestIsDNS(t *testing.T) {
	cases := []struct {
		caseName string
		host     string
		expected bool
	}{
		{"case1", "example.com", true},
		{"case2", "93.184.216.34", false},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			assert.Equal(t, isDNS(tc.host), tc.expected)
		})
	}
}

func TestIsIP(t *testing.T) {
	cases := []struct {
		caseName string
		host     string
		expected bool
	}{
		{"case1", "example.com", false},
		{"case2", "93.184.216.34", true},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			assert.Equal(t, isIP(tc.host), tc.expected)
		})
	}
}
