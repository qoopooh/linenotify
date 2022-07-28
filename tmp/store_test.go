package tmp

import (
	"testing"
)

func TestGetAndSet(t *testing.T) {
	store := Store{FileName: "store.json"}

	node := map[string]string{
		"first": "1",
	}

	store.Set(node)

	var out map[string]string
	store.Get(&out)
	if out["first"] != node["first"] {
		t.Fatalf("Set (%q) is not the same as Get (%q)", node, out)
	}
}
