package containsduplicate2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsNearByDuplicate(t *testing.T) {
	q1 := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
	assert.Equal(t, true, q1)

	q2 := containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	assert.Equal(t, true, q2)

	q3 := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	assert.Equal(t, false, q3)

}
