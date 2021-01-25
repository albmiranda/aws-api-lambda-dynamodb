// Package satellite stores struct and methods which handle satellites information
package satellite

import (
	"errors"
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

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func calculateThreeCircleIntersection(c [3]circle) (point, error) {

	p := point{}

	/* dx and dy are the vertical and horizontal distances between
	* the circle centers.
	 */
	dx := c[1].point.X - c[0].point.X
	dy := c[1].point.Y - c[0].point.Y

	/* Determine the straight-line distance between the centers. */
	d := math.Sqrt((dy * dy) + (dx * dx))

	/* Check for solvability. */
	if d > (c[0].R + c[1].R) {
		e := errors.New("circles do not intersect")
		return p, e
	}
	if d < math.Abs(c[0].R-c[1].R) {
		e := errors.New("one circle is contained in the other")
		return p, e
	}

	/* 'point 2' is the point where the line through the circle
	* intersection points crosses the line between the circle
	* centers.
	 */

	/* Determine the distance from point 0 to point 2. */
	a := ((c[0].R * c[0].R) - (c[1].R * c[1].R) + (d * d)) / (2.0 * d)

	/* Determine the coordinates of point 2. */
	point2x := c[0].point.X + (dx * a / d)
	point2y := c[0].point.Y + (dy * a / d)

	/* Determine the distance from point 2 to either of the
	* intersection points.
	 */
	h := math.Sqrt((c[0].R * c[0].R) - (a * a))

	/* Now determine the offsets of the intersection points from
	* point 2.
	 */
	rx := -dy * (h / d)
	ry := dx * (h / d)

	/* Determine the absolute intersection points. */
	intersectionPoint1x := toFixed(point2x+rx, 3)
	intersectionPoint2x := toFixed(point2x-rx, 3)
	intersectionPoint1y := toFixed(point2y+ry, 3)
	intersectionPoint2y := toFixed(point2y-ry, 3)

	// fmt.Println("INTERSECTION Circle1 AND Circle2: (", intersectionPoint1x, ",", intersectionPoint1y, ")", " AND (", intersectionPoint2x, ",", intersectionPoint2y, ")")

	/* Lets determine if circle 3 intersects at either of the above intersection points. */
	dx = intersectionPoint1x - c[2].point.X
	dy = intersectionPoint1y - c[2].point.Y
	d1 := math.Sqrt((dy * dy) + (dx * dx))

	dx = intersectionPoint2x - c[2].point.X
	dy = intersectionPoint2y - c[2].point.Y
	d2 := math.Sqrt((dy * dy) + (dx * dx))

	if math.Abs(d1-c[2].R) < 0.001 {
		p = point{
			X: intersectionPoint1x,
			Y: intersectionPoint1y,
		}
	} else if math.Abs(d2-c[2].R) < 0.001 {
		p = point{
			X: intersectionPoint2x,
			Y: intersectionPoint2y,
		}
	} else {
		e := errors.New("one of 3 circles didn't intersects")
		return p, e
	}

	return p, nil
}

// GetLocation receives data from all satellites and tries to find ship position
func GetLocation(satellites []Data) (x float32, y float32, found bool) {

	// Populate circles with point (Location) and Radius (Distances)
	circles := [3]circle{}
	for index := 0; index < len(satellites); index++ {
		var key string = satellites[index].Name
		circles[index].point.X = float64(satellitesLocation[key].X)
		circles[index].point.Y = float64(satellitesLocation[key].Y)
		circles[index].R = float64(satellites[index].Distance)
	}

	p, err := calculateThreeCircleIntersection(circles)
	if err != nil {
		return 0, 0, false
	}

	return float32(p.X), float32(p.Y), true
}
