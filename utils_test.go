package freeswitchesl

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_buildVars(t *testing.T) {
	vars := buildVars(map[string]string{
		"origination_caller_name":   "test",
		"origination_caller_number": "1234",
		"origination_callee_name":   "John Doe",
		"origination_callee_number": "7100",
	})

	// Contains since order is not guaranteed when iterating over maps
	assert.Contains(t, vars, "origination_caller_name=test")
	assert.Contains(t, vars, "origination_caller_number=1234")
	assert.Contains(t, vars, "origination_callee_name='John Doe'")
	assert.Contains(t, vars, "origination_callee_number=7100")

	// Ensure the formatting elements are contained in the string
	assert.Equal(t, 3, strings.Count(vars, ","))
	assert.True(t, strings.HasPrefix(vars, "{"))
	assert.True(t, strings.HasSuffix(vars, "}"))
}
