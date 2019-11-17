package main

import ( "fmt"
"strings" 
)	

func main() {
	fmt.Println(strings.Contains("Kanokpol","Kanokpol")) //true จริง
	fmt.Println(strings.Contains("Kanokpol","kanokpol")) //false ไม่จริง

	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	
	
}
	

