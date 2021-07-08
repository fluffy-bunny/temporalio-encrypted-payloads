package cryptconverter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmbeddedCheck(t *testing.T) {
	obj := &SensitivePayload{
		Secret: "p@ssw0rd!",
	}
	result := HasProtectMe(obj)
	require.True(t, result)

	result = HasProtectMe(obj)
	require.True(t, result)

	open := &OpenPayload{Quote: "hello"}
	result = HasProtectMe(open)
	require.False(t, result)

	result = HasProtectMe(open)
	require.False(t, result)
}
