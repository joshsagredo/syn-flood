package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSynFloodOptions(t *testing.T) {
	t.Log("fetching default options.SynFloodOptions")
	opts := GetSynFloodOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.SynFloodOptions, %v\n", opts)
}
