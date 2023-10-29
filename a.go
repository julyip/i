package main

import (
	. "github.com/ktye/wg/module"
)

func init() {
	Memory(1)
	Memory2(1)
	Data(132, "\x00\x01@\x01\x01\x01\x01\t\x10`\x01\x01\x01\x01\x01\t\xc4\xc4\xc4\xc4\xc4\xc4\xc4\xc4\xc4\xc4\x01 \x01\x01\x01\x01\x01BBBBBBBBBBBBBBBBBBBBBBBBBB\x10\t`\x01\x01\x00\xc2\xc2\xc2\xc2\xc2\xc2BBBBBBBBBBBBBBBBBBBB\x10\x01`\x01") // k_test.go: TestClass
	Data(227, ":+-*%&|<>=~!,^#_$?@.':/:\\:vbcisfzldtmdplx00BCISFZLDT0")
	Export(main, Asn, Atx, Cal, cs, dx, Kc, Kf, Ki, kinit, l2, mk, nn, repl, rx, sc, src, Srcp, tp, trap, Val)

	//            0    :    +    -    *    %    &    |    <    >    =10   ~    !    ,    ^    #    _    $    ?    @    .20  '    ':   /    /:   \    \:                  30                       35                       40                       45
	Functions(00, nul, Idy, Flp, Neg, Fst, Sqr, Wer, Rev, Asc, Dsc, Grp, Not, Til, Enl, Srt, Cnt, Flr, Str, Unq, Typ, Val, ech, nyi, rdc, nyi, scn, nyi, lst, Kst, Out, nyi, nyi, Abs, Img, Cnj, Ang, nyi, Uqs, nyi, Tok, Fwh, Las, Exp, Log, Sin, Cos, Prs)
	Functions(64, Asn, Dex, Add, Sub, Mul, Div, Min, Max, Les, Mor, Eql, Mtc, Key, Cat, Cut, Tak, Drp, Cst, Fnd, Atx, Cal, Ech, nyi, Rdc, nyi, Scn, nyi, com, prj, Otu, In, Find, Hyp, Cpx, nyi, Rot, Enc, Dec, nyi, nyi, Bin, Mod, Pow, Lgn, nyi, nyi, Rtp)
	Functions(193, tchr, tnms, tvrb, tpct, tvar, tsym)
	Functions(211, Amd, Dmd)

	Functions(220, negi, negf, negz) //220
	Functions(223, absi, absf, nyi)  //227
	Functions(226, addi, addf, addz) //234
	Functions(229, subi, subf, subz) //245
	Functions(232, muli, mulf, mulz) //256
	Functions(235, divi, divf, divz) //267
	Functions(238, mini, minf, minz) //278
	Functions(241, maxi, maxf, maxz) //289
	Functions(244, modi, sqrf, nyi)  //300

	Functions(247, cmi, cmi, cmi, cmF, cmZ, cmC, cmI, cmI, cmF, cmZ)

	Functions(271, guC, guC, guI, guI, guF, guZ, guL, gdC, gdC, gdI, gdI, gdF, gdZ, gdL) //353

	Functions(285, sum, rd0, prd, rd0, min, max)
	Functions(291, sums, rd0, prds, rd0, rd0)
	Functions(296, mtC, mtC, mtC, mtF, mtF, mtL)
	Functions(302, exp1, log1, sin1, cos1, pow2)

}

func trap() {
	s := src()
	if srcp == 0 {
		write(Ku(2608)) // 0\n
	} else {
		a := maxi(srcp-30, 0)
		for i := a; i < srcp; i++ {
			if I8(int32(s)+i) == 10 {
				a = 1 + i
			}
		}
		b := mini(nn(s), srcp+30)
		for i := srcp; i < b; i++ {
			if I8(int32(s)+i) == 10 {
				b = i
				break
			}
		}
		Write(0, 0, int32(s)+a, b-a)
		if srcp > a {
			write(Cat(Kc(10), ntake(srcp-a-1, Kc(32))))
		}
	}
	write(Ku(2654)) // ^\n
	panic(srcp)
}
func Srcp() int32 { return srcp }
