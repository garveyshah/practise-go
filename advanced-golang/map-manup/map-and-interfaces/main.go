package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["key"] = "value"
	fmt.Println(m["key"], "\n")
	m2 := map[string]interface{}{
		"key2": "value2",
	}
	fmt.Println(m2["key2"], "\n")

	m3 := make(map[string]interface{})
	m3["key3"] = "value3"
	m3["key3.1"] = "value3.1"
	fmt.Println(m3["key3"],"\n",m3["key3.1"],"\n")

	// Nested map
	m4 := make(map[string]interface{})
	m4["key1"] = make(map[string]interface{})
	m4["key1"].(map[string]interface{})["key2"] = "value"

	// Note: before accessing the nested map you have to 
	//assert that the value is actually a map, in this case using `.`

	
}
