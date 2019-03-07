package droptoken

import (
	"fmt"
	"strconv"
)

type byGameID []string

func (s byGameID) Len() int {
	return len(s)
}
func (s byGameID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byGameID) Less(i, j int) bool {
	vali, erri := strconv.Atoi(s[i][4:len(s[i])])
	valj, errj := strconv.Atoi(s[j][4:len(s[j])])

	if erri != nil || errj != nil {
		fmt.Println("Could not convert: " + s[i][4:len(s[i])] + "::" + s[j][4:len(s[j])])
		return s[i] < s[j]
	}

	return vali < valj
}

type byMoveNumber []Move

func (s byMoveNumber) Len() int {
	return len(s)
}
func (s byMoveNumber) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byMoveNumber) Less(i, j int) bool {
	return s[i].moveNumber < s[j].moveNumber
}
