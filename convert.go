package geezify

func ToGeez(num int) string {
	result := ""
	for i := 0; num != 0; i++ {
		var pad string
		evenGrouping := false
		if i != 0 && i%2 == 0 {
			pad = "፼"
			evenGrouping = true
		} else if i%2 != 0 {
			pad = "፻"
		}
		if m := num % 100; m != 0 || evenGrouping { // 00 with in the odd groupings are ignored
			result = pad + result
			// unit place 1 and middle place 10,000 unit values should also be handled
			if (m != 1 || i == 0) || (num > 1 && evenGrouping) {
				result = conv(m) + result
			}
		}
		num = num / 100
	}
	return result
}

var geezDigits = map[int]string{
	0: "",
	1: "፩", 2: "፪", 3: "፫", 4: "፬", 5: "፭", 6: "፮", 7: "፯", 8: "፰", 9: "፱",
	10: "፲", 20: "፳", 30: "፴", 40: "፵", 50: "፶", 60: "፷", 70: "፸", 80: "፹", 90: "፺",
}

// converts a base 10 value into geez
func conv(n int) string {
	return geezDigits[(n/10)*10] + geezDigits[n%10]
}
