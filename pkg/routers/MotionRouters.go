package routers

import (
	"github.com/ribeirosaimon/motion-go/pkg/health"
)

func MotionRouters() {
	health.NewHeathController()
}
