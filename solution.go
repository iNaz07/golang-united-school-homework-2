package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	switch sidesNum {
	case 0:
		//area of circle
		return math.Pi * sideLen * sideLen
	case 3:
		//are of triangle
		return sideLen * sideLen * math.Sqrt(3) / 4
	case 4:
		//area of square
		return sideLen * sideLen
	}
	return 0

}
