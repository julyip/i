package cont

func f() {
	var i int32
	for {
		i = (i + 1)
		if i < 2 {
			continue
		}
		i = (i * 2)
	}
}
