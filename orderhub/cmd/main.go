package main

import (
	"fmt"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/confighub/util"
	"github.com/ribeirosaimon/motion-go/orderHub/akafka"
	"github.com/ribeirosaimon/motion-go/orderHub/order"
)

func main() {
	gear := order.NewMotionGear()
	channel := akafka.NewKafkaMotion().CreateConn(getProperties())

	for {
		msg := <-channel
		gear.Process(msg.Value)
	}

}

func getProperties() *properties.Properties {
	propertiesFile := "config.properties"
	dir, _ := util.FindRootDir()

	return properties.MustLoadFile(fmt.Sprintf("%s/%s", dir, propertiesFile), properties.UTF8)
}
