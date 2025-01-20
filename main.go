package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

// calculateCoordinates returns x,y coordinates based on the position n
func calculateCoordinates(n int) (x, y float64) {
	// GOAL: Calculate polar coordinates and convert to Cartesian
	r := math.Sqrt(float64(n))
	theta := math.Sqrt(float64(n))

	// GOAL: Convert to Cartesian coordinates
	x = r * math.Cos(2*math.Pi*theta)
	y = -r * math.Sin(2*math.Pi*theta)

	return x, y
}

func main() {
	// GOAL: Get input filename and debug flag from command line
	inputFile := flag.String("file", "input.txt", "Path to input file containing comma-separated binary digits")
	debug := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()

	// GOAL: Read the input file and process the binary digits
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// GOAL: Process file character by character and count digits
	reader := bufio.NewReader(file)

	// CLAIM: i will count the number of digits (0 or 1) found, starting at 1
	i := 0
	for {
		char, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// GOAL: Print digits and calculate coordinates for 1s
		if char == '0' || char == '1' {
			i++
			if *debug {
				fmt.Printf("%d: %s", i, string(char))
			}
			if char == '1' {
				x, y := calculateCoordinates(i)
				if *debug {
					fmt.Printf(" at coordinates (%.15f, %.15f)", x, y)
				}
				fmt.Printf("%.12f,%.12f\n", x, y)
			}
			if *debug {
				fmt.Println()
			}
		}
	}

	if *debug {
		fmt.Printf("Total digits processed: %d\n", i)
	}
}
