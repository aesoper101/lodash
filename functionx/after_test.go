package functionx

import (
	"testing"
)

func TestAfter(t *testing.T) {
	done := After(3, func() {
		t.Log("3")
	})

	done()
	done()
	//done()
}
