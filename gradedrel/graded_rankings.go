package gradedrel

type Ranking uint

func (run Ranking) StringAtDepth(depth uint) string{
    ret := ""
    for i := uint(0); i < depth ; i++ {
        score := (uint(run) >> (i*2)) & uint(3)
               // 00 - irrel
           // 01 - partially rel
           // 10 - rel
           // 11 - highly rel
           switch score {
                case 0:
                ret += "0 "
            case 1:
                ret += "1 "
            case 2:
                ret += "2 "
            case 3:
                ret += "3 "
            default:
                ret += "? "
           }
    }
    return ret
}

func (run Ranking) Value() uint {
	return uint(run)
}
