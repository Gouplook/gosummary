package main

func main() {

	m := make(map[int]int)
	for k, v := range m {
		if v == 34 {
			println(k)
		}
	}

}
