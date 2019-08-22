package main

var (
	cols = 20
	rows = cols
)

func main() {

	splitLine()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			print("|")

			print(" o ")

			if j == cols-1 {
				print("|")
				print("\n")
			}
		}
		splitLine()
	}

}

func splitLine() {
	for num := 0; num < cols*4+1; num++ {
		print("-")
	}
	println("")
}
