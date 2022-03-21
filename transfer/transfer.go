package transfer

type Account struct {
	AccountNumber   string   `json:"account_number,omitempty" bson:"account_number,omitempty"`
	CustumerName  string             `json:"custumer_name" bson:"custumer_name,omitempty"`
	Balance int            `json:"balance" bson:"balance,omitempty"`
}

type Transfer struct {
	ToAccountNumber string `json:"to_account_number,omitempty" bson:"to_account_number,omitempty"`
	Amount  string `json:"amount,omitempty" bson:"amount,omitempty"`
}
