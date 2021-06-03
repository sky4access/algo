package array

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type ByContent []Log

func (e ByContent) Len() int      { return len(e) }
func (e ByContent) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e ByContent) Less(i, j int) bool {
	if e[i].Content < e[j].Content {
		return true
	} else if e[i].Content == e[j].Content {
		return e[i].ID < e[j].ID
	}
	return false
}

type Log struct {
	ID      string
	Content string
}

func isDigit(elem string) bool {
	r := []rune(elem)
	return unicode.IsDigit(r[0])
}

func reorderLogFiles(logs []string) []string {
	digitArray := make([]string, 0)
	letterArray := make([]Log, 0)
	// go through loop and find out if it is digit or letter log
	for _, l := range logs {
		// split log into array
		ls := strings.Split(l, " ")
		if isDigit(ls[1]) {
			digitArray = append(digitArray, l)
		} else {
			id := ls[0]
			c := strings.Join(ls[1:len(ls)], " ")
			newLog := Log{
				ID:      id,
				Content: c,
			}
			letterArray = append(letterArray, newLog)
		}
	}
	sort.Sort(ByContent(letterArray))
	fmt.Println(letterArray)
	result := make([]string, 0)
	for _, lr := range letterArray {
		str := lr.ID + " " + lr.Content
		result = append(result, str)
	}
	result = append(result, digitArray...)
	return result
}

func reorderLogFiles2(logs []string) []string {
	sort.SliceStable(sort.StringSlice(logs), func(i, j int) bool {
		log1 := logs[i]
		log2 := logs[j]
		switch {
		case isDigit2(log1[len(log1)-1]) && !isDigit2(log2[len(log2)-1]):
			return false
		case !isDigit2(log1[len(log1)-1]) && isDigit2(log2[len(log2)-1]):
			return true
		case isDigit2(log1[len(log1)-1]) && isDigit2(log2[len(log2)-1]):
			return i < j
		default:
			a, b := strings.IndexRune(log1, ' '), strings.IndexRune(log2, ' ')
			if log1[a+1:] == log2[b+1:] {
				return log1[:a] < log2[:b]
			}
			return log1[a+1:] < log2[b+1:]
		}
	})
	return logs
}

func isDigit2(i uint8) bool {
	if i >= '0' && i <= '9' {
		return true
	}
	return false
}

func reorderLogFiles3(logs []string) []string {

	lLog := []string{}
	dLog := []string{}

	for _, v := range logs {

		temp := strings.Fields(v)
		tmp := []rune(temp[1])

		if tmp[0] < 'a' || tmp[0] > 'z' {
			dLog = append(dLog, v)
			continue
		}
		lLog = append(lLog, v)
	}
	sort.SliceStable(lLog, func(i, j int) bool {

		iIndex := strings.Index(lLog[i], " ")
		jIndex := strings.Index(lLog[j], " ")

		iLog := lLog[i][iIndex:]
		jLog := lLog[j][jIndex:]

		if iLog == jLog {
			return lLog[i] < lLog[j]
		}
		return iLog < jLog
	})

	return append(lLog, dLog...)

}
