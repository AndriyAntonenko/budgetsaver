package dto

type CreateFinanceGroupPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FinanceGroup struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MembersCount int    `json:"membersCount"`
}
