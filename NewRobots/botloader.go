package NewRobots

import (
	"math/rand"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))


type token struct {
	ID int
}

func (t *token) CreateNewRobotWithBlinkFunc(name string, pin string, dur time.Duration) *gobot.Robot {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	workFunc := func() {
		gobot.Every(dur, func() {
			led.Toggle()
		})
	} 

	robot := gobot.NewRobot(
		name,
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		workFunc,
	)

	return robot
} 


func GenerateToken() *token{
	userToken := rand.Int()
	return &token{userToken}
}
