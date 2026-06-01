package main

import "fmt"

// func divide(numerator, denominator float64)(float64, error){
// 	if denominator == 0 {
// 		return 0, fmt.Errorf("Cannot divide by 0")
// 	}

// 	return numerator / denominator, nil
// }

// func main(){
// 	result, err := divide(10,20)

// 	if err != nil{
// 		fmt.Println("Error: ",err)
// 	} else {
// 		fmt.Println("Result: ", result)
// 	}
// }


type Rectangle struct {
	Width, Height float64
}

// Value Receiver -> 
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

// Pointer Receiver -> actual value is directly passed
func (r *Rectangle) Scale(factor float64){
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

func main(){
	rect := Rectangle{Width: 10, Height:5}

	fmt.Println("Initial Area:", rect.Area()) // 50

	rect.Scale(2) // Width: 20, Height: 10
	fmt.Println("New Area:", rect.Area()) // 200
}
