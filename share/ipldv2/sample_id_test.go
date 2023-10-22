package ipldv2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/eds/edstest"
)

func TestSampleID(t *testing.T) {
	square := edstest.RandEDS(t, 2)
	root, err := share.NewRoot(square)
	require.NoError(t, err)

	sid := NewSampleID(rsmt2d.Row, 2, root, 1)

	id, err := sid.Cid()
	require.NoError(t, err)

	assert.EqualValues(t, sampleCodec, id.Prefix().Codec)
	assert.EqualValues(t, sampleMultihashCode, id.Prefix().MhType)
	assert.EqualValues(t, SampleIDSize, id.Prefix().MhLength)

	data, err := sid.MarshalBinary()
	require.NoError(t, err)

	sidOut := SampleID{}
	err = sidOut.UnmarshalBinary(data)
	require.NoError(t, err)
	assert.EqualValues(t, sid, sidOut)

	err = sidOut.Validate()
	require.NoError(t, err)
}
