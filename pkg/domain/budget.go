package domain

type CreateBudgetPayload struct {
	GroupId     string `json:"groupId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Budget struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
