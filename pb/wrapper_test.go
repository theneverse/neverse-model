package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theneverse/neverse-kit/types"
)

func TestInterchain_Marshal(t *testing.T) {
	ic := &Interchain{
		ID:                   "1",
		InterchainCounter:    make(map[string]uint64),
		ReceiptCounter:       make(map[string]uint64),
		SourceReceiptCounter: nil,
	}

	ic.InterchainCounter["1"] = 2

	data, err := ic.Marshal()
	assert.Nil(t, err)

	ic1 := &Interchain{}
	err = ic1.Unmarshal(data)
	assert.Nil(t, err)

	assert.Equal(t, ic.ID, ic1.ID)
	assert.Equal(t, ic.InterchainCounter, ic1.InterchainCounter)
	assert.Equal(t, ic.ReceiptCounter, ic1.ReceiptCounter)
	assert.Equal(t, 0, len(ic1.SourceReceiptCounter))
}

func TestInterchainMeta_Marshal(t *testing.T) {
	icm := &InterchainMeta{
		Counter: make(map[string]*VerifiedIndexSlice),
	}

	data, err := icm.Marshal()
	assert.Nil(t, err)

	icm1 := &InterchainMeta{}
	err = icm1.Unmarshal(data)
	assert.Nil(t, err)

	assert.Equal(t, icm.Counter, icm1.Counter)
	assert.Equal(t, icm.L2Roots, icm1.L2Roots)

	icm.Counter["1"] = &VerifiedIndexSlice{Slice: []*VerifiedIndex{
		{Index: 1, Valid: true},
		{Index: 2, Valid: false},
	}}
	icm.L2Roots = append(icm.L2Roots, types.Hash{})

	data, err = icm.Marshal()
	assert.Nil(t, err)

	icm2 := &InterchainMeta{}
	err = icm2.Unmarshal(data)
	assert.Nil(t, err)

	assert.Equal(t, icm.Counter, icm2.Counter)
	assert.Equal(t, icm.L2Roots, icm2.L2Roots)
}
