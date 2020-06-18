package red_black_tree

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectness(t *testing.T) {
	num := 1000000

	tree := &Tree{}
	// Insert
	list := rand.Perm(num)
	for _, k := range list {
		tree.insert(k, k)
	}
	assert.True(t, tree.validate())

	// Search
	list = rand.Perm(num)
	for _, k := range list {
		value := tree.GetValue(k)
		assert.EqualValues(t, value, k)
	}

	// Remove
	list = rand.Perm(num)
	for _, k := range list {
		tree.Remove(k)
		if k%73171 == 0 {
			assert.True(t, tree.validate())
		}
	}

}
