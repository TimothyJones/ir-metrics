package binrel


type Ranking uint

func (run Ranking) StringAtDepth(depth uint) string {
	ret := ""
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			ret += "1 "
		} else {
			ret += "0 "
		}
	}
	return ret
}

func (run Ranking) Value() uint {
	return uint(run)
}
