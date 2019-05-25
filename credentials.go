package stream

import "strings"

type (
	// Credential is used to identfy a token that grants
	// authorization/authentication to a resource
	Credential string
)

// String implements the Stringer interface so as to prevent
// spilling the actual API key/secret in logs
// String formats the key/secret as a debit card (*****-9999)
func (c Credential) String() string {

	cred := string(c)

	n := len(cred)
	return strings.Repeat("*", n-5) + cred[n-5:]
}
