package main

import "fmt"

func main() {

	str1 := "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Printf("%s\n", str1)

	// Raw String
	str1 = `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Printf("%s\n", str1)
}
