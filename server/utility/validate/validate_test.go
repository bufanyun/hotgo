package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmail(t *testing.T) {
	b := IsEmail("QTT123456@163.com")
	assert.True(t, b, "ok~")
}
