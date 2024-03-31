package rest

type Account struct {
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

type AccountsEnvelope struct {
	Accounts []Account `json:"accounts"`
}

type AccountEnvelope struct {
	Account Account `json:"account"`
}
