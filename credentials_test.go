package stream

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCredential_String(t *testing.T) {

	tt := []struct {
		original Credential
		expected string
	}{
		{Credential("hdjjdkjsksk"), "******jsksk"},
		{Credential("adelowolanre"), "*******lanre"},
		{Credential("ppppppppjsksk"), "********jsksk"},
		{Credential("dddddddsksk"), "******dsksk"},
	}

	for k, v := range tt {
		t.Logf("Running iteration number  %d", k)
		require.Equal(t, v.expected, v.original.String())
	}
}
