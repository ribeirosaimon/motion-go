package sqlDomain

type Transaction struct {
	Id            uint64        `json:"id" gorm:"primary_key"`
	SessionId     string        `json:"-" gorm:"foreignkey:SessionID"`
	ProfileId     uint64        `json:"-" gorm:"foreignkey:ProfileID"`
	Value         float64       `json:"value"`
	OperationType OperationType `json:"operationType"`
	BasicSQL
}

func (t Transaction) GetId() interface{} {
	return t.Id
}

type OperationType string

const (
	BUY      OperationType = "BUY"
	SELL     OperationType = "SELL"
	WITHDRAW OperationType = "WITHDRAW"
	DEPOSIT  OperationType = "DEPOSIT"
)
