package errors

import "net/http"

// NotAuthorized is a compliant error for use when an authenticated client is
// not allowed to perform a specific action on specific resources. In other
// words, when they don't have permission.
type NotAuthorized struct {
	Message  string
	Metadata M
}

func (na NotAuthorized) Error() string      { return string(KindNotAuthorized) }
func (na NotAuthorized) GetCode() int       { return http.StatusForbidden }
func (na NotAuthorized) GetKind() K         { return KindNotAuthorized }
func (na NotAuthorized) GetMessage() string { return na.Message }
func (na NotAuthorized) GetMetadata() M     { return na.Metadata }
