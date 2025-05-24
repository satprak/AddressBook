package models

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type ContactRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateContactsRequest []ContactRequest
type CreateContactsResponse []Contact

type UpdateRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type UpdateContactsRequest []UpdateRequest
type UpdateContactsResponse []Contact

type DeleteContactsRequest []string
type DeleteResponse struct {
	Deleted int `json:"deleted"`
}

type SearchRequest struct {
	Query string `json:"query"`
}
type SearchResponse []Contact
