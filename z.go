package main

import . "github.com/ktye/wg/module"

func zk() {
	Data(600, "``x`y`z`k`l`a`b`while`\"rf.\"`\"rz.\"`\"uqs.\"`\"uqf.\"`\"gdt.\"`\"lin.\"`\"odo.\"`\"grp.\"\n`\"x.\":{,/+\"0123456789abcdef\"@(x%16;16/x:256/256+x)}\n`\"t.\":`45\n`\"p.\":`46\n`\"b.\":(`46)[`b;]\n`\"c.\":(`46)[`c;]\n`\"i.\":(`46)[`i;]\n`\"s.\":(`46)[`s;]\n`\"f.\":(`46)[`f;]\n`\"z.\":(`46)[`z;]\n`\"uqs.\":{x@&~0b~':x:^x}\n`\"uqf.\":{x@&(!#x)=x?x}\n`\"gdt.\":{[t;g](!#t)($[g;{x@>y x};{x@<y x}])/|.t}\n`\"odo.\":{{y@(#y)/!x}/:[*/x;&'x#'|*\\-1_1,|x]}\n`\"grp.\":{(x@*'g)!g:(&~x~':x i)^i:<x}\nany:`30;abs:`32;sin:`44;cos:`39;find:`31;fill:`38;imag:`33;conj:`34;angle:`35;exp:`42;log:`43\n`\"pad.\":{(|/#'x)#'x}\n`\"l.\":{\nkt:{[x;y;k;T]x:$[`T~@x;T[x;k];`pad(\"\";\"-\"),$x];(x,'\"|\"),'T[y;k]}\nd:{[x;k;kt;T]r:!x;x:.x;$[`T~@x;kt[r;x;k;T];,'[,'[`pad(k'r);\"|\"];k'x]]}\nT:{[x;k]$[`L':@'.x;,k x;(,*x),(,(#*x)#\"-\"),1_x:\" \"/:'+`pad@'$(!x),'.x]}\nt:@x;k:`kxy 1\ndd:(\"\";,\"..\")20<#x:$[(@x)':`L`D`T;x;x~*x;x;[t:`L;,x]]\nx:$[x~*x;x;(20&#x)#x]\n$[`D~t;d[x;k;kt;T];`T~t;T[x;k];x~*x;,k x;k'x],dd}\n`\"str.\":{q:{c,(\"\\\\\"/:(0,i)^@[x;i;(qs!\"tnr\\\"\\\\\")x i:&x':qs:\"\\t\\n\\r\\\"\\\\\"]),c:_34}\n$[|/x':\"\\t\\n\\r\"__!31;\"0x\",`x@x;q x]}\n`\"kxy.\":{\na:{t:@x;x:$x;$[`c~t;`str x;`s~t;\"`\",x;x]}\nd:{[x;k]r:\"!\",k@.x;n:#!x;x:k@!x;$[(1~n)|(@.x)':`D`T;\"(\",x,\")\";x],r}\nv:{[x;k;m;n]m*:(.`\".kstm\")t:@x; dd:(\"\";\"..\")m<#x;x:(m&#x)#x\nx:$[`L~t;k'x;`C~t;x;$x]\nx:$[`B~t;(*'x),\"b\";`C~t;`str x;`S~t;c,(c:\"`\")/:x;`L~t;$[1~n;*x;\"(\",(\";\"/:x),\")\"];\" \"/:x]\n((\"\";\",\")(1~n)),x,dd}\nt:@y;n:#y;k:`kxy x;m:x\n$[n~0;(.`\".kst0\")@t;`T~t;\"+\",d[+y;k];`D~t;d[y;k];y~*y;a y;v[y;k;m;n]]}\n`\".kst0\":`B`C`I`S`F`Z`L!(\"0#0b\";c,c:_34;\"!0\";\"0#`\";\"0#0.\";\"0#0a\";\"()\")\n`\".kstm\":`B`C`I`S`F`Z`L!100 100 30 30 20 10 20\n`\"k.\":`kxy 1000000\n`\"rf.\": {.5+(x?0)%4294967295.}\n`\"rf1.\":{.5+(1.+x?0)%4294967295.}        \n`\"rz.\": {(%-2*log `rf1 x)@360.*`rf x}\n`\"lin.\":{$[`L~@z;(.`\"lin.\")[x;y]'z;[dx:0.+1_-':x;dy:0.+1_-':y;b:(-2+#x)&0|x'z;(y b)+(dy b)*(z-x b)%dx b]]}\n`\"split.\":{$[`L~@x;`split@'x;\" \"\\:$[\" \"=x@-1+#x:x@&~i&~':i:\" \"=x;-1_x;x]]}\ndot:{[xt;y]{+/x*y}\\:[xt;y]}\nsolve:{qslv:{H:x 0;r:x 1;n:x 2;m:x 3;j:0;K:!m\nwhile[j<n;y[K]-:(+/(conj H[j;K])*y K)*H[j;K];K:1_K;j+:1]\ni:n-1;J:!n;y[i]%:r@i\nwhile[i;j:i_J;i-:1;y[i]:(y[i]-+/H[j;i]*y@j)%r@i]\nn#y}\nq:$[`i~@*|x;x;qr x];$[`L~@y;qslv/:[q;y];qslv[q;y]]}\nqr:{K:!m:#*x;I:!n:#x;j:0;r:n#0a;turn:$[`Z~@*x;{(-x)@angle y};{x*1. -1@y>0}]\nwhile[j<n;I:1_I\nr[j]:turn[s:abs@abs/j_x j;xx:x[j;j]]\nx[j;j]-:r[j]\nx[j;K]%:%s*(s+abs xx)\nx[I;K]-:{+/x*y}/:[(conj x[j;K]);x[I;K]]*\\:x[j;K]\nK:1_K;j+:1];(x;r;n;m)}\navg:{(+/x)%0.+#x}\nvar:{(+/x*x:(x-avg x))%-1+#x}\nstd:{%var x}\nrem:{x/x+x/y}\nej:{(y j),'x_z i j:&~0N=i:(z x)?y x}\n`\"pack.\":{w:{(`c@,#x),x};($t),$[`s~t:@x;`pack@$x;x~*x;w `c@,x;`L~@x;(`c@,#x),,/`pack@'x;(@x)':`D`T;(`pack@.x),`pack@!x;`S~t;,/`pack@$x;w `c x]}\n`\"unpack.\":{s:x;g:{[n]r:n#s;s::n_s;r};n:{*`i@g 4};u:{x;$[(t:*g 1)':\"bcifz\";*(`$t)g n[];t~\"s\";`$u 0;t~\"S\";`$u 0;t~\"L\";u'!n[];t~\"D\";(u 0)!u 0;t~\"T\";+(u 0)!u 0;(`$_t+32)g n[]]};u 0}\ncsv:{c:{s:`$'x@i:&x':\"ifzs\";n:`i$\" \"\\:-1_@[x;i;\" \"];y[a]:(y[a],''\"a\"),''y[1+a:&s=`z];s$'y n};s:$[\" \"~(*x);`split@;(*x)\\:];x:1_x;y:+s'$[`L~@y;y;\"\\n\"\\:y];$[#x;c[x;y];y]}\n`\"ceg.\":{(x i)!0-':1+i:&(1_~~':x),1b}\nhist:{$[`i~@x;hist[(x;&/y;|/y);y];(Y;(`38)[0;(`ceg@^1+((d%2.0)+-1_Y:(x 1)+(d:(--/1_x)%-1.+n)*!n)'y)@!n:_0.+*x])]}\n")
	zn := int32(3104) // should end before 8k
	x := mk(Ct, zn)
	Memorycopy(int32(x), 600, zn)
	dx(Val(x))
}
