package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partitionSysbenchLogAsReaders(t *testing.T) {
	fileName := "../../sample/kb_sysbench.log"
	content, err := ReadAsString(fileName)
	if err != nil {
		assert.NoError(t, err)
	}
	readers, err := partitionAsReaders(content)
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, 6, len(readers))
}
