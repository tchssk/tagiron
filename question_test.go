package main

import "testing"

func TestQuestionsRemove(t *testing.T) {
	cases := []struct {
		name     string
		i        int
		j        int
		expected Questions
		removeed Questions
	}{
		{
			name: "remove 1 question from the head",
			i:    0,
			j:    1,
			expected: Questions{
				QuestionSumOfAll,
			},
			removeed: Questions{
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
			},
		},
		{
			name: "remove 2 questions from the tail",
			i:    18,
			j:    20,
			expected: Questions{
				QuestionSixOrSeven,
				QuestionEightOrNine,
			},
			removeed: Questions{
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
			},
		},
		{
			name: "remove 3 questions from 2nd from the head",
			i:    1,
			j:    4,
			expected: Questions{
				QuestionSumOfCenterThree,
				QuestionSumOfUpperThree,
				QuestionSumOfRed,
			},
			removeed: Questions{
				QuestionSumOfLowerThree,
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
			},
		},
		{
			name: "remove 4 questions from 2nd from the tail",
			i:    10,
			j:    14,
			expected: Questions{
				QuestionBlueTiles,
				QuestionZero,
				QuestionOneOrTwo,
				QuestionThreeOrFour,
			},
			removeed: Questions{
				QuestionSumOfLowerThree,
				QuestionSumOfBlue,
				QuestionDifference,
				QuestionOdd,
				QuestionEven,
				QuestionNumberPairs,
				QuestionColorPairs,
				QuestionCenter,
				QuestionSerial,
				QuestionRedTiles,
				QuestionFive,
			},
		},
		{
			name: "remove 1 question from the middle",
			i:    2,
			j:    3,
			expected: Questions{
				QuestionDifference,
			},
			removeed: Questions{
				QuestionSumOfLowerThree,
				QuestionSumOfBlue,
				QuestionOdd,
				QuestionEven,
				QuestionNumberPairs,
				QuestionColorPairs,
				QuestionCenter,
				QuestionSerial,
				QuestionRedTiles,
				QuestionFive,
			},
		},
		{
			name: "remove 11 questions but remaining only 10 questions",
			i:    0,
			j:    11,
			expected: Questions{
				QuestionSumOfLowerThree,
				QuestionSumOfBlue,
				QuestionOdd,
				QuestionEven,
				QuestionNumberPairs,
				QuestionColorPairs,
				QuestionCenter,
				QuestionSerial,
				QuestionRedTiles,
				QuestionFive,
			},
			removeed: Questions{},
		},
		{
			name:     "remove but no question",
			i:        0,
			j:        1,
			expected: Questions{},
			removeed: Questions{},
		},
	}
	questions := NewQuestions()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := questions.Remove(tc.i, tc.j); len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v != tc.expected[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
			if len(questions) != len(tc.removeed) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.removeed), len(questions))
			} else {
				for i, v := range questions {
					if v != tc.removeed[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.removeed[i], i)
					}
				}
			}
		})
	}
}

func TestQuestionsAdd(t *testing.T) {
	cases := []struct {
		name      string
		questions Questions
		expected  Questions
	}{
		{
			name:      "add no question",
			questions: Questions{},
			expected:  Questions{},
		},
		{
			name: "add 1 question",
			questions: Questions{
				QuestionSumOfAll,
			},
			expected: Questions{
				QuestionSumOfAll,
			},
		},
		{
			name: "add 2 questions",
			questions: Questions{
				QuestionSumOfLowerThree,
				QuestionSumOfCenterThree,
			},
			expected: Questions{
				QuestionSumOfAll,
				QuestionSumOfLowerThree,
				QuestionSumOfCenterThree,
			},
		},
		{
			name: "add 3 questions",
			questions: Questions{
				QuestionSumOfUpperThree,
				QuestionSumOfRed,
				QuestionSumOfBlue,
			},
			expected: Questions{
				QuestionSumOfAll,
				QuestionSumOfLowerThree,
				QuestionSumOfCenterThree,
				QuestionSumOfUpperThree,
				QuestionSumOfRed,
				QuestionSumOfBlue,
			},
		},
		{
			name: "add 4 questions",
			questions: Questions{
				QuestionDifference,
				QuestionOdd,
				QuestionEven,
				QuestionNumberPairs,
			},
			expected: Questions{
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
			},
		},
		{
			name: "add 5 questions",
			questions: Questions{
				QuestionColorPairs,
				QuestionCenter,
				QuestionSerial,
				QuestionRedTiles,
				QuestionBlueTiles,
			},
			expected: Questions{
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
			},
		},
		{
			name: "add 6 questions",
			questions: Questions{
				QuestionZero,
				QuestionOneOrTwo,
				QuestionThreeOrFour,
				QuestionFive,
				QuestionSixOrSeven,
				QuestionEightOrNine,
			},
			expected: Questions{
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
			},
		},
	}
	questions := Questions{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			questions.Add(tc.questions)
			if actual := questions; len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v != tc.expected[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
		})
	}
}

func TestWhere(t *testing.T) {
	cases := map[string]struct {
		indexes  []int
		expected string
	}{
		"none": {
			indexes:  []int{},
			expected: "なし",
		},
		"left end": {
			indexes:  []int{0},
			expected: "一番左",
		},
		"left end and the right": {
			indexes:  []int{0, 1},
			expected: "一番左と左から２番目",
		},
		"left end and center": {
			indexes:  []int{0, 2},
			expected: "一番左と真ん中",
		},
		"left end and right side of center": {
			indexes:  []int{0, 3},
			expected: "一番左と右から２番目",
		},
		"left end and right end": {
			indexes:  []int{0, 4},
			expected: "一番左と一番右",
		},
		"center and the left": {
			indexes:  []int{1, 2},
			expected: "左から２番目と真ん中",
		},
		"both sides of center": {
			indexes:  []int{1, 3},
			expected: "左から２番目と右から２番目",
		},
		"left side of center and right end": {
			indexes:  []int{1, 4},
			expected: "左から２番目と一番右",
		},
		"center and the right": {
			indexes:  []int{2, 3},
			expected: "真ん中と右から２番目",
		},
		"center and right end": {
			indexes:  []int{2, 4},
			expected: "真ん中と一番右",
		},
		"right end and the left": {
			indexes:  []int{3, 4},
			expected: "右から２番目と一番右",
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := Where(tc.indexes); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}
