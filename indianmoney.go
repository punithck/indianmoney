package indianmoney

import (
	"fmt"
	"math"
	"strings"
)

// FormatMoneyInt returns comma separated number(Indian Money Format) in string
func FormatMoneyInt(num int) string {
	var isNegative bool
	if num < 0 {
		isNegative = true
	}
	num = int(math.Abs(float64(num)))              //negative int to positive int
	numInStr := fmt.Sprint(num)                    //int to string
	numDigits := []rune(numInStr)                  //string to []rune
	numDigitsInRevOrder := reverseRunes(numDigits) //reverse above array
	var commaSapNumStrInRevOrder string
	for i := 0; i < len(numDigitsInRevOrder); i++ {
		if i > 2 && i%2 == 1 {
			commaSapNumStrInRevOrder += ","
		}
		commaSapNumStrInRevOrder += string(numDigitsInRevOrder[i])
	}
	commSapNumStr := string(reverseRunes([]rune(commaSapNumStrInRevOrder)))
	if isNegative { //prefix with sign '-'
		return "-" + commSapNumStr
	}
	return commSapNumStr
}

// FormatMoneyFloat64 returns comma separated number(Indian Money Format) in string
// If precision is zero, it acts returns comma separated decimal part of number
func FormatMoneyFloat64(num float64, precision uint) string {
	decimalPart := FormatMoneyInt(int(num))
	if precision == 0 {
		return decimalPart
	}
	numInStr := fmt.Sprintf("%."+fmt.Sprint(precision)+"f", num)
	floatParts := strings.Split(numInStr, ".")
	return decimalPart + "." + floatParts[1]
}

func FormatMoneyInt64(num int64) string {
	return FormatMoneyInt(int(num))
}

//Internal function to reverse the array
func reverseRunes(chrs []rune) []rune {
	for i := 0; i < len(chrs)/2; i++ {
		j := len(chrs) - i - 1
		chrs[i], chrs[j] = chrs[j], chrs[i]
	}
	return chrs
}

