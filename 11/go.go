package main

import "fmt"

const INPUT = 9445

func Setup(serial int) *[301][301]int {
	grid := [301][301]int{}
	for x := 1; x < 301; x++ {
		rid := x + 10
		for y := 1; y < 301; y++ {
			p := (rid*y + serial) * rid
			p = p/100%10 - 5
			grid[x][y] = p
		}
	}
	return &grid
}

func SumRect(grid *[301][301]int, x, y, dx, dy int) int {
	total := 0
	// faster to iterate over slices and/or range?
	for xx := x; xx < x+dx; xx++ {
		for yy := y; yy < y+dy; yy++ {
			total += grid[xx][yy]
		}
	}
	return total
}

func Part1() {
	grid := Setup(INPUT)
	best_tp := -46
	best_x := 0
	best_y := 0
	tp := 0
	for x := 1; x < 298; x++ {
		for y := 1; y < 298; y++ {
			tp = SumRect(grid, x, y, 3, 3)
			if tp > best_tp {
				best_tp = tp
				best_x = x
				best_y = y
			}
		}
	}
	fmt.Printf("(%d,%d): %d\n", best_x, best_y, best_tp)
}

func Part2() {
	grid := Setup(INPUT)
	best_tp := 0
	best_x := 0
	best_y := 0
	best_n := 0
	n := 300
	tp0 := 0
	tp := 0
	for n*n*5 > best_tp {
		for x := 1; x < 302-n; x++ {
			if x == 1 {
				tp0 = SumRect(grid, x, 1, n, n)
			} else {
				tp0 -= SumRect(grid, x-1, 1, 1, n)
				tp0 += SumRect(grid, x+n-1, 1, 1, n)
			}
			tp = tp0
			for y := 1; y < 302-n; y++ {
				if y > 1 {
					tp -= SumRect(grid, x, y-1, n, 1)
					tp += SumRect(grid, x, y+n-1, n, 1)
				}
				if tp > best_tp {
					best_tp = tp
					best_x = x
					best_y = y
					best_n = n
				}
			}
		}
		n -= 1
	}
	fmt.Printf("(%d,%d,%d): %d\n", best_x, best_y, best_n, best_tp)
}

func IncrementRowSums(grid, row_sums *[301][301]int, n int) {
	// assumes row_sums has already been calculated for rows of n-1 length
	for add_x := n; add_x < 301; add_x++ {
		a := grid[add_x][:]
		b := row_sums[add_x-n+1][:]
		for y := 1; y < 301; y++ {
			b[y] += a[y]
		}
	}
}

func Part2a() {
	grid := Setup(INPUT)
	row_sums := &[301][301]int{} // sum of n-length row fragment starting at [x][y]
	best_tp := 0
	best_x := 0
	best_y := 0
	best_n := 0
	tp := 0
	for n := 1; n < 301; n++ {
		IncrementRowSums(grid, row_sums, n)
		for x := 1; x <= 301-n; x++ {
			rows := row_sums[x][:]
			tp = 0
			// total power of an nxn square is the sum of the n row_sums that compose it
			for _, v := range rows[1:n] {
				tp += v
			}
			for sub_y, y := 0, n; y < 301; sub_y, y = sub_y+1, y+1 {
				// as our window moves down, 1 row_sum drops out and 1 row gets added
				tp += rows[y] - rows[sub_y]
				if tp > best_tp {
					best_tp = tp
					best_x = x
					best_y = sub_y + 1
					best_n = n
				}
			}
		}
	}
	fmt.Printf("(%d,%d,%d): %d\n", best_x, best_y, best_n, best_tp)
}

func main() {
	Part1()
	Part2a()
}
