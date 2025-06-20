import "math"

func myAtoi(s string) int {
	numberAscii := map[int]int{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}

	// multiplier := 1
	num := 0
	var isExist bool
	var minus bool

	for i := 0; i < len(s); i++ {
		asciiCode := int(s[i])
		// Kalau characternya whitespace
		if asciiCode == 32 {
			if num == 0 && !isExist {
				continue
			} else {
				break
			}
			// Kalau characternya -
		} else if asciiCode == 45 {
			if num == 0 && !isExist {
				minus = true
				isExist = true
			} else {
				break
			}
			//Kalau characternya +
		} else if asciiCode == 43 {
			if num == 0 && !isExist {
                isExist = true
				continue
			} else {
				break
			}
			// Kalau characternya 0
		} else if asciiCode == 48 {
			if num == 0 {
				isExist = true
				continue
			} else {
				num = (num * 10)
			}
		} else if asciiCode >= 49 && asciiCode <= 57 {
			if num == 0 {
				num = 1
				num *= numberAscii[asciiCode]
			} else {
				num = (num * 10) + numberAscii[asciiCode]
			}
		} else {
			break
		}

        if float64(num) > math.Pow(2, 31){
            break
        }

		fmt.Println("num:", num)
	}

	if minus {
		num *= -1
	}

	if float64(num) < math.Pow(-2, 31) {
		num = int(math.Pow(-2, 31))
	} else if float64(num) > (math.Pow(2, 31) - 1) {
		num = int(math.Pow(2, 31) - 1)
	}

	return num
}

