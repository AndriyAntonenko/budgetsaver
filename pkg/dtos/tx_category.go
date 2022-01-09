package dto

type TxCategoryDto struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Creator      *string `json:"creator"`
	FinanceGroup *string `json:"financeGroup"`
}

type CreateTxCategoryDto struct {
	Name         string  `json:"name"`
	FinanceGroup *string `json:"financeGroup"`
}
