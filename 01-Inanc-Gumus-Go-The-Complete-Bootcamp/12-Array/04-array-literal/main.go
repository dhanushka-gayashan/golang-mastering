package main

func main() {

	{
		var books [4]string

		books[0] = "Kafka's Revenge"
		books[1] = "Stay Golden"
		books[2] = "Everythingship"
		books[3] += "Kafka's Revenge 2nd Edition"

		_ = books
	}

	{
		var books = [4]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		_ = books
	}

	{
		books := [4]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		_ = books
	}

	{
		books := [4]string{
			"Kafka's Revenge",
			"Stay Golden",
		}

		_ = books
	}

	// Ellipsis [...]
	{
		books := [...]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		_ = books
	}
}
