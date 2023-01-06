package stools

func ShowVersion() string {
	return "stools_v1.1.0"
}

func SAdd(numA, numB int) int {
	return numA + numB
}

// v1.0.3
func SSub(numA, numB int) int {
	if numA > numB {
		return numA - numB
	}
	return numB - numA
}
