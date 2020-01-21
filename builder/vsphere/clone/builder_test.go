package clone

import (
	"github.com/hashicorp/packer/packer"
	"testing"
)

func TestCloneBuilder_ImplementsBuilder(t *testing.T) {
	var raw interface{}
	raw = &Builder{}
	if _, ok := raw.(packer.Builder); !ok {
		t.Fatalf("Builder should be a builder")
	}
}
