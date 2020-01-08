package main

func main() {

	phone := map[string]string{
		"Jack": "01112223333",
		"Nick": "09944995343",
		"MacB": "03344563456",
	}
	_ = phone

	product := map[int]string{
		100: "Available",
		200: "Unavailable",
		300: "Available",
	}
	_ = product

	phones := map[string][]string{
		"Jack": []string{"01112223333", "02222223333", "03332223333"},
		"Nick": []string{"09992223333", "05552223333", "04442223333"},
		"MacB": []string{"08882223333", "06662223333", "08882223333"},
	}
	_ = phones

	cart := map[int]map[int]int{
		100: map[int]int{1: 10, 2: 20},
		200: map[int]int{3: 30, 4: 40},
		300: map[int]int{5: 50, 6: 60, 7: 70},
	}
	_ = cart
}
