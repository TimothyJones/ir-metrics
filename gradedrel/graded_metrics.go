package gradedrel

import (
	"log"
	"math"
)

func ERR(run Ranking,depth, R uint) float64 {
    p := float64(1)
    err := float64(0)
	for i := uint(0); i < depth ; i++ {
        score := (uint(run) >> (i*2)) & uint(3)
        err += p * (float64(score) / float64(i+1))
        p = p * float64(1-score)
    }
    return err
}

// returns the rbp0.95 score as a float64
func rbp95(run Ranking,depth,R uint) float64 {
	log.Fatal("Unimplemented")
	rbp := float64(0)
	p := 0.95
	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			rbp += float64(1) * (math.Pow(p,float64(i)))
		}
	}
	return rbp * (float64(1) - p)
}
// returns the rbp0.50 score as a float64
func rbp50(run Ranking,depth,R uint) float64 {
	log.Fatal("Unimplemented")
	rbp := float64(0)
	p := 0.50
	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			rbp += float64(1) * (math.Pow(p,float64(i)))
		}
	}
	return rbp * (float64(1) - p)
}
// returns the rbp0.85 score as a float64
func rbp85(run Ranking,depth,R uint) float64 {
	log.Fatal("Unimplemented")
	rbp := float64(0)
	p := 0.85
	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			rbp += float64(1) * (math.Pow(p,float64(i)))
		}
	}
	return rbp * (float64(1) - p)
}

// returns the sdcg score as a float64
func sdcg(run Ranking,depth,R uint) float64 {
	log.Fatal("Unimplemented")
	dcg := float64(0)
	weight := float64(0)
	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			// This is a relevant document
			dcg += (math.Pow(2,1) -1)/(math.Log2(float64(i + 2)))
		}
		weight += float64(1)/math.Log2(float64(i + 2))
	}
	return dcg / weight
}
// returns the sndcg score as a float64
func sndcg(run Ranking,depth,R uint) (dcg float64) {
	log.Fatal("Unimplemented")
	dcg = 0
	n := float64(0)
        found := 0
	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			// This is a relevant document
			dcg +=  (math.Pow(2,1) -1)/(math.Log2(float64(i + 2)))
			found++
		}
	}
        for ; found > 0; found-- {
         	n += (math.Pow(2,1)-1) / math.Log2(float64(found +1))
	}
	dcg /= n
	return
}
// returns the ndcg score as a float64
func ndcg(run Ranking,depth,R uint) (dcg float64) {
	log.Fatal("Unimplemented")
	dcg = 0
	n := float64(0)        

	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			// This is a relevant document
			dcg +=  (math.Pow(2,float64(score)) -1)/(math.Log2(float64(i + 2)))
		}
		if R > 0 {
			n +=(math.Pow(2,1) -1)/(math.Log2(float64(i + 2)))
			R--
		}
	}
	dcg /= n
	return
}

// returns the dcg score as a float64
func DCG_log(run Ranking,depth,R uint) (dcg float64) {
	dcg = 0

	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			// This is a relevant document
			dcg += (math.Pow(2,float64(score)) -1)/(math.Log2(float64(i + 2)))
		}
	}
	return
}
// returns the dcg score as a float64
func DCG_lin(run Ranking,depth,R uint) (dcg float64) {
	dcg = 0

	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			// This is a relevant document
			dcg += (math.Pow(2,float64(score)) -1) * ((float64(depth+1)-float64(i+1))/(float64(depth)));
		}
	}
	return
}
// returns the dcg score as a float64
func DCG_zipf(run Ranking,depth,R uint) (dcg float64) {
	dcg = 0

	for i := uint(0); i < depth ; i++ {
		score := (uint(run) >> (i*2)) & uint(3)
		if(score != 0) {
			// This is a relevant document
			dcg += (math.Pow(2,float64(score)) -1)/float64(i + 1)
		}
	}
	return
}

