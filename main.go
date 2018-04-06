package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func main() {
	for i := 0; i < 1; i++ {
		ts, err := host.SensorsTemperatures()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, t := range ts {
			fmt.Printf("%s: %4.2f\n", t.SensorKey, t.Temperature)
		}
	}
}
