package netstat

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetStat(t *testing.T) {
	n := new(NetStat)
	var input interface{}
	e := json.Unmarshal([]byte(`{"netstat":{"measurement": [
						"tcp_established",
						"tcp_syn_sent",
						"tcp_close"]}}`), &input)
	if e == nil {
		_, actual := n.ApplyRule(input)
		expected := []interface{}{map[string]interface{}{
			"fieldpass": []string{"tcp_established", "tcp_syn_sent", "tcp_close"},
		}}
		assert.Equal(t, expected, actual, "Expected to be equal")
	}
}