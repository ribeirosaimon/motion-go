package nosqlDomain

import (
	sqlDomain2 "github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/shopspring/decimal"
)

type ShoppingCart struct {
	Id        uint64               `json:"id"`
	Owner     sqlDomain2.Profile   `json:"owner" gorm:"foreignkey:Id"`
	ProfileId uint64               `json:"profileId"`
	Price     decimal.Decimal      `json:"price"`
	Products  []sqlDomain2.Product `json:"products" gorm:"many2many:cart_product;"`
	BasicNoSQL
}

func (s ShoppingCart) GetId() interface{} {
	return s.Id
}
