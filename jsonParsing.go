package main

// Data Structure를 안다면 Go struct에 담아서 처리하면 된다

type Record struct {
	Name string
	Age  int
	Car  []Car
}

type Car struct {
	Type   string
	Number string
}

// func main() {
// 	args := os.Args
// 	if len(args) == 1 {
// 		fmt.Println("need a filename!!")
// 		return
// 	}

// 	filename := args[1]
// 	fileData, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var pasredData map[string]interface{}
// 	json.Unmarshal([]byte(fileData), &pasredData)

// 	for key, value := range pasredData {
// 		fmt.Println("Key : ", key, ", Value : ", value)
// 	}
// }
