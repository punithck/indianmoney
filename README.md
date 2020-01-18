# indianmoney
  This library is used to format number in Indian money type

## Quick Start

```
go get github.com/punithck/indianmoney
```
** Examples **
`
0---> 0
1---> 1
12---> 12
123---> 123
1234---> 1,234
12345---> 12,345
123456---> 1,23,456
1234567---> 12,34,567
12345678---> 1,23,45,678
-123456789---> -12,34,56,789
1234567899---> 1,23,45,67,899
-12345678999---> -12,34,56,78,999
123456789999---> 1,23,45,67,89,999
0.0---> 0.00
1.1---> 1.10
12.12---> 12.12
123.123---> 123.123
1234.123---> 1,234.1230
12345.1234---> 12,345.12345
123456.12345---> 1,23,456.12345
1234567.12345---> 12,34,567.12
-12345678.123456---> -1,23,45,678.123
123456789.123456---> 12,34,56,789.12
-1234567899.123456---> -1,23,45,67,899.123456
12345678999.123456---> 12,34,56,78,999.123455
123456789999.123456---> 1,23,45,67,89,999.123459
`

example.go

```
package main

import (
	"fmt"

	"github.com/punithck/indianmoney"
)

func main() {
	//Integer value formatting
	fmt.Println("0--->", indianmoney.FormatMoneyInt(0))
	fmt.Println("1--->", indianmoney.FormatMoneyInt(1))
	fmt.Println("12--->", indianmoney.FormatMoneyInt(12))
	fmt.Println("123--->", indianmoney.FormatMoneyInt(123))
	fmt.Println("1234--->", indianmoney.FormatMoneyInt(1234))
	fmt.Println("12345--->", indianmoney.FormatMoneyInt(12345))
	fmt.Println("123456--->", indianmoney.FormatMoneyInt(123456))
	fmt.Println("1234567--->", indianmoney.FormatMoneyInt(1234567))
	fmt.Println("12345678--->", indianmoney.FormatMoneyInt(12345678))
	fmt.Println("-123456789--->", indianmoney.FormatMoneyInt(-123456789))
	fmt.Println("1234567899--->", indianmoney.FormatMoneyInt(1234567899))
	fmt.Println("-12345678999--->", indianmoney.FormatMoneyInt(-12345678999))
	fmt.Println("123456789999--->", indianmoney.FormatMoneyInt(123456789999))

	//Float value formatting
	fmt.Println("0.0--->", indianmoney.FormatMoneyFloat64(0.0, 2))
	fmt.Println("1.1--->", indianmoney.FormatMoneyFloat64(1.1, 2))
	fmt.Println("12.12--->", indianmoney.FormatMoneyFloat64(12.12, 2))
	fmt.Println("123.123--->", indianmoney.FormatMoneyFloat64(123.123, 3))
	fmt.Println("1234.123--->", indianmoney.FormatMoneyFloat64(1234.123, 4))
	fmt.Println("12345.1234--->", indianmoney.FormatMoneyFloat64(12345.12345, 5))
	fmt.Println("123456.12345--->", indianmoney.FormatMoneyFloat64(123456.12345, 5))
	fmt.Println("1234567.12345--->", indianmoney.FormatMoneyFloat64(1234567.123456, 2))
	fmt.Println("-12345678.123456--->", indianmoney.FormatMoneyFloat64(-12345678.123456, 3))
	fmt.Println("123456789.123456--->", indianmoney.FormatMoneyFloat64(123456789.123456, 2))
	fmt.Println("-1234567899.123456--->", indianmoney.FormatMoneyFloat64(-1234567899.123456, 6))
	fmt.Println("12345678999.123456--->", indianmoney.FormatMoneyFloat64(12345678999.123456, 6))
	fmt.Println("123456789999.123456--->", indianmoney.FormatMoneyFloat64(123456789999.123456, 6))
}
```

## To format int to INR type
```
indianmoney.FormatMoneyInt(12345) or
indianmoney.FormatMoneyInt64(12345)
```

## To format float64 to INR type
```
indianmoney.FormatMoneyFloat64(12345.12345, 5) //max_precision_supported  = 6
```
