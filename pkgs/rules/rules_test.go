package rules

import (
	"encoding/json"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestSetRules(t *testing.T) {
	testFile := "./test_rules.json"
	r := &Rules{}

	r.SetRules(testFile)

	j, _ := json.Marshal(r)
	f, err := os.ReadFile(testFile)
	assert.Check(t, err)
	assert.Equal(t, string(j), string(f))
	os.Remove(testFile)
}
