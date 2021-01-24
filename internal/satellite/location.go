// Package satellite stores struct and methods which handle satellites information
package satellite

import (
	"math"
)

type point struct {
	X float64
	Y float64
}

type circle struct {
	point
	R float64
}

func intersect(a *circle, b *circle) (p []point) {
	dx, dy := b.X-a.X, b.Y-a.Y
	lr := a.R + b.R                //radius and
	dr := math.Abs(a.R - b.R)      //radius difference
	ab := math.Sqrt(dx*dx + dy*dy) //center distance

	if ab <= lr && ab > dr {
		theta1 := math.Atan(dy / dx)
		ef := lr - ab
		ao := a.R - ef/2
		theta2 := math.Acos(ao / a.R)
		theta := theta1 + theta2
		xc := a.X + a.R*math.Cos(theta)
		yc := a.Y + a.R*math.Sin(theta)
		p = append(p, point{xc, yc})

		if ab < lr { //two intersections
			theta3 := math.Acos(ao / a.R)
			theta = theta3 - theta1
			xd := a.X + a.R*math.Cos(theta)
			yd := a.Y - a.R*math.Sin(theta)
			p = append(p, point{xd, yd})
		}
	}
	return p
}

func matchesBetweenArrays(satellites [3][]point) (p []point) {
	intersec := make(map[point]bool)

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
			p := append(p, item)
			_ = p
		}
	}

	for _, item := range satellites[2] {
		if _, ok := intersec[item]; ok {
			p = append(p, item)
		}
	}

	return
}

// GetLocation receives data from all satellites and tries to find ship position
func GetLocation(satellites []Data) (x float32, y float32, found bool) {

	// Populate circles with point (Location) and Radius (Distances)
	circle := [3]circle{}
	for index := 0; index < len(satellites); index++ {
		var key string = satellites[index].Name
		circle[index].point.X = float64(satellitesLocation[key].X)
		circle[index].point.Y = float64(satellitesLocation[key].Y)
		circle[index].R = float64(satellites[index].Distance)
	}

	// Find every intersection between two circles (satellites)
	intersecBetweenTwocircles := [3][]point{}
	for index := 0; index < len(intersecBetweenTwocircles); index++ {
		intersecBetweenTwocircles[index] = intersect(&circle[index%len(intersecBetweenTwocircles)],
			&circle[(index+1)%len(intersecBetweenTwocircles)])
	}

	// Find common value between all inntersections otherwise return
	var p []point = matchesBetweenArrays(intersecBetweenTwocircles)
	if len(p) == 0 {
		x = 0
		y = 0
	} else {
		found = true
		x = float32(p[0].X)
		y = float32(p[0].Y)
	}

	return
}
