package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorgerLogs(t *testing.T) {
	arr := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	result := reorderLogFiles(arr)
	assert.Equal(t, []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}, result)
}
