package root

import "testing"

func TestSendUserCmd(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestSendUserCmd_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendUserCmd([]byte{})
		})
	}
}