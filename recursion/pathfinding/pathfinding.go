package pathfinding

// pathfinding recursion basecases
// 1. off the map
// 2. its a wall
// 3. its the end
// 4. if we have seen it before

type Point struct {
	x int
	y int
}

var dir [4][2]int = [4][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func Walk(
	maze []string,
	wall string,
	curr Point,
	end Point,
	seen [][]bool,
	path []Point,
) bool {

	// bc 1
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// bc 2
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	// bc 3
	if curr.x == end.x && curr.y == end.y {
		return true
	}

	// bc 4
	if seen[curr.y][curr.x] {
		return false
	}

	// pre condition
	path = append(path, curr)

	for i := 0; i < len(dir); i++ {
		if (Walk(maze, wall, Point{x: dir[i][0], y: dir[i][1]}, end, seen, path)) {
			return true
		}
	}

	path = path[:len(path)-1]

	seen[curr.y][curr.x] = true

	return false
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	return nil
}
