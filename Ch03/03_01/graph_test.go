package graph

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseGraph(t *testing.T) {

	var data = `
ready → running
running → waiting
running → terminated
waiting → running
`

	r := strings.NewReader(data)
	graph, err := ParseGraph(r)
	require.NoError(t, err)
	expected := map[string][]string{
		"ready":   {"running"},
		"running": {"waiting", "terminated"},
		"waiting": {"running"},
	}
	require.Equal(t, expected, graph)
}
