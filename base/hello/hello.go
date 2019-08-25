package hello

var (
	cols   = 5
	rows   = cols
	status = []int{
		1, 0, 0, 1, 1,
		0, 1, 0, 1, 0,
		0, 0, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1,
	}
)

func Hello() {

	if len(status) != cols*rows {
		println("要素が一致しません、設定し直してください")
		return
	}

	splitLine()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			print("|")

			openOrClose(status[i*cols+j])

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

func openOrClose(status int) {
	if status == 0 {
		print(" x ")
	} else {
		print(" o ")
	}
}
