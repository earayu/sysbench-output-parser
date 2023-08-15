package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

const RunningStringFlag = "sysbench 1.0.20 (using system LuaJIT 2.1.0-beta3)"

func partitionAsReaders(content string) ([]io.Reader, error) {
	if content == "" {
		return nil, fmt.Errorf("content is empty")
	}
	readers := make([]io.Reader, 0)
	regex := regexp.MustCompile(`cd ./sysbench/src/lua.*\n`)
	chunk := regex.Split(content, -1)

	for _, c := range chunk {
		c = strings.Trim(c, "\n")
		if !strings.HasPrefix(c, RunningStringFlag) {
			continue
		}
		readers = append(readers, strings.NewReader(c))
	}
	return readers, nil
}
