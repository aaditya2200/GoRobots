package NewRobots

import (
	"math/rand"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Robot struct {
	name *gobot.Robot
	ID *int
}

type token struct {
	ID int
}

func (t *token) CreateNewRobot(name string) *gobot.Robot {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1 * time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot(
		name,
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)
	return robot

}

func GenerateToken() *token{
	userToken := rand.Int()
	return &token{userToken}
}
