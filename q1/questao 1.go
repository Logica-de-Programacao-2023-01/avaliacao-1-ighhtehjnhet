package main

import "fmt"

func main() {

	var weight int

	if weight <= 0 {
		return false, fmt.Errorf("te, que ser maior que o peso da melancia")

	} else if weight%2 == 0 && weight > 2 {
		return true, nil
	} else {
		return false, nil

	}
	return false, nil
}
