package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	type args struct {
		scenario      string
		filePath      string
		errorExpected bool
	}
	tests := []args{
		{
			scenario:      "success",
			filePath:      "testdata/testdata_success.json",
			errorExpected: false,
		},
		{
			scenario:      "bad file path",
			filePath:      "bad/path",
			errorExpected: true,
		},
		{
			scenario:      "bad config file",
			filePath:      "testdata/testdata_error.txt",
			errorExpected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			_, err := InitDB(test.filePath)
			if !test.errorExpected {
				assert.NoError(t, err)
				return
			}
			assert.Error(t, err)
		})
	}
}
