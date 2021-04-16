package model

import (
	"github.com/sirupsen/logrus"
	"regexp"
)

var TargetExp *regexp.Regexp
var LineExp *regexp.Regexp

func Init(exp string) {
	tmp, err := regexp.Compile(exp)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	TargetExp = tmp

	tmp2, err := regexp.Compile("\n")
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	LineExp = tmp2
}

func FindIndex(source []byte) (res []*Target) {
	targetsPos := TargetExp.FindAllIndex(source, -1)
	linesEndPos := LineExp.FindAllIndex(source, -1)
	j, preLineEnd, curLineEnd, preLineNum := 0, 0, 0, 0

	for _, pos := range targetsPos {
		start := pos[0]
		for j < len(linesEndPos){
			preLineEnd = curLineEnd
			curLineEnd = linesEndPos[j][0]

			if start > curLineEnd {
				j++
				preLineNum = j
			} else {
				break
			}
		}
		res = append(res, &Target{
			LineNum:     preLineNum,
			LineContent: string(source[preLineEnd:curLineEnd]),
		})
	}
	return
}
