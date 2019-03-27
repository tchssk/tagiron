package main

import (
	"math/rand"
)

type Questions []string

const (
	QuestionSumOfAll         = "数字タイルすべての数の合計は？"
	QuestionSumOfLowerThree  = "小さいほうから３枚の数の合計は？"
	QuestionSumOfCenterThree = "中央の３枚の数の合計は？"
	QuestionSumOfUpperThree  = "大きいほうから３枚の数の合計は？"
	QuestionSumOfRed         = "赤の数の合計は？"
	QuestionSumOfBlue        = "青の数の合計は？"
	QuestionDifference       = "数字タイルの最大の数から、最小の数を引いた数は？"
	QuestionOdd              = "奇数は何枚ある？"
	QuestionEven             = "偶数は何枚ある？（０も含む）"
	QuestionNumberPairs      = "同じ数字タイルのペアは何組ある？"
	QuestionColorPairs       = "同じ色がとなり合っている数字タイルはどこ？"
	QuestionCenter           = "中央の数字タイルは５以上？４以下？"
	QuestionSerial           = "数が連続している数字タイルはどこ？"
	QuestionRedTiles         = "赤の数字タイルは何枚ある？"
	QuestionBlueTiles        = "青の数字タイルは何枚ある？"
	QuestionZero             = "０はどこ？"
	QuestionOneOrTwo         = "１または２はどこ？（どちらかひとつ選ぶ）"
	QuestionThreeOrFour      = "３または４はどこ？（どちらかひとつ選ぶ）"
	QuestionFive             = "５はどこ？"
	QuestionSixOrSeven       = "６または７はどこ？（どちらかひとつ選ぶ）"
	QuestionEightOrNine      = "８または９はどこ？（どちらかひとつ選ぶ）"
)

func (q *Questions) Shuffle() {
	qq := *q
	for i := range qq {
		j := rand.Intn(i + 1)
		qq[i], qq[j] = qq[j], qq[i]
	}
	q = &qq
}

func (q *Questions) Pull(i, j int) Questions {
	qq := *q
	if len(qq) == 0 {
		return Questions{}
	}
	if i < 0 {
		i = 0
	}
	if j < 1 {
		j = 1
	}
	if max := len(qq) - 1; i > max {
		i = max
	}
	if max := len(qq); j > max {
		j = max
	}
	if i > j {
		i = j - 1
	}
	questions := qq[i:j]
	var qqq Questions
	qqq = append(qqq, qq[:i]...)
	qqq = append(qqq, qq[j:]...)
	new := make(Questions, len(qqq))
	copy(new, qqq)
	*q = new
	return questions
}

func (q *Questions) Push(questions Questions) {
	qq := *q
	qq = append(qq, questions...)
	n := make(Questions, len(qq))
	copy(n, qq)
	*q = n
}

func NewQuestions() Questions {
	return Questions{
		QuestionSumOfAll,
		QuestionSumOfLowerThree,
		QuestionSumOfCenterThree,
		QuestionSumOfUpperThree,
		QuestionSumOfRed,
		QuestionSumOfBlue,
		QuestionDifference,
		QuestionOdd,
		QuestionEven,
		QuestionNumberPairs,
		QuestionColorPairs,
		QuestionCenter,
		QuestionSerial,
		QuestionRedTiles,
		QuestionBlueTiles,
		QuestionZero,
		QuestionOneOrTwo,
		QuestionThreeOrFour,
		QuestionFive,
		QuestionSixOrSeven,
		QuestionEightOrNine,
	}
}
