package reverse

func Reverse(s []int) {
	var i, j = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func Reverse2(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse3(s []int) {
	for i, l := 0, len(s); i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
