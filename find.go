package main

import (
	. "github.com/ktye/wg/module"
)

type f3i = func(int32, int32, int32) int32

func Fnd(x, y K) K { // x?y
	r := K(0)
	xt, yt := tp(x), tp(y)
	if xt < 16 {
		if yt == Tt {
			return grp(x, y)
		} else {
			return deal(x, y)
		}
	}
	if xt > Lt {
		if xt == Tt {
			trap() //nyi t?..
		}
		r = x0(x)
		return Atx(r, Fnd(r1(x), y))
	} else if xt == yt {
		yn := nn(y)
		r = mk(It, yn)
		rp := int32(r)
		e := ep(r)
		if xt == Lt {
			yp := int32(y)
			for rp < e {
				SetI32(rp, fndl(x, x0(K(yp))))
				rp += 4
				yp += 8
			}
		} else {
			for i := int32(0); i < yn; i++ {
				yi := ati(rx(y), i)
				SetI32(rp, fnd(x, yi, xt-16))
				dx(yi)
				rp += 4
			}
		}
	} else if xt == yt+16 {
		r = Ki(fnd(x, y, yt))
	} else if xt == Lt {
		r = Ki(fndl(x, rx(y)))
	} else if yt == Lt {
		return Ecr(18, l2(x, y))
	} else {
		trap() //type
	}
	dxy(x, y)
	return r
}
func fnd(x, y K, t T) int32 {
	r := int32(0)
	xn := nn(x)
	if xn == 0 {
		return nai
	}
	xp, yp := int32(x), int32(y)
	xe := ep(x)
	s := int32(0)
	switch t - 2 {
	case 0: // ct
		r = inC(yp, xp, xe)
	case 1: // it
		s = 2
		r = inI(yp, xp, xe)
	case 2: // st
		s = 2
		r = inI(yp, xp, xe)
	case 3: // ft
		s = 3
		//r = inF(F64(yp), xp, xe)
		r = inF(yp, xp, xe)
	default: // zt
		s = 4
		//r = inZ(F64(yp), F64(yp+8), xp, xe)
		r = inZ(yp, xp, xe)
	}
	if r == 0 {
		return nai
	}
	return (r - xp) >> s
}
func fndl(x, y K) int32 {
	xn := nn(x)
	xp := int32(x)
	dx(y)
	r := int32(0)
	for r < xn {
		if match(K(I64(xp)), y) != 0 {
			return r
		}
		r++
		xp += 8
	}
	return nai
}
func idx(x, a, b int32) int32 {
	for i := a; i < b; i++ {
		if x == I8(i) {
			return i - a
		}
	}
	return -1
}

func Find(x, y K) K { // find[pattern;string] returns all matches (It)
	xt, yt := tp(x), tp(y)
	if xt != yt || xt != Ct {
		trap() //type
	}
	xn, yn := nn(x), nn(y)
	if xn == 0 || yn == 0 {
		dxy(x, y)
		return mk(It, 0)
	}
	r := mk(It, 0)
	xp, yp := int32(x), int32(y)
	y0 := yp
	e := yp + yn + 1 - xn
	for yp < e { // todo rabin-karp / knuth-morris / boyes-moore..
		if findat(xp, yp, xn) != 0 {
			r = cat1(r, Ki(yp-y0))
			yp += xn
		} else {
			yp++
		}
		continue
	}
	dxy(x, y)
	return r
}
func findat(xp, yp, n int32) int32 {
	for i := int32(0); i < n; i++ {
		if I8(xp+i) != I8(yp+i) {
			return 0
		}
	}
	return 1
}

func Mtc(x, y K) K {
	r := Ki(match(x, y))
	dxy(x, y)
	return r
}
func match(x, y K) int32 {
	yn := int32(0)
	if x == y {
		return 1
	}
	xt, yt := tp(x), tp(y)
	if xt != yt {
		return 0
	}
	if xt > 16 {
		xn := nn(x)
		yn = nn(y)
		if xn != yn {
			return 0
		}
		if xn == 0 {
			return 1
		}
		xp, yp := int32(x), int32(y)
		if xt < Dt {
			return Func[251+xt].(f3i)(xp, yp, ep(y))
		} else {
			if match(K(I64(xp)), K(I64(yp))) != 0 {
				return match(K(I64(xp+8)), K(I64(yp+8)))
			}
			return 0
		}
	}
	xp, yp := int32(x), int32(y)
	if xt < ft {
		return I32B(xp == yp)
	}
	switch int32(xt-ft) - 3*I32B(xt > 9) {
	case 0: // ft
		return I32B(0 == cmF(xp, yp)) //eqf(F64(xp), F64(yp))
	case 1: // zt
		return I32B(0 == cmZ(xp, yp)) //return eqz(F64(xp), F64(xp+8), F64(yp), F64(yp+8))
	case 2: // composition
		yn = 8 * nn(y)
	case 3: // derived
		yn = 16
	case 4: // projection
		yn = 24
	case 5: // lambda
		return match(K(I64(xp+16)), K(I64(yp+16))) // compare strings
	default: // xf
		return I32B(I64(xp) == I64(yp))
	}
	for yn > 0 { // composition, derived, projection
		yn -= 8
		if match(K(I64(xp+yn)), K(I64(yp+yn))) == 0 {
			return 0
		}
	}
	return 1
}
func mtC(xp, yp, e int32) int32 {
	ve := e &^ 7
	for yp < ve {
		if I64(xp) != I64(yp) {
			return 0
		}
		xp += 8
		yp += 8
	}
	for yp < e {
		if I8(xp) != I8(yp) {
			return 0
		}
		xp++
		yp++
	}
	return 1
}
func mtF(xp, yp, e int32) int32 {
	for yp < e {
		if cmF(xp, yp) != 0 { //eqf(F64(xp), F64(yp)) == 0 {
			return 0
		}
		xp += 8
		yp += 8
		continue
	}
	return 1
}
func mtL(xp, yp, e int32) int32 {
	for yp < e {
		if match(K(I64(xp)), K(I64(yp))) == 0 {
			return 0
		}
		xp += 8
		yp += 8
		continue
	}
	return 1
}
func In(x, y K) K {
	xt, yt := tp(x), tp(y)
	if xt == yt && xt > 16 {
		return Ecl(30, l2(x, y))
	} else if xt+16 != yt {
		trap() //type
	}
	dx(y)
	return in(x, y, xt)
}
func in(x, y K, xt T) K {
	xp, yp := int32(x), int32(y)
	e := ep(y)
	switch xt - 2 {
	case 0: //ct
		e = inC(xp, yp, e)
	case 1: //it
		e = inI(xp, yp, e)
	case 2: //st
		e = inI(xp, yp, e)
	case 3: //ft
		dx(x)
		//e = inF(F64(xp), yp, e)
		e = inF(xp, yp, e)
	default: //zt
		dx(x)
		//e = inZ(F64(xp), F64(xp+8), yp, e)
		e = inZ(xp, yp, e)
	}
	return Ki(I32B(e != 0))
}
func inC(x, yp, e int32) int32 {
	// maybe splat x to int64
	for yp < e {
		if x == I8(yp) {
			return yp //used in idxc
		}
		yp++
	}
	return 0
}
func inI(x, yp, e int32) int32 {
	for yp < e {
		if x == I32(yp) {
			return yp //used in idxi
		}
		yp += 4
	}
	return 0
}

// func inF(x float64, yp int32, e int32) int32 {
func inF(xp, yp, e int32) int32 {
	for yp < e {
		//if eqf(x, F64(yp)) != 0 {
		if cmF(xp, yp) == 0 {
			return yp
		}
		yp += 8
	}
	return 0
}

// func inZ(re, im float64, yp int32, e int32) int32 {
func inZ(xp, yp, e int32) int32 {
	for yp < e {
		//if eqz(re, im, F64(yp), F64(yp+8)) != 0 {
		if cmZ(xp, yp) == 0 {
			return yp
		}
		yp += 16
	}
	return 0
}
