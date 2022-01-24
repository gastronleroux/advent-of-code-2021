package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func absDiffInt(x, y int) int {
	diff := x - y
	if diff < 0 {
		diff *= -1
	}
	return diff
}

type Coord struct {
	x, y int
}

func (c *Coord) setFromString(coord, value string) {
	var err error
	if coord == "x" {
		c.x, err = strconv.Atoi(value)
	} else if coord == "y" {
		c.y, err = strconv.Atoi(value)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func (c1 Coord) betweenHorAndVer(c2 Coord) []Coord {
	var from Coord
	var to Coord
	var pointsToCheck []Coord
	if c1.x == c2.x && c1.y > c2.y || c1.y == c2.y && c1.x > c2.x {
		from = c2
		to = c1
	} else if c1.x == c2.x && c1.y <= c2.y || c1.y == c2.y && c1.x <= c2.x {
		from = c1
		to = c2
	} else {
		return pointsToCheck
	}
	for i := from.x; i <= to.x; i++ {
		for j := from.y; j <= to.y; j++ {
			pointsToCheck = append(pointsToCheck, Coord{x: i, y: j})
		}
	}
	return pointsToCheck
}

func (c1 Coord) betweenDia(c2 Coord) []Coord {
	var from Coord
	var to Coord
	var pointsToCheck []Coord
	if absDiffInt(c1.x, c2.x) == absDiffInt(c1.y, c2.y) {
		if c1.x < c2.x {
			from = c1
			to = c2
		} else {
			from = c2
			to = c1
		}
	} else {
		return pointsToCheck
	}
	for i := 0; i <= to.x-from.x; i++ {
		x := from.x + i
		var y int
		if from.y < to.y {
			y = from.y + i
		} else {
			y = from.y - i
		}
		pointsToCheck = append(pointsToCheck, Coord{x: x, y: y})
	}
	return pointsToCheck
}

func part1(points map[Coord]int, from, to Coord) {
	for _, coordBetween := range from.betweenHorAndVer(to) {
		if points[coordBetween] < 2 {
			points[coordBetween]++
		}
	}
}

func part2(points map[Coord]int, from, to Coord) {
	for _, coordBetween := range append(from.betweenDia(to), from.betweenHorAndVer(to)...) {
		if points[coordBetween] < 2 {
			points[coordBetween]++
		}
	}
}

func main() {
	var partFunction func(map[Coord]int, Coord, Coord)
	if len(os.Args) > 1 && os.Args[1] == "2" {
		partFunction = part2
	} else {
		partFunction = part1
	}
	points := make(map[Coord]int)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var from, to Coord

		coords := strings.Split(scanner.Text(), " -> ")
		splitCoords1 := strings.Split(coords[0], ",")
		splitCoords2 := strings.Split(coords[1], ",")

		from.setFromString("x", splitCoords1[0])
		from.setFromString("y", splitCoords1[1])
		to.setFromString("x", splitCoords2[0])
		to.setFromString("y", splitCoords2[1])

		partFunction(points, from, to)
	}

	var sum int
	for _, v := range points {
		if v == 2 {
			sum++
		}
	}

	fmt.Println(sum)
}
