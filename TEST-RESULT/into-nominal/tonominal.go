package intonominal

// * Rp. 100.000
// * Rp. 50.000
// * Rp. 20.000
// * Rp. 10.000
// * Rp. 5.000
// * Rp. 2.000
// * Rp. 1.000
// * Rp. 500
// * Rp. 200
// * RP. 100

func InToNominal(num int) map[string]int {
	var result = make(map[string]int)

	for num > 0 {
		switch {
		case num >= 100000:
			num -= 100000
			result["Rp. 100.000"] += 1
		case num >= 50000:
			num -= 50000
			result["Rp. 50.000"] += 1
		case num >= 20000:
			num -= 20000
			result["Rp. 20.000"] += 1
		case num >= 10000:
			num -= 10000
			result["Rp. 10.000"] += 1
		case num >= 5000:
			num -= 5000
			result["Rp. 5.000"] += 1
		case num >= 2000:
			num -= 2000
			result["Rp. 2.000"] += 1
		case num >= 1000:
			num -= 1000
			result["Rp. 1.000"] += 1
		case num >= 500:
			num -= 500
			result["Rp. 500"] += 1
		case num >= 200:
			num -= 200
			result["Rp. 200"] += 1
		case num >= 100:
			num -= 100
			result["Rp. 100"] += 1
		}
	}

	return result
}
