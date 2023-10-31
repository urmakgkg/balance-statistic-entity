package statistic_entity

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	a := Accrual{
		Ls: "001001",
		Header: Header{
			ServiceName: "serviceName",
		},
	}

	strJson, err := json.Marshal(a)
	assert.NoError(t, err)

	println(string(strJson))
}
