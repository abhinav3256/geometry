package main

import "fmt"

func main() {
	var area int
	var perimeter int
	var operation int
	fmt.Println("Enter shape- options: [1] Rectangle, [2] Circle, [3] Triangle")
	var shape int
	fmt.Scanln(&shape)

	if shape == 1 {
		fmt.Println("Enter operation type- options: [1] Area , [2] Perimeter")

		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Println("Enter Retangle Length")
			var len int
			fmt.Scanln(&len)
			fmt.Println("Enter Retangle width")
			var width int
			fmt.Scanln(&width)
			area = rectangleArea(len, width)
			fmt.Println("Area of Triangle is:", area)
		} else if operation == 2 {
			fmt.Println("Enter Retangle Length")
			var len int
			fmt.Scanln(&len)
			fmt.Println("Enter Retangle width")
			var width int
			fmt.Scanln(&width)
			perimeter = rectanglePerimeter(len, width)
			fmt.Println("Perimeter of Rectangle is :", perimeter)
		} else {
			fmt.Println("Invalid Operation")
		}
	} else if shape == 2 {
		fmt.Println("Enter operation type- options: [1] Area , [2] Perimeter")

		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Println("Enter Circle radius")
			var radius float32
			fmt.Scanln(&radius)

			var cArea = circleArea(radius)
			fmt.Println("Area of Triangle is:", cArea)
		} else if operation == 2 {
			fmt.Println("Enter Circle radius")
			var radius float32
			fmt.Scanln(&radius)

			var cPerimeter = circlePerimeter(radius)
			fmt.Println("Perimeter of Circle is:", cPerimeter)
		} else {
			fmt.Println("Invalid Operation")
		}
	} else if shape == 3 {
		fmt.Println("Enter operation type- options: [1] Area , [2] Perimeter")

		fmt.Scanln(&operation)

		if operation == 1 {
			fmt.Println("Enter base of triangle")
			var base int
			fmt.Scanln(&base)

			fmt.Println("Enter height of triangle")
			var height int
			fmt.Scanln(&height)

			area = triangleArea(base, height)
			fmt.Println("Area of Triangle is:", area)
		} else if operation == 2 {
			fmt.Println("Enter side A")
			var sideA int
			fmt.Scanln(&sideA)

			fmt.Println("Enter side B")

			var sideB int
			fmt.Scanln(&sideB)

			fmt.Println("Enter side C")
			var sideC int
			fmt.Scanln(&sideC)

			perimeter = trianglePerimeter(sideA, sideB, sideC)
			fmt.Println("Perimeter of Triangle is:", perimeter)
		} else {
			fmt.Println("Invalid Operation")
		}
	} else {
		fmt.Println("Invalid Shape")
	}

}
func rectangleArea(len, width int) int {
	return len * width
}

func rectanglePerimeter(len, width int) int {
	return 2 * (len + width)

}

func circleArea(radius float32) float32 {
	var radiusSquare float32 = radius * radius
	var area = 3.14 * radiusSquare
	return area
}

func circlePerimeter(radius float32) float32 {

	var perimeter = 2 * 3.14 * radius
	return perimeter
}

func triangleArea(base, height int) int {

	var area = (base * height) / 2
	return area
}

func trianglePerimeter(a, b, c int) int {

	return a + b + c

}
