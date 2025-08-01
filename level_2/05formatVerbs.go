package main

import "fmt"

func main_05() {

	// --- General formatting verbs ---
	// %v prints value in default format
	// %#v prints value in Go-syntax format
	// %T prints type of value
	// %% prints % sign

	float_num := 15.5
	fmt.Printf("float_num = %v (default)\n", float_num)
	fmt.Printf("float_num = %#v (Go-syntax)\n", float_num)
	fmt.Printf("float_num has type %T\n", float_num)
	fmt.Printf("%% - percent sign\n")
	another_float_num := 45_678.9
	fmt.Printf("another_float_num = %v\n\n", another_float_num)

	str := "Hello"
	fmt.Printf("str = %v (default)\n", str)
	fmt.Printf("str = %#v (Go-syntax)\n", str)
	fmt.Printf("str har type %T\n\n", str)

	// --- Integer formatting verbs ---
	// %b base 2
	// %#b base 2 with 0b
	// %o base 8
	// %#o base 8 with 0
	// %O base 8 with 0o
	// %#O base 8 with 0o0
	// %x base 16 in lowercase
	// %X base 16 in uppercase
	// %#x base 16 with 0x in lowercase
	// %#X base 16 with 0X in uppercase
	// %d base 10
	// %+d base 10 with sign
	// %4d pads number with spaces to width of 4
	// %-4d left-justifies number in width of 4
	// %04d pads number with zeros to width of 4
	int_num := 31
	fmt.Printf("int_num = %b or %#b (base 2)\n", int_num, int_num)
	fmt.Printf("int_num = %o or %#o (base 8)\n", int_num, int_num)
	fmt.Printf("int_num = %o or %#o (base 8)\n", int_num, int_num)
	fmt.Printf("int_num = %O or %#O (base 8)\n", int_num, int_num)
	fmt.Printf("int_num = %x or %#x(base 16)\n", int_num, int_num)
	fmt.Printf("int_num = %X or %#X(base 16)\n", int_num, int_num)
	fmt.Printf("int_num = %d (base 10)\n", int_num)
	fmt.Printf("int_num = %+d (with sign)\n", int_num)
	fmt.Printf("int_num = %4d (width 4 with spaces, default right)\n", int_num)
	fmt.Printf("int_num = %-4d (width 4 with spaces, left)\n", int_num)
	fmt.Printf("int_num = %04d (width 4 with zeros)\n\n", int_num)

	// --- String formatting verbs ---
	// %s prints value as plain string
	// %q prints value as double-quoted string
	// %8s prints value as right-justified plain string with width 8
	// %-8s prints value as left-justified plain string with width 8
	// %x prints value as hex dump of byte values
	// % x prints value as hex dump with spaces
	txt := "hello"
	fmt.Printf("txt = %s (plain string)\n", txt)
	fmt.Printf("txt = %q (double-quoted string)\n", txt)
	fmt.Printf("txt = %8s (width 8, default right)\n", txt)
	fmt.Printf("txt = %-8s (width 8, left)\n", txt)
	fmt.Printf("txt = %x (byte values)\n", txt)
	fmt.Printf("txt = % x (bite values with spaces)\n\n", txt)

	// --- Boolean formating verbs ---
	// %t prints true or false
	t := true
	f := false
	fmt.Printf("t = %t\n", t)
	fmt.Printf("f = %t\n", f)
	fmt.Printf("t = %v\n\n", f)

	// --- Float formatting verbs ---
	// %e scientific notation with exponent
	// %f decimal point, no exponent
	// %.2f default width and precision 2
	// %6.3f width 6 and precision 3
	// %g exponent as needed, only necessary digits
	float_num2 := 9.16368
	fmt.Printf("float_num2 = %e (with exponent)\n", float_num2)
	fmt.Printf("float_num2 = %f (no exponent)\n", float_num2)
	fmt.Printf("float_num2 = %.2f (default width and precision 2)\n", float_num2)
	fmt.Printf("float_num2 = %6.3f (width 6 and precision 3)\n", float_num2)
	fmt.Printf("float_num2 = %g (exponent as needed)\n", float_num2)
}
