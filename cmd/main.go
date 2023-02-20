package main

import (
	"github.com/ribeirosaimon/motion-go/pkg/config"
	"github.com/ribeirosaimon/motion-go/pkg/routers"
)

func main() {
	config.StartupEnginer(routers.MotionRouters)
}
