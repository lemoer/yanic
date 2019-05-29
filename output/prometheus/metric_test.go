package prometheus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetric(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input  Metric
		err    string
		output string
	}{
		{
			input: Metric{Name: "test1"},
			err:   "no value of metric found",
		},
		{
			input:  Metric{Name: "test2", Value: 3},
			output: "test2 3",
		},
		{
			input: Metric{Name: "test3", Value: 3.2,
				Labels: map[string]interface{}{
					"site_code": "lola",
				},
			},
			output: `test3{site_code="lola"} 3.2`,
		},
		{
			input: Metric{Name: "test4", Value: "0",
				Labels: map[string]interface{}{
					"frequency": float32(3.2),
				},
			},
			output: `test4{frequency="3.2000"} 0`,
		},
		{
			input: Metric{Name: "test5", Value: 3,
				Labels: map[string]interface{}{
					"node_id": "lola",
					"blub":    3.3423533,
				},
			},
			output: `test5{node_id="lola",blub="3.3424"} 3`,
		},
	}

	for _, test := range tests {
		output, err := test.input.String()

		if test.err == "" {
			assert.NoError(err)
			assert.Equal(test.output, output)
		} else {
			assert.EqualError(err, test.err)
		}
	}
}