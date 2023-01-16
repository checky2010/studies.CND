package rabbitmq

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getEnv(t *testing.T) {
	assert.Equal(t, "yes", getEnv("not-set", "yes"))

	err := os.Setenv("definitely-set", "other-yes")
	if err != nil {
		t.Error("Shouldn't have failed")
	}
	assert.Equal(t, "other-yes", getEnv("definitely-set", "no"))
}
