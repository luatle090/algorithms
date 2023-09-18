package algorithms

// import "sort"

// type Point struct {
// 	x, y float32
// }

// func graham(points []Point) []Point {
// 	// check các point trong slice?
// 	n := len(point)
// 	if n < 3 {
// 		return point
// 	}
// 	pointLowestY, i := GetPointLowestYAxis(points)

// 	sort.Slice(points, func(i, j int) bool {
// 		return points[i] < points[j]
// 	})
// 	return
// }

// func isLeftTurn() {

// }

// func GetPointLowestYAxis(points []Point) (Point, int) {
// 	pointLowestY := points[0]
// 	index := 0
// 	for i := 1; i < len(points); i++ {
// 		if points[i].y < pointLowestY.y {
// 			pointLowestY = points[i]
// 			index = i
// 		}
// 	}
// 	return pointLowestY, index
// }
