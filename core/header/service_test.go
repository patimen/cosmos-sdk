package header

import (
	"crypto/sha256"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInfo_Bytes(t *testing.T) {
	sum := sha256.Sum256([]byte("test-chain"))
	info := Info{
		Height:  12345,
		Hash:    sum[:],
		Time:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		AppHash: sum[:],
		ChainID: "test-chain",
	}

	expectedBytes := []byte{
		0x39, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Height (little-endian)
		0x26, 0xb0, 0xb8, 0x3e, 0x72, 0x81, 0xbe, 0x3b, 0x11, 0x76, 0x58, 0xb6, 0xf2, 0x63, 0x6d, 0x3, 0x68, 0xca, 0xd3, 0xd7, 0x4f, 0x22, 0x24, 0x34, 0x28, 0xf5, 0x40, 0x1a, 0x4b, 0x70, 0x89, 0x7e, // Hash
		0x80, 0x0, 0x92, 0x65, 0x0, 0x0, 0x0, 0x0, // Time (little-endian)
		0x26, 0xb0, 0xb8, 0x3e, 0x72, 0x81, 0xbe, 0x3b, 0x11, 0x76, 0x58, 0xb6, 0xf2, 0x63, 0x6d, 0x3, 0x68, 0xca, 0xd3, 0xd7, 0x4f, 0x22, 0x24, 0x34, 0x28, 0xf5, 0x40, 0x1a, 0x4b, 0x70, 0x89, 0x7e, // Apphash
		0x0A,                                                       // ChainID length
		0x74, 0x65, 0x73, 0x74, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, // ChainID
	}

	bytes, err := info.Bytes()
	require.NoError(t, err)
	require.Equal(t, expectedBytes, bytes)
}

func TestInfo_FromBytes(t *testing.T) {
	info := Info{}

	// Test case 1: Valid byte slice
	bytes := []byte{
		0x39, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Height (little-endian)
		0x26, 0xb0, 0xb8, 0x3e, 0x72, 0x81, 0xbe, 0x3b, 0x11, 0x76, 0x58, 0xb6, 0xf2, 0x63, 0x6d, 0x3, 0x68, 0xca, 0xd3, 0xd7, 0x4f, 0x22, 0x24, 0x34, 0x28, 0xf5, 0x40, 0x1a, 0x4b, 0x70, 0x89, 0x7e, // Hash
		0x80, 0x0, 0x92, 0x65, 0x0, 0x0, 0x0, 0x0, // Time (little-endian)
		0x26, 0xb0, 0xb8, 0x3e, 0x72, 0x81, 0xbe, 0x3b, 0x11, 0x76, 0x58, 0xb6, 0xf2, 0x63, 0x6d, 0x3, 0x68, 0xca, 0xd3, 0xd7, 0x4f, 0x22, 0x24, 0x34, 0x28, 0xf5, 0x40, 0x1a, 0x4b, 0x70, 0x89, 0x7e, // Apphash
		0x0A,                                                       // ChainID length
		0x74, 0x65, 0x73, 0x74, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, // ChainID
	}

	err := info.FromBytes(bytes)
	require.NoError(t, err)
	require.Equal(t, int64(12345), info.Height)
	require.Equal(t, []byte{0x26, 0xb0, 0xb8, 0x3e, 0x72, 0x81, 0xbe, 0x3b, 0x11, 0x76, 0x58, 0xb6, 0xf2, 0x63, 0x6d, 0x3, 0x68, 0xca, 0xd3, 0xd7, 0x4f, 0x22, 0x24, 0x34, 0x28, 0xf5, 0x40, 0x1a, 0x4b, 0x70, 0x89, 0x7e}, info.Hash)
	require.Equal(t, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), info.Time)
	require.Equal(t, []byte{0x26, 0xb0, 0xb8, 0x3e, 0x72, 0x81, 0xbe, 0x3b, 0x11, 0x76, 0x58, 0xb6, 0xf2, 0x63, 0x6d, 0x3, 0x68, 0xca, 0xd3, 0xd7, 0x4f, 0x22, 0x24, 0x34, 0x28, 0xf5, 0x40, 0x1a, 0x4b, 0x70, 0x89, 0x7e}, info.AppHash)
	require.Equal(t, "test-chain", info.ChainID)

}