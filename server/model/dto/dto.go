package dto

type CreateRequest struct {
	UserName       string `json:"userName"`
	MembershipType string `json:"membershipType"`
}

type UpdateRequest struct {
	ID             string `json:"id"`
	UserName       string `json:"userName"`
	MembershipType string `json:"membershipType"`
}
