package addressbook

// These are primarily models for the HTTP/JSON endpoints.  For protobuf, use
// pb.

type ContactRequest struct {
	ID         int64  `json:"id"`
	TenantID   int64  `json:"tenant_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Active     bool   `json:"active"`
	Address    string `json:"address"`
	SomeSecret string `json:"some_secret" encryption:"true"`
}

type ContactResponse struct {
	ID         int64  `json:"id"`
	TenantID   int64  `json:"tenant_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Active     bool   `json:"active"`
	Address    string `json:"address"`
	SomeSecret string `json:"some_secret" encryption:"true"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ReadContactRequest struct {
	ID int64 `json:"id"`
}

type DeleteContactRequest struct {
	ID int64 `json:"id"`
}

// Data is an array of ContactResponse structs
type StandardPayloadResponse struct {
	Data          []any                  `json:"data"`
	NextPageToken int64                  `json:"next_page_token"`
	Errors        []StandardPayloadError `json:"errors"`
	CorrelationID string                 `json:"correlation_id"`
}

type StandardPayloadError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}
