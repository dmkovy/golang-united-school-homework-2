package solution

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type myCalcSquareInt int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
	SidesCircle   myCalcSquareInt = 0
	SidesTriangle myCalcSquareInt = 3
	SidesSquare   myCalcSquareInt = 4
)

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum myCalcSquareInt) float64 {
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

//func main() {
//	fmt.Println(CalcSquare(10.0, SidesTriangle))
//	fmt.Println(CalcSquare(10.0, SidesSquare))
//	fmt.Println(CalcSquare(10.0, SidesCircle))
//	fmt.Println(CalcSquare(10.0, 2))
//}
