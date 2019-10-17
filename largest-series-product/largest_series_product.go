package lsproduct

import (
	"fmt"
	"strconv"
)

type Product struct {
	result uint64
	digits []uint64
}

func LargestSeriesProduct(dgtsStr string, sp int) (int, error) {
	if sp < 0 {
		return -1, fmt.Errorf("span must be greater than zero")
	}
	if sp > len(dgtsStr) {
		return -1, fmt.Errorf("span larger than digits lenght")
	}
	if sp == 0 {
		return 1, nil
	}

	var dgts []uint64
	var prods []Product
	for _, dgtStr := range dgtsStr {
		n, err := strconv.ParseUint(string(dgtStr), 10, 64)
		if err != nil {
			return -1, err
		}
		dgts = append(dgts, n)
	}

	total := len(dgts)
	if total == sp {
		var pro Product
		pro.digits = append(pro.digits, dgts[0])
		for s := 1; s < sp; s++ {
			pro.digits = append(pro.digits, dgts[s])
		}
		pro.CountProduct()
		return int(pro.result), nil
	}

	for i := 0; i <= total-sp; i++ {
		var pro Product
		pro.digits = append(pro.digits, dgts[i])
		for s := 1; s < sp; s++ {
			pro.digits = append(pro.digits, dgts[i+s])
		}
		pro.CountProduct()
		prods = append(prods, pro)
	}

	var rst uint64
	for _, p := range prods {
		if p.result > rst {
			rst = p.result
		}
	}
	return int(rst), nil
}

func (p *Product) CountProduct() {
	var c uint32
	for _, dig := range p.digits {
		if c == 0 {
			p.result = dig
			c++
			continue
		}
		p.result *= dig
		c++
	}
}
