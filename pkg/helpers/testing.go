package helpers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertErrorIs will assert two errors are equal while returning a better error message
func AssertErrorIs(t *testing.T, err error, target error) {
	if target == nil {
		assert.NoError(t, err)
	} else if err == nil {
		assert.Error(t, err)
	} else {
		assert.True(t, errors.Is(err, target), fmt.Sprintf("Received incorrect error:\nReceived: %+v", err))
	}
}
