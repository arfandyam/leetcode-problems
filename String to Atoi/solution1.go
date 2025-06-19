import (
	"strconv"
	"math"
)

func myAtoi(s string) int {
	var myNum string
    var zeroStr string
    isZero := false
	symbol := false

	for i := 0; i < len(s); i++ {
		asciiCode := int(s[i])
		// Kalau belum ada number
		if len(myNum) == 0 {
			if symbol {
				if asciiCode >= 48 && asciiCode <= 57 {
					myNum += string(s[i])
					symbol = false
				} else {
					break
				}
				// Kalau char berupa plus, continue
			} else if asciiCode == 43 {
				symbol = true
				continue
				// Kalau char berupa minus
			} else if asciiCode == 45 {
				myNum += string(s[i])
				symbol = true
			} else if asciiCode >= 48 && asciiCode <= 57 {
				myNum += string(s[i])
                // Kalau char berupa whitespace, continue
			} else if asciiCode == 32 {
				continue
			} else {
                myNum += "0"
				break
            }

		} else {
			// Kalau flag symbol true
			if symbol {
				// kalau char berupa 0, continue
				if asciiCode == 48 {
					continue
				} else if asciiCode >= 49 && asciiCode <= 57 {
					myNum += string(s[i])
                    symbol = false
				} else {
					break
				}
			} else if isZero {
				if asciiCode >= 49 && asciiCode <= 57 {
                    myNum += zeroStr
					myNum += string(s[i])
                    zeroStr = ""
					symbol = false
                    isZero = false
				} else if asciiCode == 48 {
					zeroStr += "0"
                    symbol = false
				} else {
                    break
                }
			} else {
                if asciiCode >= 49 && asciiCode <= 57 {
					myNum += string(s[i])
					symbol = false
				} else if asciiCode == 48 {
					zeroStr += "0"
                    isZero = true
                    symbol = false
				} else {
                    break
                }
            }
		}

		fmt.Println("myNum:", myNum)
        fmt.Println("zeroStr:", zeroStr)
        fmt.Println("isZero:", isZero)
        fmt.Println("symbol:", symbol)
	}

    if isZero {
        myNum += zeroStr
    }
	if len(myNum) == 0 {
		return 0
	}

	if myNum[0] == 48 {
		myNum = myNum[1:]
	}
	finalNum, _ := strconv.Atoi(myNum)

    fmt.Println(finalNum)

	if float64(finalNum) < math.Pow(-2, 31) {
		finalNum = int(math.Pow(-2, 31))
	} else if float64(finalNum) > (math.Pow(2, 31))-1 {
		finalNum = int((math.Pow(2, 31)) - 1)
	}

	return finalNum
}

