package domain

type FinanceGroupRole string

const (
	Owner  FinanceGroupRole = "owner"
	Admin  FinanceGroupRole = "admin"
	Member FinanceGroupRole = "member"
)
