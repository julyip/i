package main

import (
	. "github.com/ktye/wg/module"
)

var ps int32

func Prs(x K) K { return parse(Tok(x)) } // `p"src"  `p(token list)
func parse(x K) K {
	if tp(x) != Lt {
		trap(Type)
	}
	pp = int32(x)
	pe = pp + 8*nn(x)
	r := es()
	if pp != pe {
		trap(Parse)
	}
	lfree(x)
	return r
}
func es() K {
	r := mk(Lt, 0)
	for {
		n := next()
		if n == 0 {
			break
		}
		if n == 59 {
			continue
		}
		pp -= 8
		x := e(t()) &^ 1
		if x == 0 {
			break
		}
		if nn(r) != 0 {
			r = cat1(r, 256) // drop
		}
		r = Cat(r, x)
	}
	return r
}
func e(x K) K { // Lt
	r := K(0)
	xv := x & 1
	x &^= 1
	if x == 0 {
		return 0
	}
	xs := ps
	y := t()
	yv := y & 1
	y &^= 1
	if y == 0 {
		return x + xv
	}
	if yv != 0 && xv == 0 {
		r = e(t())
		ev := r & 1
		r &^= 1
		a := pasn(x, y, r)
		if a != 0 {
			return a
		}
		if r == 0 || ev == 1 { // 1+ (projection)
			x = cat1(ucat(cat1(cat1(ucat(l1(0), x), Ki(2)), 27), y), 92)
			if ev == 1 { // 1+-
				return cat1(ucat(r, x), 91) + 1
			}
			return x + 1
		}
		return dyadic(ucat(r, x), y) // dyadic
	}
	r = e(rx(y) + yv)
	ev := r & 1
	r &^= 1
	dx(y)
	if xv == 0 {
		return cat1(ucat(r, x), 83|K(xs)<<32) // juxtaposition
	} else if (r == y && xv+yv == 2) || ev == 1 {
		return cat1(ucat(r, x), 91) + 1 // composition
	}
	return idiom(monadic(ucat(r, x))) // monadic
}
func t() K { // Lt
	r := next()
	if r == 0 {
		return 0
	}
	if tp(r) == 0 && int32(r) < 127 {
		if is(int32(r), 32) != 0 {
			pp -= 8
			return 0
		}
	}
	verb := K(0)
	if r == K('(') {
		r = plist(41)
		r = rlist(r&^1, 0)
	} else if r == K('{') {
		r = plam(ps)
	} else if r == K('[') {
		r = es()
		if next() != K(']') {
			trap(Parse)
		}
		return r
	} else if tp(r) == st {
		r = l2(r, 20|(K(ps)<<32)) // .`x (lookup)
	} else {
		rt := tp(r)
		if rt == 0 {
			r, verb = quote(r)|K(ps)<<32, 1
		} else if rt == St {
			if nn(r) == 1 {
				r = Fst(r)
			}
		}
		r = l1(r)
	}
f:
	for {
		n := next()
		if n == 0 {
			break f
		}
		a := int32(n)
		tn := tp(n)
		ks := K(ps) << 32
		if tn == 0 && a > 20 && a < 27 { // +/
			r, verb = cat1(r, n), 1
		} else if n == 91 { // [
			verb = 0
			n = plist(93)
			p := K(84) + 8*(n&1) // 92(project) or call(84)
			n &^= 1
			s := pspec(r, n)
			if s != 0 {
				return s
			}
			if nn(n) == 1 {
				r = cat1(ucat(Fst(n), r), 83|ks)
			} else {
				n = rlist(n, 2)
				r = cat1(Cat(n, r), p|ks)
			}
		} else {
			pp -= 8
			break f // within else-if
		}
	}
	return r + verb
}
func pasn(x, y, r K) K {
	l := K(I64(int32(y)))
	v := int32(l)
	sp := 0xffffff & int32(l>>32)
	if nn(y) == 1 && tp(l) == 0 && v == 449 || (v > 544 && v < 565) {
		dx(y)
		xn := nn(x)
		if xn > 2 { // indexed amd/dmd
			if v > 544 { // indexed-modified
				l -= 96
			}
			s := ati(rx(x), xn-3)
			lp := 0xff000000ffffffff & lastp(x)
			// (+;.i.;`x;.;@) -> x:@[x;.i.;+;rhs] which is (+;.i.;`x;.;211 or 212)
			// lp+128 is @[amd..] or .[dmd..]
			if lp == 92 {
				lp = 84 // x[i;]:.. no projection
			}
			x = cat1(cat1(ucat(l1(l), ldrop(-2, x)), 20), (K(sp)<<32)|(lp+128))
			y = l2(s, 448) // s:..
		} else if v == 449 || v == 545 {
			s := Fst(x) // (`x;.)
			if loc != 0 && v == 449 {
				loc = Cat(loc, rx(s))
			}
			x = l1(s)
			y = l1(448) // asn
		} else { // modified
			y = cat1(l2(unquote(l-32), Fst(rx(x))), 448)
		}
		return dyadic(ucat(r, x), y)
	}
	return 0
}
func plam(s0 int32) K {
	r := K(0)
	slo := loc
	loc = 0
	ar := int32(-1)
	n := next()
	if n == 91 { // argnames
		n := plist(93) &^ 1
		ln := nn(n)
		loc = Ech(4, l1(n)) // [a]->,(`a;.)  [a;b]->((`a;.);(`b;.))
		if ln > 0 && tp(loc) != St {
			trap(Parse)
		}
		ar = nn(loc)
		if ar == 0 {
			dx(loc)
			loc = mk(St, 0)
		}
	} else {
		pp -= 8
		loc = mk(St, 0)
	}
	c := es()
	n = next()
	if n != 125 {
		trap(Parse)
	}
	cn := nn(c)
	cp := int32(c)
	if ar < 0 {
		ar = 0
		for i := int32(0); i < cn; i++ {
			r = K(I64(cp))
			if tp(r) == 0 && int32(r) == 20 {
				r = K(I64(cp - 8))
				y := int32(r) >> 3
				if tp(r) == st && y > 0 && y < 4 {
					ar = maxi(ar, y)
				}
			}
			cp += 8
		}
		loc = Cat(ntake(ar, rx(xyz)), loc)
	}
	i := Add(seq(1+ps-s0), Ki(s0-1))
	s := atv(rx(src()), i)
	r = l3(c, Unq(loc), s)
	loc = slo
	return l1(slam(r, ar, s0))
}
func slam(r K, ar, s0 int32) K {
	rp := int32(r)
	SetI32(rp-12, ar)
	return K(rp) | (K(s0) << 32) | K(lf)<<59
}
func pspec(r, n K) K {
	ln := nn(n)
	v := K(I64(int32(r)))
	if nn(r) == 1 && ln > 2 { // $[..] cond
		if tp(v) == 0 && int32(v) == 465 {
			dx(r)

			//n = Out(n)
			return l2(n, 22) // Cnd
			//return pcnd(n, ln)
		}
	}
	if nn(r) == 2 && ln > 1 && int32(v) == 64 { // while[..]
		dx(r)
		r = Fst(rx(n))
		return cat1(cat1(Enl(join(256, ndrop(1, n))), r), 86)
		//return pwhl(n, ln-1)
	}
	return 0
}

/*
func pwhl(x K, xn int32) K {
	c := Fst(rx(x))
	return cat1(cat1(Enl(join(256, ndrop(1, x))), c), 86)

		r := cat1(Fst(rx(x)), 0)
		p := nn(r) - 1
		r = cat1(r, 384) // jif
		r = cat1(r, 256) // drop
		xp := int32(x)
		sum := int32(2)
		for i := int32(0); i < xn; i++ {
			if i != 0 {
				r = cat1(r, 256)
			}
			xp += 8
			y := x0(K(xp))
			sum += 1 + nn(y)
			r = ucat(r, y)
		}
		r = cat1(cat1(r, Ki(-8*(2+nn(r)))), 320) // jmp back
		SetI64(int32(r)+8*p, int64(Ki(8*sum)))   // jif
		dx(x)
		return ucat(l1(0), r) // null for empty while
}
func pcnd(x K, xn int32) K {
	nxt := int32(0)
	sum := int32(0)
	xp := int32(x) + 8*xn
	state := int32(1)
	for xp != int32(x) {
		xp -= 8
		r := K(I64(xp))
		if sum > 0 {
			state = 1 - state
			if state != 0 {
				r = cat1(cat1(r, Ki(nxt)), 384) // jif
			} else {
				r = cat1(cat1(r, Ki(sum)), 320) // j
			}
			SetI64(xp, int64(r))
		}
		nxt = 8 * nn(r)
		sum += nxt
	}
	return flat(x)
}
*/
func plist(c K) K {
	p := K(0)
	r := mk(Lt, 0)
	n := int32(0)
	for {
		b := next()
		if b == 0 || b == c {
			break
		}
		if n == 0 {
			pp -= 8
		}
		if n != 0 && b != 59 {
			trap(Parse)
		}
		n++
		x := e(t()) &^ 1
		if x == 0 {
			p = 1
		}
		r = cat1(r, x)
	}
	return r + p
}
func rlist(x, p K) K {
	n := nn(x)
	if n == 0 {
		return l1(x)
	}
	if n == 1 {
		return Fst(x)
	}
	if p != 2 {
		p = clist(x, n)
		if p != 0 {
			return l1(p)
		}
	}
	return cat1(cat1(flat(Rev(x)), Ki(n)), 27)
}
func clist(x K, n int32) K {
	p := int32(x)
	for i := int32(0); i < n; i++ {
		xi := K(I64(p))
		t := tp(xi)
		if t != Lt {
			return 0
		}
		if nn(xi) != 1 {
			return 0
		}
		if tp(K(I64(int32(xi)))) == 0 {
			return 0
		}
		p += 8
	}
	return uf(flat(x))
}

func next() K {
	if pp == pe {
		return 0
	}
	r := K(I64(pp))
	ps = 0xffffff & int32(r>>32)
	r = r &^ (K(0xffffff) << 32)
	pp += 8
	return r
}
func lastp(x K) K { return K(I64(int32(x) + 8*(nn(x)-1))) }
func dyadic(x, y K) K {
	l := lastp(y)
	if quoted(l) != 0 {
		return cat1(ucat(x, ldrop(-1, y)), 64+unquote(l))
	}
	return cat1(ucat(x, y), 128)
}
func monadic(x K) K {
	l := lastp(x)
	if quoted(l) != 0 {
		r := cat1(ldrop(-1, x), unquote(l))
		if int32(l) == 449 { // :x (return: identity+jump)
			return cat1(cat1(r, Ki(1048576)), 320)
		} else {
			return r
		}
	}
	return cat1(x, 83) // dyadic-@
}
func ldrop(n int32, x K) K { return explode(ndrop(n, x)) }
func svrb(p int32) int32 {
	x := K(I64(p))
	return I32B(int32(x) < 64 && tp(x) == 0) * int32(x)
}
func idiom(x K) K {
	l := int32(x) + 8*(nn(x)-2)
	i := svrb(l) + svrb(l+8)<<6
	if i == 262 || i == 263 { // *& 6 4 -> 40
		i = 34 // 6->40(Fwh) 7->41(Las)
	} else if i == 1166 { // ?^ 14 18
		i = 23 // 14->37(Uqs)
	} else {
		return x
	}
	SetI64(l, I64(l)+int64(i))
	return ndrop(-1, x)
}
