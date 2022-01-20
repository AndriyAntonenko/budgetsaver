package dto

type CreateBudgetPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupId     string `json:"groupId"`
	Currency    string `json:"currency"`
}

type Budget struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupId     string `json:"groupId"`
	Creator     string `json:"creator"`
	Currency    string `json:"currency"`
}
