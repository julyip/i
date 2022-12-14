package main

import (
	. "github.com/ktye/wg/module"
)

var tot, top int32

func init() {
	Memory(1)
	// no @ $ /: \: ': ` [ ]  { }  may be assigned to as user functions but not infix

	// x(space)y          is (,x),y
	// e.g.:   l:4 5 6    is (4;5;6)
	//         1 l 2      is (1;4 5 6;2)  or as:  1 (4 5 6) 2      (space matters)
	//
	// "abc" is quotation/literal, also used for lambdas: f:"(+/x)%#x"
	//
	// juxtaposition
	// ay cal
	// vy idx
	Data(0, "\x2b\x2d\x2a\x25\x26\x7c\x3c\x3e\x3d\x21\x3a\x7e\x2c\x5e\x23\x5f\x3f\x2e\x20\x27\x2f\x5c") // +-*%&|<>=!:~,^#_?. '/\  (22)

	//              :    ~    ,    ^    #    _    ?    .         '    /    \
	Functions(00, asn, mtc, cat, cut, tak, drp, rol, cal, spc, ech, ovr, scn)                                    //ay
	Functions(12, asn, mtc, cat, ctv, tkv, dpv, fnd, atx, spc, inn, spl, jon)                                    //vy
	Functions(24, add, sub, mul, div, min, max, les, mor, eql, mod)                                              //scalar dyadic
	Functions(34, flp, neg, fst, rot, wer, rev, grd, gdn, grp, til, idn, not, enl, str, cnt, lst, unq, val, enl) //monadic
	//              +    -    *    %    &    |    <    >    =    !    :    ~    ,    ^    #    _    ?    .  spc    '   /  \
	//             43   45   42   37   38  124   60   62   61   33   58  126   44   94   35   95   63   46   32   39  47  92
}
func main() {
	k1()
}
func k1() {
	tot = 65536
	top = 24
	rm(256) //keys at 24
	rm(256) //vals at 270
}

func v(x int32) int32 { return x >> 1 }
func w(x int32) int32 { return x << 1 }
func n(x int32) int32 {
	if x&1 != 0 {
		return -1
	}
	return I32(v(x) - 4)
}
func mk(x int32) int32 {
	x = x*4 + 4
	for tot < top+x {
		Memorygrow(1)
		tot += 65536
	}
	r := top
	SetI32(r, x)
	top += x
	return 1 + w(r)
}
func rm(x int32) int32 { // reset make, use with c1
	r := mk(x)
	SetI32(v(r)-1, 0)
	return r
}
func c1(x, y int32) int32 {
	p := I32(v(x))
	SetI32(p, 1+p)
	SetI32(v(x)+4*p, y)
	return x
}
func l2(x, y int32) int32 { return cat(enl(x), enl(y)) }
func el(x int32) int32 {
	if n(x)&1 != 0 {
		return x
	}
	return enl(x)
}
func ec2(f, x, y int32) int32 {
	rn := max(cnt(x), cnt(y)) >> 1
	r := mk(rn)
	p := v(r)
	for i := int32(0); i < rn; i++ {
		SetI32(p+4*i, cal(f, cat(enl(atx(x, i)), enl(atx(y, i)))))
	}
	return r
}
func seq(x, o, m int32) int32 {
	r := rm(x)
	for i := int32(0); i < x; i++ {
		c1(r, (i+o)%m)
	}
	return r
}
func lup(x int32) int32 {
	return x //nyi
}

type f1 = func(int32) int32
type f2 = func(int32, int32) int32

//dyadic
func add(x, y int32) int32 { return x + y }
func sub(x, y int32) int32 { return x - y }
func mul(x, y int32) int32 { return x * y }
func div(x, y int32) int32 { return x / y }
func mod(x, y int32) int32 { return x % y }
func min(x, y int32) int32 {
	if x < y {
		return x
	} else {
		return y
	}
}
func max(x, y int32) int32 {
	if x > y {
		return x
	} else {
		return y
	}
}
func les(x, y int32) int32 { return I32B(x < y) }
func mor(x, y int32) int32 { return I32B(x > y) }
func eql(x, y int32) int32 { return I32B(x == y) }

func ny2(x, y int32) int32 { return y }

func asn(x, y int32) int32 { return ny2(x, y) }
func mtc(x, y int32) int32 { return w(I32B(x == y)) }
func nul(x, y int32) int32 { return 0 }
func mtv(x, y int32) int32 {
	xn := n(x)
	if x == y {
		return 2
	}
	if xn != n(y) {
		return 0
	}
	return ovr(38, ec2(126, x, y)) // &/x~'y
}
func cat(x, y int32) int32 { // x,y
	x, y = el(x), el(y)
	xn, yn := n(x), n(y)
	r := mk(xn + yn)
	Memorycopy(v(r), v(x), 4*xn)
	Memorycopy(v(r)+4*xn, v(y), 4*yn)
	return r
}
func cut(x, y int32) int32 { return ny2(x, el(y)) }                           // a^y
func ctv(x, y int32) int32 { return ny2(x, el(y)) }                           // v^y
func tak(x, y int32) int32 { return atx(el(y), cal(66, l2(til(x), cnt(x)))) } // a#v
func tkv(x, y int32) int32 { return atx(y, wer(inn(el(y), x))) }              // v#y
func drp(x, y int32) int32 { // a_y
	yn := n(y)
	xv := v(x)
	if x < 0 {
		if xv < -yn {
			return atx(x, seq(-xv, xv+yn, yn))
		}
		return tak(x, w(-xv))
	} else {
		if xv > y {
			return mk(0)
		}
		return atx(x, seq(yn-xv, xv, yn))
	}
}
func dpv(x, y int32) int32 { return atx(y, wer(not(inn(el(y), x)))) } // v_y
func rol(x, y int32) int32 { return ny2(x, y) }                       // a?y
func fnd(x, y int32) int32 { return fst(ec2(126, x, y)) }             // v?a
func fnx(x, y int32) int32 { return ec2(63, enl(x), y) }              // v?v   x?/:y
func cal(x, y int32) int32 { // a.a  a.v
	y = el(y)
	yn := n(y)
	i := int32(0)
	for ; i < 21; i++ {
		if x == I8(i) {
			break
		}
	}
	if i > 21 {
		return exe(lup(x), y)
	}
	x = fst(y)
	y = I32(4 + v(y))
	xa := I32B(n(x) < 0)
	ya := I32B(n(y) < 0)
	mo := I32B(yn < 2)
	if i < 10 { //scalar
		i += 48
		if mo != 0 {
			return nd(i, x, y, xa+ya)
		}
		return nm(i+10, x, xa)
	}
	i += 24*ya + 12*xa
	if mo != 0 {
		return Func[i+48].(f1)(x)
	}
	return Func[i].(f2)(x, y)
}
func atx(x, y int32) int32 { // v.a  (also a.v)
	nn := n(y)
	if nn < 0 {
		return ec2(46, enl(x), y)
	}
	nn = n(x)
	if nn < 0 {
		return x
	}
	y = v(y)
	if y < 0 || nn > y {
		return 0
	} else {
		return I32(v(x) + 4*y)
	}
}
func spc(x, y int32) int32 { return cat(enl(x), y) } // x(space)y
func ech(x, y int32) int32 { // a'a  a'v
	yn := v(cnt(y))
	r := mk(yn)
	p := v(r)
	for i := int32(0); i < yn; i++ {
		SetI32(p, cal(x, atx(y, w(i))))
		p += 4
	}
	return r
}
func inn(x, y int32) int32 {
	yn := w(n(y))
	if n(x) < 0 {
		return mor(yn, fnx(y, x))
	} else {
		return ovr(124, mor(yn, fnd(y, x))) // |/(#y)>y?x
	}
}
func ovr(x, y int32) int32 { // a/a  a/v
	yn := v(cnt(y))
	r := fst(x)
	for i := int32(1); i < yn; i++ {
		r = cal(x, atx(y, w(i)))
	}
	return r
}
func spl(x, y int32) int32 { return cut(cat(0, wer(ec2(126, enl(x), fst(y)))), x) } // v/a    (0,&(,x)~'*y)^x
func scn(x, y int32) int32 { // a\a  a\v
	yn := v(cnt(y))
	r := mk(yn)
	p := v(r)
	f := fst(x)
	SetI32(p, f)
	for i := int32(0); i < yn; i++ {
		p += 4
		f = cal(x, cat(enl(f), enl(atx(y, w(i)))))
		SetI32(p, f)
	}
	return r
}
func jon(x, y int32) int32 { return ovr(44, ec2(44, x, enl(y))) } // v\a   ,/x,',y

func nm(f, x, a int32) int32 {
	if a != 0 {
		return ech(46, cat(f, enl(x)))
	}
	return w(Func[f].(f1)(v(x)))
}
func nd(f, x, y, a int32) int32 {
	if a != 0 {
		return ec2(I8(f), x, y)
	}
	return w(Func[f].(f2)(v(x), v(y)))
}

func ecv(x, y int32) int32 { return ec2(39, x, enl(y)) } // v'a  v'a

func ovv(x, y int32) int32 { return ec2(47, x, enl(y)) } // v/a  v/v

//monadic
func nyi(x int32) int32 { return x }
func flp(x int32) int32 { return nyi(x) } // +x
func neg(x int32) int32 { // -x
	if n(x) < 0 {
		x = v(x)
		if x < 0 {
			x = -x
		}
		return w(x)
	}
	return ech(90, x)
}
func fst(x int32) int32 { return atx(x, 0) }              // *x
func rot(x int32) int32 { return cat(drp(2, x), fst(x)) } // %x
func wer(x int32) int32 { // &x
	x = el(x)
	r := mk(0)
	xn := n(x)
	for i := int32(0); i < xn; i++ {
		j := w(i)
		r = cat(r, tak(j, atx(x, j)))
	}
	return r
}
func rev(x int32) int32 { // |x
	x = el(x)
	xn := n(x)
	r := rm(xn)
	for i := int32(0); i < xn; i++ {
		c1(r, xn-i-1)
	}
	return atx(x, r)
}
func grd(x int32) int32 { return nyi(el(x)) }               // <x
func gdn(x int32) int32 { return nyi(el(x)) }               // >x
func grp(x int32) int32 { return nyi(el(x)) }               // =x
func til(x int32) int32 { return seq(v(x), 0, 2147483647) } // !0
func idn(x int32) int32 { return x }                        // :x
func not(x int32) int32 { // ~a
	if n(x) < 0 {
		return w(I32B(v(x) != 0))
	}
	return ech(252, x)
}
func enl(x int32) int32 { r := mk(1); SetI32(r, x); return r } // ,x
func str(x int32) int32 { // ^x
	if n(x) < 0 {
		x = v(x)
		if x < 0 {
			return cat(90, str(w(-x)))
		} else if x == 0 {
			return enl(0)
		}
		r := rm(10)
		i := int32(0)
		for x > 0 {
			c1(r, x%10)
			x -= 10
			x /= 10
			i++
		}
		return tak(w(i), r)
	}
	return ech(188, x)
}
func cnt(x int32) int32 { // #x
	xn := n(x)
	if xn < 0 {
		return 2
	} else {
		return w(xn)
	}
}
func lst(x int32) int32 { return atx(el(x), w(n(x)-1)) } // _x
func unq(x int32) int32 { return nyi(x) }                // ?x
func val(x int32) int32 { // .x
	xn := n(x)
	if xn < 0 {
		return lup(x)
	}
	return exe(x, 0)
}
func tok(r, x, t int32) int32 {
	xn := n(x)
	r := rm(xn)
	x = v(x)
	xe := x + 4*xn
	t := 0
	q := rm(xn)
	for x < xe {
		c := I32(x)
		if 0 < n(q) {
			if c == 68 { // "
				c1(r, q)
				q = rm(xn)
				continue
			}
			c1(q, c)
			continue
		}
		if c >= 96 && c <= 114 {
			t *= 10
			t += v(c - 36)
			continue
		}
		c1(r, x)
		x += 4
	}
	return r
}
func prs(r, x, t int32) int32 {
	// 11-(1+1)+-1
	//          -   1)     >   A       ? /    >d
	//          +  -1)     m   B       - -    m>
	//          )  +1)     >   C       1 -    >d
	//          1  )+1)    >   D       1 1    c
	//                                 ( ?    <>
	//          +  1)+1)   >   A     else     >
	//          1  +1)+1)  d   E
	//          (  1)+1)   b<  F      -/1()
	//          1    +1)   d   E
	//          -    1)    >   A
	//          1    -1)   d   E
	//          1    1)    j   G
	//               1)
	//          +    +)    m
	//
	// x == (      => move<y over (, remove )  next
	// y == /      => move> dyadic
	// y == -
	//   | x == 1  => move> dyadic
	//   | x == -  => monadic, move>
	// else        => move
}
func exe(x, a int32) int32 {
	x = rev(run(tok(x)), 0)
}
