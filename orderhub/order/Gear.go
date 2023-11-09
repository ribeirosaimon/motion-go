package order

import (
	"fmt"

	"github.com/ribeirosaimon/motion-go/confighub/domain/nosqlDomain"
)

type MotionGear struct {
	BrokerStocks []nosqlDomain.SummaryStock
}

var motionGear *MotionGear

func NewMotionGear() *MotionGear {
	if motionGear == nil {
		motionGear = &MotionGear{}
	}
	return motionGear
}

func (g *MotionGear) Process(value []byte) {
	fmt.Println(string(value))
}
