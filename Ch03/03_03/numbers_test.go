package numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindNums(t *testing.T) {
	text := "English: 42, Arabic: ٤٢"
	nums, err := FindNums(text)
	require.NoError(t, err)
	require.Equal(t, []int{42, 42}, nums)
}
