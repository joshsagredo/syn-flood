package options

import "testing"

func TestGetSynFloodOptions(t *testing.T) {
	t.Log("fetching default options.SynFloodOptions")
	opts := GetSynFloodOptions()
	t.Logf("fetched default options.SynFloodOptions, %v\n", opts)
}
