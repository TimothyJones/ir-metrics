package binrel

import (
	"math"
)


func ERR(run Ranking, depth, R uint) float64 {
	p := float64(1)
	err := float64(0)
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i) // the gain function we use is 0 -> 0, 1 -> 1
		if score != 0 {
			score = 1
		}
		err += p * (float64(score) / float64(i+1))
		p = p * float64(1-score)
	}
	return err
}

func SP(run Ranking, depth, R uint) float64 {
	sp := float64(0)
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i) // the gain function we use is 0 -> 0, 1 -> 1
		if score != 0 {
			sp += float64(1) * PatK(run, i+1, R)
		}
	}

	return sp
}

// returns the mrr score as a float64
func MRR(run Ranking, depth, R uint) float64 {
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			return float64(1) / float64(i+1)
		}
	}
	return 0
}

func R_at_K(run Ranking, depth, R uint) float64 {
	found := 0
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			found++
		}
	}

	return float64(found) / float64(R)
}

func PatK(run Ranking, depth, R uint) float64 {
	found := 0
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			found++
		}
	}

	return float64(found) / float64(depth)
}

func RPrec(run Ranking, depth, R uint) float64 {
	if depth <= R {
		return PatK(run, depth, R)
	} else {
		return PatK(run, depth, depth)
	}
}

// returns the rbp0.95 score as a float64
func RBP95(run Ranking, depth, R uint) float64 {
	rbp := float64(0)
	p := 0.95
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			rbp += float64(1) * (math.Pow(p, float64(i)))
		}
	}
	return rbp * (float64(1) - p)
}

// returns the rbp0.50 score as a float64
func RBP50(run Ranking, depth, R uint) float64 {
	rbp := float64(0)
	p := 0.50
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			rbp += float64(1) * (math.Pow(p, float64(i)))
		}
	}
	return rbp * (float64(1) - p)
}

// returns the rbp0.85 score as a float64
func RBP85(run Ranking, depth, R uint) float64 {
	rbp := float64(0)
	p := 0.85
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			rbp += float64(1) * (math.Pow(p, float64(i)))
		}
	}
	return rbp * (float64(1) - p)
}

// returns the sdcg score as a float64
func SDCG(run Ranking, depth, R uint) float64 {
	dcg := float64(0)
	weight := float64(0)
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			// This is a relevant document
			dcg += (math.Pow(2, 1) - 1) / (math.Log2(float64(i + 2)))
		}
		weight += float64(1) / math.Log2(float64(i+2))
	}
	return dcg / weight
}

// returns the sndcg score as a float64
func SNDCG(run Ranking, depth, R uint) (dcg float64) {
	dcg = 0
	n := float64(0)
	found := 0
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			// This is a relevant document
			dcg += (math.Pow(2, 1) - 1) / (math.Log2(float64(i + 2)))
			found++
		}
	}
	for ; found > 0; found-- {
		n += (math.Pow(2, 1) - 1) / math.Log2(float64(found+1))
	}
	dcg /= n
	return
}

// returns the ndcg score as a float64
func NDCG(run Ranking, depth, R uint) (dcg float64) {
	dcg = 0
	n := float64(0)

	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			// This is a relevant document
			dcg += (math.Pow(2, 1) - 1) / (math.Log2(float64(i + 2)))
		}
		if R > 0 {
			n += (math.Pow(2, 1) - 1) / (math.Log2(float64(i + 2)))
			R--
		}
	}
	dcg /= n
	return
}

// returns the dcg score as a float64
func DCG_log(run Ranking, depth, R uint) (dcg float64) {
	dcg = 0

	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			// This is a relevant document
			dcg += (math.Pow(2, 1) - 1) / (math.Log2(float64(i + 2)))
		}
	}
	return
}

// returns the snap score as a float64
func SNAP(run Ranking, depth, R uint) (ap float64) {
	var found int // the number of rel docs we've found
	ap = 0        // average precision

	// Iterate over each result in the result list
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			// This is a relevant document
			found++
			ap += (1.0 * (float64(found) / float64(i+1)))
		}
	}
	ap /= float64(found)

	return
}

// returns the ap score as a float64
func AP(run Ranking, depth, R uint) (ap float64) {
	var found int // the number of rel docs we've found
	ap = 0        // average precision

	// Iterate over each result in the result list
	for i := uint(0); i < depth; i++ {
		score := run & (1 << i)
		if score != 0 {
			// This is a relevant document
			found++
			ap += (1.0 * (float64(found) / float64(i+1)))
		}
	}
	ap /= float64(R)

	return
}

