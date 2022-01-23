package elon

import "fmt"

func (c *Car) Drive() {
    if c.batteryDrain > c.battery {
        return
    }
    c.distance += c.speed
    c.battery -= c.batteryDrain
}

func (c Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
    numberOfDrives := c.battery / c.batteryDrain
    if trackDistance <= c.speed * numberOfDrives {
        return true
    }
	return false
}
