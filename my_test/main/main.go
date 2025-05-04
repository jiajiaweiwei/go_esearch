package main

func main() {
	m := make(map[string]string)
	m["1"] = "1"
	value, ok := m["2"]
	if !ok {
		println("nothing.")
	}
	println(value)
}
