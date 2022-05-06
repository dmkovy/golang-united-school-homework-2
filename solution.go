package solution

import "math"

type mySidesNumber int

const SidesTriangle mySidesNumber = 3
const SidesSquare mySidesNumber = 4
const SidesCircle mySidesNumber = 0

func CalcSquare(sideLen float64, sidesNum mySidesNumber) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3) * sideLen * sideLen / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	}
	return 0
}
