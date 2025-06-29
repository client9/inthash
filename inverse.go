package inthash

/*

// https://jeffhurchalla.com/2022/04/25/a-faster-multiplicative-inverse-mod-a-power-of-2/
uint32_t multiplicative_inverse(uint32_t a)
{
    assert(a%2 == 1);  // the inverse (mod 2<<32) only exists for odd values
    uint32_t x0 = (3*a)^2;
    uint32_t y = 1 - a*x0;
    uint32_t x1 = x0*(1 + y);
    y *= y;
    uint32_t x2 = x1*(1 + y);
    y *= y;
    uint32_t x3 = x2*(1 + y);
    return x3;
}

uint64_t multiplicative_inverse(uint64_t a)
{
    assert(a%2 == 1);  // the inverse (mod 2<<64) only exists for odd values
    uint64_t x0 = (3*a)^2;
    uint64_t y = 1 - a*x0;
    uint64_t x1 = x0*(1 + y);
    y *= y;
    uint64_t x2 = x1*(1 + y);
    y *= y;
    uint64_t x3 = x2*(1 + y);
    y *= y;
    uint64_t x4 = x3*(1 + y);
    return x4;
}
*/

func inverse32(a uint32) uint32 {
	if a&1 == 0 {
		panic("must be odd")
	}

	x0 := (3 * a) ^ 2
	y := 1 - a*x0
	x1 := x0 * (1 + y)
	y *= y
	x2 := x1 * (1 + y)
	y *= y
	x3 := x2 * (1 + y)
	return x3
}

func inverse64(a uint64) uint64 {
	if a&1 == 0 {
		panic("must be odd")
	}

	x0 := (3 * a) ^ 2
	y := 1 - a*x0
	x1 := x0 * (1 + y)
	y *= y
	x2 := x1 * (1 + y)
	y *= y
	x3 := x2 * (1 + y)
	y *= y
	x4 := x3 * (1 + y)
	return x4
}
