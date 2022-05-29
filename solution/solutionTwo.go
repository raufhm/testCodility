package solution

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func SolutionTwo(T []string, R []string) (result float64) {
	lenGroup, err := GetLen(T[0])
	if err != nil {
		log.Error(err)
		return
	}
	groupMap := map[string][]string{}
	inputMap := map[string]string{}

	// do map for each group and map status for each test type
	for i, dataT := range T {
		if _, exist := groupMap[dataT[:lenGroup]]; !exist {
			groupMap[dataT[:lenGroup]] = []string{}
		}
		groupMap[dataT[:lenGroup]] = append(groupMap[dataT[:lenGroup]], dataT)
		inputMap[dataT] = R[i]
	}

	// count score if only OK in each grup
	var score float64
	for _, v := range groupMap {
		var isOK bool
		for _, vv := range v {
			if inputMap[vv] == "OK" {
				isOK = true
			} else {
				isOK = false
				break
			}
		}
		if isOK == true {
			score++
		}
	}

	// count result
	result = math.Floor(score * 100 / float64(len(groupMap)))
	return
}

func GetLen(input string) (length int, err error) {
	re := regexp.MustCompile("[0-9]+")
	number := re.FindAllString(input, -1)
	if len(number) == 0 {
		err = fmt.Errorf("error occured")
		return
	}
	length = strings.IndexAny(input, number[0]) + 1
	return
}
