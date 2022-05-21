package main
import "fmt"

func main() {
		fmt.Println(discountSalary(50000))
		avg, e := average(10, 19)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(avg)
	
		
	}

	fmt.Println(calculateSalary("A", 162))

	oper := calculateValues(ave)
	r, e := oper(10, 8)
	fmt.Println(r, e)

	opera := calculateFood(cat)
	re, e := opera(10)
	fmt.Println(re, e)
}