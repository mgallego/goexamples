package main

import (
	"fmt"
	"math"
)

func main() {

}

func intTypes() {
	// Declare a variable with a type of int (and a default value of 0).
	var myInt int

	// Declare variable with type int64
	var largeInt int64

	// Use short declaration to create variable with type int and a value
	// of 3.
	i := 3

	// Use a type conversion if you want to use short declaration with other
	// integer types.
	u := uint64(4)

	// These are different ways of writing the same value.
	decInt := 1000         // Base 10 Notation          (no prefix)
	hexInt := 0x3E8        // Hexadecimal Notation      (Prefix: "0x")
	octInt := 01750        // Octal Notation            (Prefix: "0")
	octInt = 0o1750        // Alternate Octal Notation  (Prefix: "0o")
	binInt := 0b1111101000 // Binary Notation           (Prefix: "0b")

	// Optional underscores can be used to separate digits for readability.
	withSep := 1_000
	hexWithSep := 0x3_E_8

	// Notation is purely cosmetic, and does not affect value:
	fmt.Println(decInt, hexInt, octInt, binInt, withSep, hexWithSep)
	// Output: 1000 1000 1000 1000 1000 1000

	// Negative integer literals:
	negativeInt := -10
	negativeHexInt := -0xa
}

func floatingPointTypes() {
	// Declare a variable doubleFloat with type float64 (and a default
	// value of 0).
	var doubleFloat float64

	// Declare a variable with type float32.
	var singleFloat float32

	// Use short declaration to create a variable f with type float64 and a
	// value of 12.1.
	f := 12.1

	// Use a type conversion if you want to use short declaration with float32.
	g := float32(12.1)

	// The following are different ways to write the same value.
	floatVal = 12.         // fractional part is optional
	floatVal = 12e0        // decimal point not required with exponent
	floatVal = 012.        // leading zero is fine
	floatVal = .12e+2      // integer part is optional
	floatVal = 1_2.        // underscores allowed in floats, too
	floatVal = float64(12) // explicit type for otherwise integer literal

	fmt.Println(floatVal, myFloat)
	// Output: 12 12

	// Floats can be negative.
	negativeFloat := -12.0

	fmt.Println(negativeFloat)
	// Output: -12
}

func specialFLoatingPointValues() {
	f := 2.0
	// Make an "infinitely huge" number.
	posInf := math.Pow(f, 10_000)
	fmt.Println(posInf) // output: +Inf

	// Make an "infinitely small" number.
	negInf := f - math.Pow(f, 10_000)
	fmt.Println(negInf) // output: -Inf

	// The second argument to math.IsInf lets you check for a specific infinity:
	// <0: negative infinity
	// >0: positive infinity
	// 0: either infinity
	fmt.Println(math.IsInf(posInf, 0))  // true
	fmt.Println(math.IsInf(negInf, 1))  // false
	fmt.Println(math.IsInf(negInf, -1)) // true

	// Make an invalid number.
	notANumber := math.Log(-f)
	fmt.Println(notANumber) // output: NaN

	fmt.Println(math.IsNaN(notANumber)) // true
}
func complexNumbers() {
	cmplxValue1 := complex(1.1, 2.2)                   // complex128
	cmplxValue2 := complex(float32(1.1), float32(2.2)) // complex64

	cmplxValue3 := complex(1, 2)       // complex128, equivalent to 1+2i
	cmplxValue4 := complex(0, -3)      // complex128, equivalent to 0-3i
	cmplxValue5 := complex(-3.3, -4.4) // complex128, equivalent to -3.3-4.4i

	cmplxValue1 = 1.1 + 2.2i  // equivalent to complex(1.1, 2.2)
	cmplxValue3 = 1 + 2i      // equivalent to complex(1,2)
	cmplxValue4 = -3i         // equivalent to complex(0, -3)
	cmplxValue5 = -3.3 - 4.4i // equivalent to complex(-3.3,-4.4)

	cmplxValue1 := complex(1.1, 2.2)
	fmt.Println(real(c1), imag(c1)) // "1.1 2.2"

}

func mathematicalOperators() {
	a := 7
	b := 3

	var i int
	i = a + b // 10
	i = a - b // 4
	i = b - a // -4
	i = a * b // 21
	i = a % b // 1

	// Divide by zero can be caught by the compiler if the 0 is a literal.
	//c = a / 0  // error: invalid operation: division by zero

	// Otherwise, divide by zero is a runtime panic.
	zero := 0
	c = a / zero // panic: runtime error: integer divide by zero

	// Float divide by zero does not panic and will result in an infinite
	// value.
	f := 1.0
	f = f / float64(zero) // +Inf

	u := uint64(1)
	u = u + i // error: mismatched types uint64 and int
	// Type conversion is needed to allow different types to work together.
	u = u + uint64(i) // 2
}
