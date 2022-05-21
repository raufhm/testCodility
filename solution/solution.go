package solution

import "strconv"

func Solution(T []string, R []string) int {
	var result int
	a := T[0]
	length := groupLength(a)
	var groupNumber []string
	var totalGroup []string

	for i, groupName := range T {
		groupNumber = append(groupNumber, string(groupName[length+1]))
		if i != 0 {
			exist := checkExist(totalGroup, string(groupName[length+1]))
			if !exist {
				totalGroup = append(totalGroup, string(groupName[length+1]))
			}
		} else {
			totalGroup = append(totalGroup, string(groupName[length+1]))
		}

	}
	score := checkScore(groupNumber, R)
	result = score * 100 / len(totalGroup)
	//check decimal point

	return result
}

func checkExist(totalGroup []string, inp string) bool {
	for _, check := range totalGroup {
		if check == inp {
			return true
		}
	}
	return false
}

func checkScore(groupNumber []string, R []string) (score int) {
	var nilai bool
	for i := 0; i < len(R); i++ {
		nilai = false
		if R[i] == "OK" {
			nilai = true
			for j := 0; j < len(R); j++ {
				if j == i {
					continue
				}
				if groupNumber[i] == groupNumber[j] {
					if R[j] == "OK" {
						nilai = true
					} else {
						nilai = false
						break
					}
				}
			}
			if nilai == true {
				score += 1
			}
		}
	}
	return score
}

func groupLength(input string) (length int) {
	i := len(input) - 1 // test1a
	for {
		if i < 0 {
			break
		}
		if _, err := strconv.Atoi(string(input[i])); err == nil {
			length = i - 1
			break
		}
		i--
	}

	return length
}
