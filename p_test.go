package i

import (
	"math"
	"math/cmplx"
	"reflect"
	"testing"
)

func TestP(t *testing.T) {
	testCases := []struct {
		s    s
		p, r v
	}{
		{"", nil, nil},
		{"1", z1, z1},
		{"1;2\n3", l{";", z1, z2, z3}, z3},
		{"1 2 3", zv{1, 2, 3}, zv{1, 2, 3}},
		{"1.2 3a90", zv{1.2, c(0, 3)}, zv{1.2, c(0, 3)}},
		{" -1  3a180 0i2 ", zv{-1, -3, c(0, 2)}, zv{-1, -3, c(0, 2)}},
		{"`a", l{"`", "a"}, "a"},
		{"`", l{"`", ""}, ""},
		{"`a`b", sv{"a", "b"}, sv{"a", "b"}},
		{"`a1 `bZ3\"+-,:\"", sv{"a1", "bZ3", "+-,:"}, sv{"a1", "bZ3", "+-,:"}},
		{`"x"`, l{"`", "x"}, "x"},
		{`""`, l{"`", ""}, ""},
		{`"xy"`, l{"`", "xy"}, "xy"},
		{`"x\ny"`, l{"`", "x\ny"}, "x\ny"},
		{`0x00`, l{"`", "\x00"}, "\x00"},
		{`0x0103`, sv{"\x01", "\x03"}, sv{"\x01", "\x03"}},
		{`[a:1;beta:2 3]`, l{"!", sv{"a", "beta"}, l{nil, z1, zv{2, 3}}}, [2]l{l{"a", "beta"}, l{z1, zv{2, 3}}}},
		{`*[a:1;b:2]`, l{"*", l{"!", sv{"a", "b"}, l{nil, z1, z2}}}, z1},
		{"x3:1 2i3", l{":", "x3", zv{1, c(2, 3)}}, zv{1, c(2, 3)}},
		{"+", "+", "fn"},
		{"1+2", l{"+", z1, z2}, z3},
		{"1+", l{"+", z1, nil}, "fn"},
		{"*1", l{"*", z1}, z1},
		{"a:1;a", l{";", l{":", "a", z1}, "a"}, z1},
		{"x3:2", l{":", "x3", z2}, z2},
		{"x:1;x+1", l{";", l{":", "x", z1}, l{"+", "x", z1}}, z2},
		{"1+(1;2;3)", l{"+", z1, l{nil, z1, z2, z3}}, zv{2, 3, 4}},
		{"(1+2)-3", l{"-", l{"+", z1, z2}, z3}, z0},
		{"1 2 3[0 2]", l{zv{1, 2, 3}, zv{0, 2}}, zv{1, 3}},
		{"`a`b`c[2]", l{sv{"a", "b", "c"}, z2}, "c"},
		{"(1;2;3)[0 1]", l{l{nil, z1, z2, z3}, zv{0, 1}}, zv{1, 2}},
		{"(1;;2)", l{nil, z1, nil, z2}, l{z1, nil, z2}},
		{"(1;2;)", l{nil, z1, z2, nil}, l{z1, z2, nil}},
		{"(;1;2)", l{nil, nil, z1, z2}, l{nil, z1, z2}},
		{"(1;2;())", l{nil, z1, z2, l{nil}}, l{z1, z2, l{}}},
		{"(1;(2;3);4)[1;1]", l{l{nil, z1, l{nil, z2, z3}, z4}, z1, z1}, z3},
		{"(1;(2;3;(4;0)))", l{nil, z1, l{nil, z2, z3, l{nil, z4, z0}}}, l{z1, l{z2, z3, l{z4, z0}}}},
		{"(1;(2;3;(4;0)))[1;2;0]", l{l{nil, z1, l{nil, z2, z3, l{nil, z4, z0}}}, z1, z2, z0}, z4},
		{"(1;(2;3;(4;0))).1 2 0", l{".", l{nil, z1, l{nil, z2, z3, l{nil, z4, z0}}}, zv{1, 2, 0}}, z4},
		{"[a:1;b:[c:3;d:4]]", l{"!", sv{"a", "b"}, l{nil, z1, l{"!", sv{"c", "d"}, l{nil, z3, z4}}}}, [2]l{l{"a", "b"}, l{z1, [2]l{l{"c", "d"}, l{z3, z4}}}}},
		{"[a:1;b:[c:3;d:4]]`a`b", l{l{"!", sv{"a", "b"}, l{nil, z1, l{"!", sv{"c", "d"}, l{nil, z3, z4}}}}, sv{"a", "b"}}, l{z1, [2]l{l{"c", "d"}, l{z3, z4}}}},
		{"[a:1;b:[c:3;d:4]].`b`d", l{".", l{"!", sv{"a", "b"}, l{nil, z1, l{"!", sv{"c", "d"}, l{nil, z3, z4}}}}, sv{"b", "d"}}, z4},
		{"+[3;4]", l{"+", z3, z4}, c(7, 0)},
		{`a:+;a[1;2]`, l{";", l{":", "a", "+"}, l{"a", z1, z2}}, z3},
		{"(+)[1;2]", l{"+", z1, z2}, z3},
		{"`a`b!3 4", l{"!", sv{"a", "b"}, zv{3, 4}}, [2]l{l{"a", "b"}, l{z3, z4}}},
		{"[a:(1+2);b:4]", l{"!", sv{"a", "b"}, l{nil, l{"+", z1, z2}, z4}}, [2]l{l{"a", "b"}, l{z3, z4}}},
		{"a:1;a+:1", l{";", l{":", "a", z1}, l{"+:", "a", z1}}, z2},
		{"a:1 2;a[1]:3", l{";", l{":", "a", zv{1, 2}}, l{":", l{"a", z1}, z3}}, zv{1, 3}},
		{"a:1 2;a[1]+:3", l{";", l{":", "a", zv{1, 2}}, l{"+:", l{"a", z1}, z3}}, zv{1, 5}},
		{"{2*x}`a`b!1 2", l{l{"λ", l{"*", z2, "x"}}, l{"!", sv{"a", "b"}, zv{1, 2}}}, [2]l{l{"a", "b"}, l{z2, z4}}},
		{`(3 3⍴⍳6)[0 1;1 2]`, l{l{"⍴", zv{3, 3}, l{"⍳", c(6, 0)}}, zv{0, 1}, zv{1, 2}}, l{zv{1, 2}, zv{4, 5}}}, // TODO
		{"+/1 2 3", l{l{"/", "+"}, zv{1, 2, 3}}, zi(6)},
		{"+/3", l{l{"/", "+"}, z3}, z3},
		{"+/,3", l{l{"/", "+"}, l{",", z3}}, z3},
		{"2-/3 7 8", l{l{"/", "-"}, z2, zv{3, 7, 8}}, zi(-16)},
		{`-\3 7 8`, l{l{`\`, "-"}, zv{3, 7, 8}}, zv{3, -4, -12}},
		{`2-\3 7 8`, l{l{`\`, "-"}, z2, zv{3, 7, 8}}, zv{-1, -8, -16}},
		{`1 2 3-\3 7 8`, l{l{`\`, "-"}, zv{1, 2, 3}, zv{3, 7, 8}}, l{zv{-2, -1, -0}, zv{-9, -8, -7}, zv{-17, -16, -15}}},
		{`3=\\0 0 0 0`, l{l{`\`, l{l{`\`, "="}, nil}}, z3, zv{0, 0, 0, 0}}, l{zv{0, 0, 0, 0}, zv{0, 1, 0, 1}, zv{0, 0, 1, 1}, zv{0, 1, 1, 1}}},
		{`{%x}'4 5`, l{l{"'", l{"λ", l{"%", "x"}}}, zv{4, 5}}, zv{0.25, 0.2}},
		{`15%'3 5`, l{l{"'", "%"}, c(15, 0), zv{3, 5}}, zv{5, 3}},
		{`-':3 5 2`, l{l{"':", "-"}, zv{3, 5, 2}}, zv{3, 2, -3}},
		{`=':3 3 4 4 5`, l{l{"':", "="}, zv{3, 3, 4, 4, 5}}, zv{3, 1, 0, 1, 0}},
		{`99,':2 3 4`, l{l{"':", ","}, c(99, 0), zv{2, 3, 4}}, l{zv{2, 99}, zv{3, 2}, zv{4, 3}}},
		{`2 3,/:4 5 6`, l{l{"/:", ","}, zv{2, 3}, zv{4, 5, 6}}, l{zv{2, 3, 4}, zv{2, 3, 5}, zv{2, 3, 6}}},
		{`2 3 4 ,\: 5 6 7`, l{l{`\:`, ","}, zv{2, 3, 4}, zv{5, 6, 7}}, l{zv{2, 5, 6, 7}, zv{3, 5, 6, 7}, zv{4, 5, 6, 7}}},
		{`2,:/3`, l{l{"/", ",:"}, z2, z3}, l{zv{3}}},
		{`f:(100>);f (2*)/1`, l{";", l{":", "f", l{">", c(100, 0), nil}}, l{l{"/", l{"*", z2, nil}}, "f", z1}}, c(128, 0)},
		{`⌈√:/0.5`, l{"⌈", l{l{"/", "√:"}, c(0.5, 0)}}, z1},
		{`⌈{sqr x}/0.5`, l{"⌈", l{l{"/", l{"λ", l{"sqr", "x"}}}, c(0.5, 0)}}, z1},
		{`f:{3};f[]`, l{";", l{":", "f", l{"λ", z3}}, l{"f", nilad(true)}}, z3},
		{`{x+y}[1;2]`, l{l{"λ", l{"+", "x", "y"}}, z1, z2}, z3},
		{`{y+2*x}\1 2 3`, l{l{`\`, l{"λ", l{"+", "y", l{"*", z2, "x"}}}}, zv{1, 2, 3}}, zv{1, 4, 11}},
		{`a:1;{a:2}[];a`, l{";", l{":", "a", z1}, l{l{"λ", l{":", "a", z2}}, nilad(true)}, "a"}, z1},
		{`{1;2}[]`, l{l{"λ", z1, z2}, nilad(true)}, z2},
		{`a:1;{a::2}[];a`, l{";", l{":", "a", z1}, l{l{"λ", l{"::", "a", z2}}, nilad(true)}, "a"}, z2},
		{`a:1;{a+::2}[];a`, l{";", l{":", "a", z1}, l{l{"λ", l{"+::", "a", z2}}, nilad(true)}, "a"}, z3},
		{`$[1;2;3]`, l{"$", z1, z2, z3}, z2},
		{`$[0;2;0;3;4]`, l{"$", z0, z2, z0, z3, z4}, z4},
		{`{$[x>1;x*o x-1;1]}8`, l{l{"λ", l{"$", l{">", "x", z1}, l{"*", "x", l{"o", l{"-", "x", z1}}}, z1}}, c(8, 0)}, z(40320)},
		{`{$[x>100;x;∇ x+1]}1`, l{l{"λ", l{"$", l{">", "x", c(100, 0)}, "x", l{"∇", l{"+", "x", z1}}}}, z1}, c(101, 0)},
		{`{$[x>100;x;∇[x+1]]}1`, l{l{"λ", l{"$", l{">", "x", c(100, 0)}, "x", l{"∇", l{"+", "x", z1}}}}, z1}, c(101, 0)},
		{`{-x}@2 3`, l{"@", l{"λ", l{"-", "x"}}, zv{2, 3}}, zv{-2, -3}},
		{`{x+y}.(1 2;3 4)`, l{".", l{"λ", l{"+", "x", "y"}}, l{nil, zv{1, 2}, zv{3, 4}}}, zv{4, 6}},
		{`{x<100}{2*x}\1`, l{l{`\`, l{"λ", l{"*", z2, "x"}}}, l{"λ", l{"<", "x", c(100, 0)}}, z1}, zv{1, 2, 4, 8, 16, 32, 64, 128}},
	}
	for _, tc := range testCases {
		p := P(tc.s)
		tt(t, tc.p, p, "P: %q %+v\n", tc.s, tc.p)
		r := E(p, nil)
		if tc.r == "fn" && rval(r).Kind() == reflect.Func {
		} else {
			tt(t, tc.r, r, "E: %q %+v\n", tc.s, tc.r)
		}
	}
}

func TestScan(t *testing.T) {
	type iv [10]int //0     1     2     3     4     5     6     7     8     9
	var f = [10]sf{sNum, sNam, sStr, sVrb, sAsn, sIov, sAdv, sViw, sDct, sWsp}
	var testCases = []struct {
		s s
		r iv
	}{
		{` `, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{`0`, iv{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`23`, iv{2, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`+1`, iv{2, 0, 0, 1, 0, 0, 0, 0, 0, 0}}, // +1 is a number
		{`-1`, iv{2, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
		{`1e`, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`1.`, iv{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`.5`, iv{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},  // no number: .5 use 0.5
		{`1i`, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},  // no number: 1i
		{`0i1`, iv{3, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, // complex i: 0i1
		{`-1i0`, iv{4, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
		{`i`, iv{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`-i`, iv{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
		{`1.23e+06i-1.23e-13`, iv{18, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`13.275a275.2`, iv{12, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`π`, iv{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, // name!
		{`a`, iv{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`a2`, iv{0, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`a2/`, iv{0, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"`a", iv{0, 0, 2, 0, 0, 0, 0, 0, 0, 0}},
		{"`a3", iv{0, 0, 3, 0, 0, 0, 0, 0, 0, 0}},
		{"`a3.", iv{0, 0, 3, 0, 0, 0, 0, 0, 0, 0}},
		{`"a"`, iv{0, 0, 3, 0, 0, 0, 0, 0, 0, 0}},
		{"`a_3", iv{0, 0, 2, 0, 0, 0, 0, 0, 0, 0}},
		{`"a"b`, iv{0, 0, 3, 0, 0, 0, 0, 0, 0, 0}},
		{`"x\ny"`, iv{0, 0, 6, 0, 0, 0, 0, 0, 0, 0}},
		{`"a\"b\n"b`, iv{0, 0, 8, 0, 0, 0, 0, 0, 0, 0}},
		{`+`, iv{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
		{`⍟3`, iv{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
		{`⍟:`, iv{0, 0, 0, 2, 0, 0, 0, 0, 0, 0}},
		{`+:`, iv{0, 0, 0, 2, 0, 0, 0, 0, 0, 0}},
		{`1:`, iv{1, 0, 0, 0, 0, 2, 0, 0, 0, 0}},
		{`/`, iv{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}},
		{`':`, iv{0, 0, 0, 0, 0, 0, 2, 0, 0, 0}},
		{`⍨`, iv{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}},
		{`::x`, iv{0, 0, 0, 0, 0, 0, 0, 2, 0, 0}},
		{`[`, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`[:`, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`[3:`, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{`[a:`, iv{0, 0, 0, 0, 0, 0, 0, 0, 3, 0}},
		{`[a3:`, iv{0, 0, 0, 0, 0, 0, 0, 0, 4, 0}},
		{`[ab3:`, iv{0, 0, 0, 0, 0, 0, 0, 0, 5, 0}},
		{`  \n `, iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 2}},
		{"  \t\r ", iv{0, 0, 0, 0, 0, 0, 0, 0, 0, 5}},
	}
	for _, tc := range testCases {
		for k, f := range f {
			if n := f([]rune(tc.s)); n != tc.r[k] {
				t.Fatalf("%q: f[%d] got %d, exp %d", tc.s, k, n, tc.r[k])
			}
		}
	}
}
func TestNum(t *testing.T) {
	testCases := []struct {
		s s
		z z
	}{
		{"1.23", c(1.23, 0)},
		{"π", c(math.Pi, 0)},
		{"-π", c(-math.Pi, 0)},
		{"-πi𝜀", c(-math.Pi, 1.0E-14)},
		{"1.2e-14", c(1.2e-14, 0)},
		{"∞", c(math.Inf(1), 0)},
		{"-∞", c(math.Inf(-1), 0)},
		{"0i-∞", c(0, math.Inf(-1))},
		{"ø", c(math.NaN(), 0)},
		{"1.23a0", c(1.23, 0)},
		{"1.23i3", c(1.23, 3)},
	}
	for _, tc := range testCases {
		p := p{b: rv(tc.s)}
		z := p.num(tc.s)
		if cmplx.IsNaN(z) && cmplx.IsNaN(tc.z) {
			continue
		}
		tt(t, tc.z, z, "num: %s %v %v\n", tc.s, tc.z, z)
	}
}
func TestBeg(t *testing.T) {
	testCases := [][2]string{
		{"xyz", "xyz"},
		{"/xyz", ""},
		{`/ a:[a:1;b:[d:3;f:4]]`, ""},
		{"/x\nyz", "\nyz"},
		{"ab/x\nyz", "ab/x\nyz"},
		{"ab /x\nyz", "ab \nyz"},
		{"ab /x;yz", "ab "},
		{`ab" /x;"yz`, `ab" /x;"yz`},
		{"1+2", "1+ 2"},
		{"a-2", "a- 2"},
		{"a-b", "a-b"},
		{"-13", "-13"},
		{"2.5e-03", "2.5e-03"},
		{"a+b", "a+b"},
		{"a*2.3i-5", "a*2.3i-5"},
		{"`a/`b`c", "`a/`b`c"},
		{"a:-1 -1", "a:-1 -1"},
	}
	for _, tc := range testCases {
		r := string(beg(rv(tc[0])))
		tt(t, tc[1], r, "beg %q %q\n", tc[0], tc[1])
	}
}
