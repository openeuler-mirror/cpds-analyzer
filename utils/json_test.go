package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestSaveAsJsonFile(t *testing.T) {
	testFile := "./test.json"

	type jsonSubData struct {
		D string `json:"D"`
	}

	type jsonData struct {
		A string       `json:"A"`
		B int          `json:"B"`
		C *jsonSubData `json:"C"`
	}

	j := &jsonData{
		A: "cpds",
		B: 2,
		C: &jsonSubData{
			D: "json",
		},
	}

	SaveAsJsonFile(testFile, j)

	f, err := os.Open(testFile)
	assert.Check(t, err)

	d, err := ioutil.ReadAll(f)
	assert.Check(t, err)

	m, _ := json.Marshal(j)
	assert.Equal(t, string(d), string(m))

	os.Remove(testFile)
}
