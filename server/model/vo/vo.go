package vo

type UpdateResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"userName"`
	MembershipType string `json:"membershipType"`
}

type GetResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type CreateResponse struct {
	ID             string `json:"id"`
	MembershipType string `json:"membershipType"`
}
