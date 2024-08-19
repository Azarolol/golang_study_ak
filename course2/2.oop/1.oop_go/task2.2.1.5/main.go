package main

import "fmt"

type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

type BaseMover struct {
	speed int
}

type SlowMover struct {
	BaseMover
}

type FastMover struct {
	BaseMover
}

func (m SlowMover) Move() string {
	return fmt.Sprint("Slow mover... ", m.BaseMover.Move())
}

func (m FastMover) Move() string {
	return fmt.Sprint("Fast mover! ", m.BaseMover.Move())
}

func (m BaseMover) Move() string {
	return fmt.Sprint("Moving at speed: ", m.speed)
}

func (m BaseMover) Speed() int {
	return m.speed
}

func (m BaseMover) MaxSpeed() int {
	return 120
}

func (m BaseMover) MinSpeed() int {
	return 10
}

func main() {
	var movers []Mover
	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}
	movers = append(movers, fm, sm)

	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
