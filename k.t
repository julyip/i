1 /1
1.1 -0w /1.1 -0w
1. 0n 0w -1.1 /1. 0n 0w -1.1
1. 0n 0w -1.1 -0w -0n 2p .5p /1. 0n 0w -1.1 -0w 0n 6.283185 1.570796
1. /1.
.1 /0.1
1.1 /1.1
1a /1a
1a 2a20 0a 1.0a 1.a 0na 0wa /1a 2a20 0a 1a 1a 0na 0wa
1a90 /1a90
1.+2. /3.
1 2+3 /4 5
1+2 3 /3 4
1 2+3 4 /4 6
(1;2;3) /1 2 3
(`a`b!1 2;`a`b!3 4) /+`a`b!(1 3;2 4)
(`a`b!1 2;`a`c!3 4) /(`a`b!1 2;`a`c!3 4)
-(1 2 3) /-1 -2 -3
-(1;2 3) /(-1;-2 -3)
0x /""
0x6162 /"ab"
0x6162,"ab" /"abab"
@0x61 /`c
0x0a /"\n"
x:1;0x6162x /"b"
#"abc\n\r\t\\\"9" /9
"abc\n\r\t\\\"9" /"abc\n\r\t\\\"9"
0+"\t\n\r\"\\" /9 10 13 34 92
- 0101b /0 -1 0 -1
-"aBcAZ" /"abcaz"
abs"azAZb" /"AZAZB"
1a /1a
."1+2" /3
.() /
`2 /+
`1+(+) /-
`t"1+2" /(1;+;2)
0+`t"1+ \\2" /1 2 29 2
`p"1+2" /(2;1;+)
`p `t"1+2" /(2;1;+)
+[2;] /2+
+[;2] /+[;2]
f:{x+y};f 3 /{x+y}[3;]
({x+y}3)4 /7
(({x+y+z}3)4)5 /12
({x+y+z}3)4 /{x+y+z}[3;4;]
({x+y+z}[;1;])[;2] /{x+y+z}[;1;2]
.(1;2;`66) /3
1 /1
1b /1b
101b /101b
`x /`x
``x /``x
`"a" /`a
"" /""
"3" /"3"
1+2 /3
1-2 /-1
1 2 3+4 /5 6 7
1 2 3+4a /5a 6a 7a
1-2 3 /-1 -2
1 2+(3;4 5) /(4;6 7)
3+`a`b!1 2 /`a`b!4 5
(`a`b!7 8)+`a`b!1 2 /`a`b!8 10
,1 /,1
,`a`b!1 2 /+`a`b!(,1;,2)
1 2,3.5,1. 2.,1a,1a 2a45,"cd",1b /(1;2;3.5;1.;2.;1a;1a;2a45;"c";"d";1b)
(`a`b!1 2),`c`d!3 4 /`a`b`c`d!1 2 3 4
(+`a`b!(1 2;3 4)),`a`b!8 9 /+`a`b!(1 2 8;3 4 9)
t,t:+`a`b!(1 2;3 4) /+`a`b!(1 2 1 2;3 4 3 4)
(+`a`b!(1 2;3 4)),'+`c`d!(`e`f;3 4.) /+`a`b`c`d!(1 2;3 4;`e`f;3. 4.)
"aBc"<"abc" /010b
~1b /0b
~1 0 2 /010b
~```a`b /1100b
~" ab" /000b
1a=1 2 /10b
- 1b /-1
1-1b /0
-|1+2 /-3
-`a`b!1 2 /`a`b!-1 -2
0. 1.!0 /0. 1.!0 0
1 2 15 17%8 /0 0 1 2
10 20%2. /5. 10.
((4;8 8);1 2 3 4)%(2 3;2) /((2;2 2);0 1 1 2)
2\2 4 6 /1 2 3
2/1 2 3 /1 0 1
101b&110b /100b
101b|110b /111b
#`a`b`c!!3 /3
#+`a`b!(1 2 3;4 5 6) /3
3#!5 /0 1 2
-3#!5 /2 3 4
6#!3 /0 1 2 0N 0N 0N
-6#!3 /0N 0N 0N 0 1 2
"ab"#"abc" /"ab"
2#`a`b`c!1 2 3 /`a`b!1 2
2#+`a`b!(1 2 3;4 5 6) /+`a`b!(1 2;4 5)
2#'`a`b`c!1 2 3 /`a`b`c!(1 1;2 2;3 3)
2#`a!+`a`b!(1 2 3 4;5 6 7 8) /(+(,`a)!,1 2)!+(,`b)!,5 6
`a#+`a`b!(1 2;3 4) /+(,`a)!,1 2
`b`a#+`a`b!(1 2;3 4) /+`b`a!(3 4;1 2)
(3>)#0 1 2 3 4 /0 1 2
3_!5 /3 4
-3_!5 /0 1
6_!3 /!0
-6_!3 /!0
0+_!5 /0 1 2 3 4
1_("ab";"cd") /,"cd"
"ab"_"abc" /,"c"
_`x`y`z /8 16 24
`b_`a`b`c!(1 2;3 4;5 6) /`a`c!(1 2;5 6)
`c`b_`a`b`c!(1 2;3 4;5 6) /(,`a)!,1 2
`b_+`a`b`c!(1 2;3 4;5 6) /+`a`c!(1 2;5 6)
2_+`a`b`c!(1 2 3;4 5 6;7 8 9) /+`a`b`c!(,3;,6;,9)
(3>)_0 1 2 3 4 /3 4
2 5^"alphabeta" /("pha";"beta")
3^!8 /(0 1;2 3;4 5)
"ABC"^"abcCdeAgh" /("abc";"de";"gh")
`a!1 /(,`a)!,1
+(1;2 3;4 5 6.) /((1;2;4.);(1;3;5.);(1;0N;6.))
+("abc";"def") /("ad";"be";"cf")
+`a`b!(1 2;3 4) /+`a`b!(1 2;3 4)
{+/x*y}\:[(1 2 3;4 5 6);7 8 9] /50 122
"b"\:"abc" /(,"a";,"c")
"x"\:"abxdexfg" /("ab";"de";"fg")
"xd"\:"abxdexfg" /("ab";"exfg")
"xy"\:"abcdx" /,"abcdx"
"x"/:("ab";,"c";"ef") /"abxcxef"
"xy"/:("ab";,"c";"ef") /"abxycxyef"
#/:5 3 1 /1
#\:5 3 1 /(5 3 1;3;1)
1 2+/:2 /3 4
1+\:2 3 /3 4
{4/1+x}/:1 /0
{4/1+x}\:1 /1 2 3 0
_(97;-2.3 2.3;2a90) /("a";-3 2;0.)
=`b`a`a`c`b /`b`a`c!(0 4;1 2;,3)
?4 3 3 2 4 /4 3 2
?"alpha" /"alph"
1 2 2 1?2 /1
0001001b?1b /3
1 2 2 1?1 2 3 /0 1 0N
1110010b?00101b /3 3 0 3 0
1111b?0000b /0N 0N 0N 0N
3 4 -5 8 8?8 -4 9 4 /3 0N 0N 1
3 4 5?!10 /0N 0N 0N 0 1 2 0N 0N 0N 0N
3 4 -5 8 8?8 -4 9 4 1 1 1 1 1 /3 0N 0N 1 0N 0N 0N 0N 0N
"abc"?"al" /0 0N
"abc"?"alphabeta" /0 0N 0N 0N 0 1 0N 0N 0
"abc"?"xxxbc" /0N 0N 0N 1 2
"abcdefghijklmnopqrstuvxwyz"?"f" /5
"abcdefghijklmnopqrstuvxwyz"?"z" /25
(!20)?3 /3
(!20)?17 /17
1 2 1024?!10 /0N 0 1 0N 0N 0N 0N 0N 0N 0N
(!0)?1 /0N
1 2 3.?2 4. /1 0N
1a20 2a30?2a30 4a50 /1 0N
`z`x`y?`a`b`c`x`y`z`a`b`c`x`y`z /0N 0N 0N 1 2 0 0N 0N 0N 1 2 0
("abc";"de")?"de" /1
("abc";"de")?("de";"gh") /1 0N
((1;"a");(2;"b"))?,(2;"b") /,1
"abc"?(("bc";"a");"cc") /((1 2;0);2 2)
(`a`b!1 2)?2 /`b
3 in 0 1 2 /0b
3 in !5 /1b
"a"in"abc" /1b
"ad"in"abc" /10b
10b in\: 111b /10b
3 4.in 0.+!4 /10b
1a in 2 1 3a /1b
in 000b /0b
in 010b /1b
"ab" find "aaabcabca" /2 5
"aa" find "aaaaaaa" /0 2 4
"ab" find "xab" /,1
"ab" find "" /!0
"" find "ab" /!0
abs(1;2a30) /(1;2.)
(1%%2)~imag 1a45 /1b
imag -conj 1 2 imag(3 4.;5 6) /(3. 4.;5. 6.)
(3;4 5)angle 80 /(3a80;4a80 5a80)
angle(3;4 5)angle 80 /(80.;80. 80.)
exp -0.5 0.5 1 2 3 /0.60653 1.648721 2.718281 7.389056 20.085536
sin -3 -2 -1 0 1 2 3 /-0.14112 -0.909297 -0.84147 0. 0.84147 0.909297 0.14112
cos -3 -2 -1 0 1 2 3 /-0.989992 -0.416146 0.540302 1. 0.540302 -0.416146 -0.989992
2 exp -2 -1.5 -1 0 1 1.5 2 /0.25 0.353553 0.5 1. 2. 2.828427 4.
log -1 0 0.1 0.5 1 2 /0n -0w -2.302585 -0.693147 0. 0.693147
2 log -2 2 3 4 5 /0n 1. 1.584962 2. 2.321928
10 log 10 11 10000 /1. 1.041392 4.
3.14 log 10 11 10000 /2.012357 2.095654 8.049429
fill(0N;"a b";0.0 0n 2.0 3.0) /(1b;010b;0100b)
0 fill(0N;1 0N 2;3) /(0;1 0 2;3)
"x"fill"a b" /"axb"
"123"fill"a b" /"a2b"
-1 fill 0 0N 1 /0 -1 1
3=3 /1b
$1b /"1b"
3 4 5=4 /010b
1 2 3 /1 2 3
+/1 2 3 /6
3+/1 2 3 /9
+/0.+!1000 /499500.
-/5 /5
2-/5 /-3
2-/!0 /2
+/`a`b!(1 2 3;4 5 6) /5 7 9
+/'`a`b!(1 2 3;4 5 6) /`a`b!6 15
10+/+`a`b!(1 2 3;4 5 6) /`a`b!16 25
+/+`a`b!(1 2 3;4 5 6) /`a`b!6 15
+/'+`a`b!(1 2 3;4 5 6) /5 7 9
+\`a`b!(1 2 3;4 5 6) /`a`b!(1 2 3;5 7 9)
+\'`a`b!(1 2 3;4 5 6) /`a`b!(1 3 6;4 9 15)
+\+`a`b!(1 2 3;4 5 6) /+`a`b!(1 3 6;4 9 15)
*/18#1b /1b
*/3<!18 /0b
*/111b /1b
*/'(0#0b;000b;001b;111b) /1001b
,/(1 2 3;,4) /1 2 3 4
,/() /()
,/!0 /!0
,/0#0. /0#0.
&/'(0#0b;000b;001b) /100b
|/'(0#0b;000b;111b;18#0b;0b,18#1b) /00101b
|/0#0. /-0w
|/`x`z`y /`z
|/0#` /`
&/1.1 2.2 /1.1
&/`x`z`y /`x
-2.&/1.1 2.2 /-2.
+\10010b /1 1 1 2 2
4+\5 2 1 /9 11 12
-\5 2 3 1 /5 3 0 -1
3-\5 2 3 1 /-2 -4 -7 -8
*\8 2 1 -3 /8 16 16 -48
+\!0 /!0
-\1 /1
3-\1 /2
|\00b /00b
|\0001101b /0001111b
&\111010101b /111000000b
-'1. 2. /-1. -2.
{x+y*z}'[1 2 3;4;5 6 7] /21 26 31
-'[3;1] /2
f:{-x};f'1 2;f'1 2;f'1 2 /-1 -2
1 2 3'0 1 2 3 4 /-1 0 1 2 2
1 2 3.0'0 1 2 3 4. /-1 0 1 2 2
"abc"'"Aabcd" /-1 0 1 2 2
v:{$[1<#x;j@*x;x]};j:{v x};j(1;2 3) /1
-'1 2a /1a180 2a180
-'!5 /0 -1 -2 -3 -4
-':3 2 4 0 /0 -1 2 -4
2-':3 2 4 0 /1 -1 2 -4
=':3 2 2 3 /1010b
2~':3 2 2 3 /0010b
-':,1 /,1
3-':,1 /,-2
{x-y}':3 2 1 /0 -1 -1
&':1 1 0 0 1 1 /1 1 0 0 0 1
&':110011b /1 1 0 0 0 1
=':0 0 1 1 1 0 1 /1101100b
=':0011101b /1101100b
3.%':12 60. /4. 5.
3.%':12 60 /(4.;5)
+':1a 1a90 1a270 /2a 1.414213a45 0a
<':001100b /000010b
2<':3 4 3 /001b
2>':3 4 3. /110b
=':1 0w 0n 0n 2 2. /100101b
1(*-)/:(1;2 3) /0 -1
-/:[1;2 3] /-1 -2
1{x-y}/:`a`b!(1 2;3 4) /`a`b!(0 -1;-2 -3)
(3+2) /5
(1+2)*3 /9
() /()
*() /""
+/() /""
2#() /("";"")
(1+2;2) /3 2
(;1) /(;1)
(1;;) /(1;;)
1 2@3 /0N
1 2@1b /2
1 2@10b /2 1
"abc"@!5 /"abc  "
3 4 5@(1;(2;0 1)) /(4;(5;3 4))
1 2 3.@-1 2 /0n 3.
1 2 3a@-1 2 /0na 3a
(2 1)0 /2
1 2 3[1] /2
1 2 3[2 1][0] /3
1@2 3 4 /1 1 1
(`0).(`x;5);x /5
.[1 2;0] /1
.[(1 2;3 4);1 0] /3
(3^!9).(,1;0 1) /,3 4
(1 2;3 4)[1;0] /3
(1 2;3 4)[1;1 0] /4 3
(1 2;3 4)[1 0;1] /4 2
(1 2;3 4)[;1] /2 4
(`a`b!1 2)`a /1
(`a`b!1 2)`b`a /2 1
(+`a`b`c!(1 2 3;4 5 6;7 8 9))[;`b`c] /(4 5 6;7 8 9)
(+`a`b!(1 2;3 4))1 /`a`b!2 4
(+`a`b!(1 2;3 4))1 0 /+`a`b!(2 1;4 3)
(+`a`b`c!(1 2 3;4 5 6;7 8 9))1 2 /+`a`b`c!(2 3;5 6;8 9)
(+`a`b!(1 2;3 4))[`a][1] /2
@[!5;2 3;9] /0 1 9 9 4
@[1 2;0;3] /3 2
@[1 2;1;+;1] /1 3
@[1 2 3.;1 2;4.] /1. 4. 4.
@[1 2 3a;1 2;4a] /1a 4a 4a
.[2^!6;1 2;9] /(0 1 2;3 4 9)
.[2^!6;1 2;9 9] /(0 1 2;(3;4;9 9))
.[2^!6;1 2;+;1] /(0 1 2;3 4 6)
.[2^!6;(1;1 2);+;1] /(0 1 2;3 5 6)
.[2^!6;1 2;+;1 2] /(0 1 2;(3;4;6 7))
.[3^!9;(1 2;0 2);(0 1;2 3)] /(0 1 2;0 4 1;2 7 3)
.[3^!9;(1 2;0 2);-1] /(0 1 2;-1 4 -1;-1 7 -1)
.[3^!9;(;1);9] /(0 9 2;3 9 5;6 9 8)
.[3^!9;(1;);9] /(0 1 2;9 9 9;6 7 8)
.[3^!9;(1;);*;2] /(0 1 2;6 8 10;6 7 8)
.[3^!9;(0 1;);9] /(9 9 9;9 9 9;6 7 8)
.[+`a!,!5;(`a;1);5] /+(,`a)!,0 5 2 3 4
x:2^!6;x[1;2]:0;x /(0 1 2;3 4 0)
x:!5;x[2]:2. 3.;x /(0;1;2. 3.;3;4)
x:!5;x[2 3]:2. 3.;x /(0;1;2.;3.;4)
d:`a`b!1 2;d[`a]:3;d /`a`b!3 2
@[2 3!4 5;3;+;1] /2 3!4 6
d:`a`b!(1 2!3 4;5);d[`a;2 1]:3 4;d /`a`b!(1 2!4 3;5)
+- /+-
(*-)1 2 /-1
+-* /+-*
(&*-)[4;3] /,0
(-1+*)[3;4] /11
(- 1+*)[3;4] /-13
1+ /1+
(1+)2 /3
1+- /1+-
1+/- /1+/-
-1+/- /-1+/-
x:1 /
x:1;x /1
x*x:2 /4
x::1;x /1
x+:x:1;x /2
(1+x;x:3) /4 3
x[1]+:*x:1 2;x /1 3
x:!5;x[2 3]:9;x /0 1 9 9 4
x:!5;x[2 3]+:1 2;x /0 1 3 5 4
f:+;f[3;4] /7
f:1+/;f 3 4 5 /13
(-+)[3;5] /-8
1;2 /2
1;;;2 /2
;;2 /2
{x+y} /{x+y}
.{x+y} /((`y;.;`x;.;+);`x`y;"{x+y}";2)
(.{a:3*x;a:5;x:2})1 /`x`a
(.{a:3*x})3 /1
{x+y}[3;4] /7
{z*x+y}[3;4;5] /35
{{x+y}[x;y]}[1;2] /3
$23 /"23"
$(+':) /"+':"
`i$"31" /31
`F$"1 2 3" /1. 2. 3.
`I$/:("1";"") /(,1;!0)
`$"abc" /`abc
`b`B`i`i`s`S`f$'("0b";"101b";"3";"3b";"`abc";"``a`b";"3.14") /(0b;101b;3;;`abc;``a`b;3.14)
`I`Z$\:"3 4 5" /(3 4 5;3a 4a 5a)
`k@23 /"23"
`l 11 22 33 /,"11 22 33"
`l@`a`be!122 3 /("`a |122";"`be|3")
`l@+`a`b!(11 2;3 4) /("a  b";"----";"11 3";"2  4")
`l@`a!+`a`b!(1 2;3 4) /("a|b";"-|-";"1|3";"2|4")
`l@`x`yy!+`a`b!(1 2;3 4) /("  |a b";"--|---";"x |1 3";"yy|2 4")
`c 1 2 3 /0x010000000200000003000000
`b 1 2 3 256 /1000100010000100b
`z 1 1. /,1.414213a45
1>2 /0b
1<2. /1b
1<2a /1b
3 4<`a`b!5 2 /`a`b!10b
`x=`y`x`z /010b
"alphabetagamma"="m" /00000000000110b
+/"alpha"="a" /2
+/"abc"="rbx" /1
1101001b~1101001b /1b
(!18)~!18 /1b
(!18)~1_!19 /0b
"0123X56789abcdefgh"~"0123456789abcdefgh" /0b
"0123456789abcdeXgh"~"0123456789abcdefgh" /0b
(1 2.;1 2.0)~'(1 3.;1 2.) /01b
(+/;+/;1+*;{x+y};{x+y};1+;1+)~'(+/;-/;1+*;{x+y};{x-y};1+;1-) /1011010b
(1 2;1 2;1a;0n 0w;1 2a90;`a`b!2 3)~'(1 2;2 1;2a;0n 0w;1 2a90;`a`b!2 3) /100111b
*|!100000 /99999
*`a`b!2 1 /2
!`a`b!1 2 /`a`b
!+`a`b!(1 2;3 4) /`a`b
.`a`b!1 2 /1 2
^2 1 3 4 /1 2 3 4
^1 5 8 2 -5 /-5 1 2 5 8
^10010010b /00000111b
^"abracadabra" /"aaaaabbcdrr"
^("alpha";"abc";"";"ab") /("";"ab";"abc";"alpha")
^-2.1 1.1 0w 0n -0w /0n -0w -2.1 1.1 0w
0n<0w /1b
0n>0w /0b
{x@<x}@-2.1 1.1 0w 0n -0w 0.5 /0n -0w -2.1 0.5 1.1 0w
{x@>x}@-2.1 1.1 0w 0n -0w 0.5 /0w 1.1 0.5 -2.1 -0w 0n
<("alpha";"abc";"";"ab") /2 3 1 0
^-!8 /-7 -6 -5 -4 -3 -2 -1 0
>3 4 /1 0
>4 3 /0 1
<1 8 1 2 5 9 /0 2 3 4 1 5
>1 8 1 2 5 9 /5 1 4 3 0 2
<0 3 2 5 6 5 1 4 9 2 8 0 2 1 3 7 1 5 8 1 8 3 3 1 5 5 4 2 6 1 /0 11 6 13 16 19 23 29 2 9 12 27 1 14 21 22 7 26 3 5 17 24 25 4 28 15 10 18 20 8
>0 3 2 5 6 5 1 4 9 2 8 0 2 1 3 7 1 5 8 1 8 3 3 1 5 5 4 2 6 1 /8 10 18 20 15 4 28 3 5 17 24 25 7 26 1 14 21 22 2 9 12 27 6 13 16 19 23 29 0 11
$[1;2;3+4] /2
$[1>2;2;3+4] /7
{$[x>3; :1+x;x];x}'!5 /0 1 2 3 5
while[1-1;2+2] /
n:2;x:3;2*while[n>0;n-:1;x+:1] /10
(3^!9)dot!3 /5 14 23
A:(3 5 -8 12.;-2 3 3 0.;7 -8 2 1.);b:(+A)dot x:1 2 3;0.00001>|/abs x-A solve b /1b
A:+0a0+(1 -2a90 3;5a90 3 2;2 3 1;4 -1 1);0.0001>|/abs r-A solve(+A)dot r:1a30 2a30 3a30 /1b
0 2^+`a`b!(1 2 3 4;5 6 7 8) /(+`a`b!(1 2;5 6);+`a`b!(3 4;7 8))
t:+`a`b`c!(1 2 3 4;2 2 3 3;4 3 2 1);t@=t`b /2 3!(+`a`b`c!(1 2;2 2;4 3);+`a`b`c!(3 4;3 3;2 1))
t:+`a`b`c!(1 2 3 4;2 2 3 3;4 3 2 1);(`b_t)@=t`b /2 3!(+`a`c!(1 2;4 3);+`a`c!(3 4;2 1))
`b?+`a`b`c!(1 2 3 4;2 2 3 3;4 3 2 1) /2 3!(+`a`c!(1 2;4 3);+`a`c!(3 4;2 1))
{a+b}.`a`b!1 2 /3
(+`a`b!(1 2 3;4 5 6)){a>1} /+`a`b!(2 3;5 6)
+/`b?+`a`b`c!(1 2 3 4;2 2 3 3;4 3 2 1) /+`a`c!(4 6;6 4)
`b!+`a`b`c!(1 2;3 4;5 6) /(+(,`b)!,3 4)!+`a`c!(1 2;5 6)
*`a!+`a`b!(1 2;3 4) /(,`b)!,3
ej[`k;+`k`a!(`x`x`y`z;1 2 3 4);+`k`b!(`x`y;`p`q)] /+`k`a`b!(`x`x`y;1 2 3;`p`p`q)
x:10000?20;0.1>abs(1;9.5;20*%1%12.)-(20=#?x;avg x;std x) /111b
x:?10000;0.01>abs(0.5;%1%12.)-(avg x;std x) /11b
x:?-10000;0.02>abs(0 1)-(avg x;std x) /11b