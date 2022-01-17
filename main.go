package main

import (
	"fmt"
	"math"
)

func main() {
	var multiplier1, multiplier2 int16

	fmt.Println("введите значение первого множителя:")
	fmt.Scan(&multiplier1)

	fmt.Println("введите значение второго множителя:")
	fmt.Scan(&multiplier2)

	fmt.Printf("%v %v результат: ", multiplier1, multiplier2)
	if multiplier1 > 0 && multiplier2 > 0 {
		result := uint32(multiplier1) * uint32(multiplier2)
		switch {
		case result <= math.MaxUint8:
			{
				temp := uint8(result)
				fmt.Printf("%T\n", temp)
				break
			}
		case result <= math.MaxUint16:
			{
				temp := uint16(result)
				fmt.Printf("%T\n", temp)
				break
			}
		case result <= math.MaxUint32:
			{
				temp := uint32(result)
				fmt.Printf("%T\n", temp)
				break
			}
		}
	} else {
		result := int32(multiplier1) * int32(multiplier2)
		switch {
		case result > math.MinInt8 && result <= math.MaxInt8:
			{
				temp := int8(result)
				fmt.Printf("%T\n", temp)
				break
			}
		case result > math.MinInt16 && result <= math.MaxInt16:
			{
				temp := int16(result)
				fmt.Printf("%T\n", temp)
				break
			}
		case result > math.MinInt32 && result <= math.MaxInt32:
			{
				temp := int32(result)
				fmt.Printf("%T\n", temp)
				break
			}
		}
	}

}
