package main

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}
 
type Circle struct {
    Point
    R float64
}
 
func Intersect(a *Circle, b *Circle) (p []Point) {
	dx, dy := b.X - a.X, b.Y - a.Y
	lr := a.R + b.R //radius and
	dr := math.Abs(a.R - b.R) //radius difference
	ab := math.Sqrt(dx * dx + dy * dy) //center distance

    if ab <= lr && ab > dr {
		theta1 := math.Atan(dy / dx)
		ef := lr - ab
		ao := a.R - ef / 2
		theta2 := math.Acos(ao / a.R)
		theta := theta1 + theta2
		xc := a.X + a.R * math.Cos(theta)
		yc := a.Y + a.R * math.Sin(theta)
		p = append(p, Point{xc, yc})

        if ab < lr { //two intersections
			theta3 := math.Acos(ao / a.R)
			theta = theta3 - theta1
			xd := a.X + a.R * math.Cos(theta)
			yd := a.Y - a.R * math.Sin(theta)
			p = append(p, Point{xd, yd})
        }
    }
	return p
}

func MatchesBetweenArrays(satellites [3][]Point) (point []Point) {
	intersec := make(map[Point]bool)

	for index := 0; index < len(satellites); index++ {
		if len(satellites[index]) == 0 {
			return
		}
	}

	for _, item := range satellites[0] {
		intersec[item] = true
	}

	for _, item := range satellites[1] {
		if _, ok := intersec[item]; ok {
			point := append(point, item)
			_ = point
		}
	}

	for _, item := range satellites[2] {
		if _, ok := intersec[item]; ok {
			point = append(point, item)
		}
	}

	return
}


func GetLocation(satellites []SatelliteData) (x float32, y float32, found bool) {

	// Populate circles with Point (Loccation) and Radius (Distances)
	circle := [3]Circle{}
	for index := 0; index < len(satellites); index++ {
		var key string = satellites[index].Name
		circle[index].Point.X = float64(satellitesLocation[key].X)
		circle[index].Point.Y = float64(satellitesLocation[key].Y)
		circle[index].R = float64(satellites[index].Distance)
	}

	// Find every intersection between two circles (satellites)
	intersecBetweenTwoCircles := [3][]Point{}
	for index := 0; index < len(intersecBetweenTwoCircles); index++ {
		intersecBetweenTwoCircles[index] = Intersect(&circle[index % len(intersecBetweenTwoCircles)],
				  						  &circle[(index+1) % len(intersecBetweenTwoCircles)])
	}

	// Find common value between all inntersections otherwise return 
	var p []Point = MatchesBetweenArrays(intersecBetweenTwoCircles)
	if len(p) == 0 {
		x = 0
		y = 0
	} else {
		found = true
		x = float32(p[0].X)
		y = float32(p[0].Y)
	}

	return x,y,found
}