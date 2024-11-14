package freq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordFreq(t *testing.T) {
	text := `
	Außer Hunden sind Bücher die besten Freunde des Menschen,
	in Hunden dagegen ist es zum Lesen zu dunkel.
	`

	expected := map[string]int{
		"Außer":1, 
		"Bücher":1, 
		"Freunde":1, 
		"Hunden":2, 
		"Lesen":1, 
		"Menschen":1, 
		"besten":1, 
		"dagegen":1, 
		"des":1, 
		"die":1, 
		"dunkel":1, 
		"es":1, 
		"in":1, 
		"ist":1, 
		"sind":1, 
		"zu":1, 
		"zum":1,
	}
        	            
	freqs := WordFreq(text)
	require.Equal(t, expected, freqs)
}
