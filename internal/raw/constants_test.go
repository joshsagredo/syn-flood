package raw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpRegex(t *testing.T) {
	expected := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	assert.Equal(t, expected, IpRegex)
}

func TestDnsRegex(t *testing.T) {
	expected := "^(([a-zA-Z]|[a-zA-Z][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$"
	assert.Equal(t, expected, DnsRegex)
}
