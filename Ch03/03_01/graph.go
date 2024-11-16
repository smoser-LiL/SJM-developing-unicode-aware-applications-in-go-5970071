package graph

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ParseGraph parses a graph containing lines like "A → B"
// and returns map of src → []dest.
func ParseGraph(r io.Reader) (map[string][]string, error) {
	graph := make(map[string][]string)
	s := bufio.NewScanner(r)
	lnum := 0
	for s.Scan() {
		lnum++
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			continue
		}

		i := strings.Index(line, "→")
		if i == -1 {
			return nil, fmt.Errorf("%d: %q: no → in line", lnum, line)
		}

		src := strings.TrimSpace(line[:i])
		dest := strings.TrimSpace(line[i+1:])
		graph[src] = append(graph[src], dest)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}
