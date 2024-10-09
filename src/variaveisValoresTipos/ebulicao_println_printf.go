package main

import "fmt"

const ebulicaoF  = 212.0
// ebulicaoF := 212.0 // irá apresentar um erro, pois criar variável por atriução precisar ser em um bloco

func main() {
	tempF := ebulicaoF  // cria variável atribui valor por inferencia, precisa estar em um bloco
	tempC := (tempF - 32) * 5 / 9 // cria variável atribui valor por inferencia

	// fmt.Println("A tempenratura de ebulição da água em ºF = ", tempF)
	// fmt.Println("A tempenratura de ebulição da água em ºC = ", tempC)
	
	// print format %g -> float
	fmt.Printf("A tempenratura de ebulição da água em ºF = %g  e a tempenratura de ebulição da água em ºC = %g", tempF, tempC)

/*

Integer:
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"


Floating-point and complex constituents:
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

String and slice of bytes (treated equivalently with these verbs):
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte

Slice:
%p	address of 0th element in base 16 notation, with leading 0x

Pointer:
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.

The default format for %v is:
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p




*/



}