package sqlDomain

type Transaction struct {
	Id            uint64        `json:"id" gorm:"primary_key"`
	SessionId     Session       `json:"sessionId" gorm:"foreignkey:Id"`
	ProfileId     Profile       `json:"profileId" gorm:"foreignkey:Id"`
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
