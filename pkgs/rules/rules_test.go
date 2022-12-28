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

func TestLoadRules(t *testing.T) {
	testFile := "./test.json"
	testData := "{\"cpu_max\":90,\"cpu_min\":50,\"disk_max\":85,\"disk_min\":60,\"mem_max\":80,\"mem_min\":70}"
	f, _ := os.Create(testFile)
	defer f.Close()
	f.Write([]byte(testData))

	r := &Rules{}
	r.LoadRules(testFile)

	assert.Equal(t, r.Cpu_max, 90)
	assert.Equal(t, r.Disk_min, 60)
	assert.Equal(t, 80, r.Memory_max)
	os.Remove(testFile)
}
