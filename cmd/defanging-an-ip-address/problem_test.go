package defanging_an_ip_address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceIPAddresses(t *testing.T) {
	p1 := defangIPaddr("1.1.1.1")
	assert.Contains(t, p1, "[.]", "you did it")

	p2 := defangIPaddr("255.100.50.0")
	assert.Contains(t, p2, "[.]", "you did it")

}
