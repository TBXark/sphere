package idgenerator

import (
	"math"
	"testing"
)

func TestNextId(t *testing.T) {
	t.Log(NextId())
	t.Log(math.MaxInt32)
}
