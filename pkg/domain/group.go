package domain

type groupType string

const (
	PersonaGroup groupType = "personal"
	FamilyGroup  groupType = "family"
)

type CreateGroupPayload struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
