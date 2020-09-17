package utils

func BubbleSort(elements []int) []int {
	flag := true

	for flag {
		flag = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i+1] < elements[i] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				flag = true
			}
		}
	}
	return elements
}
