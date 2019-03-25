package main

import "testing"

func TestQuestionsPull(t *testing.T) {
	cases := []struct {
		name     string
		n        int
		expected Questions
		pulled   Questions
	}{
		{
			name: "pull 1 question",
			n:    1,
			expected: Questions{
				QuestionSumOfAll,
			},
			pulled: Questions{
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
			name: "pull 2 questions",
			n:    2,
			expected: Questions{
				QuestionSumOfLowerThree,
				QuestionSumOfCenterThree,
			},
			pulled: Questions{
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
			name: "pull 3 questions",
			n:    3,
			expected: Questions{
				QuestionSumOfUpperThree,
				QuestionSumOfRed,
				QuestionSumOfBlue,
			},
			pulled: Questions{
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
			name: "pull 4 questions",
			n:    4,
			expected: Questions{
				QuestionDifference,
				QuestionOdd,
				QuestionEven,
				QuestionNumberPairs,
			},
			pulled: Questions{
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
			name: "pull 5 questions",
			n:    5,
			expected: Questions{
				QuestionColorPairs,
				QuestionCenter,
				QuestionSerial,
				QuestionRedTiles,
				QuestionBlueTiles,
			},
			pulled: Questions{
				QuestionZero,
				QuestionOneOrTwo,
				QuestionThreeOrFour,
				QuestionFive,
				QuestionSixOrSeven,
				QuestionEightOrNine,
			},
		},
		{
			name: "pull 6 questions",
			n:    6,
			expected: Questions{
				QuestionZero,
				QuestionOneOrTwo,
				QuestionThreeOrFour,
				QuestionFive,
				QuestionSixOrSeven,
				QuestionEightOrNine,
			},
			pulled: Questions{},
		},
		{
			name:     "pull but no question",
			n:        1,
			expected: Questions{},
			pulled:   Questions{},
		},
	}
	questions := NewQuestions()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := questions.Pull(tc.n); len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v != tc.expected[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
			if len(questions) != len(tc.pulled) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.pulled), len(questions))
			} else {
				for i, v := range questions {
					if v != tc.pulled[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.pulled[i], i)
					}
				}
			}
		})
	}
}

func TestQuestionsPush(t *testing.T) {
	cases := []struct {
		name      string
		questions Questions
		expected  Questions
	}{
		{
			name:      "push no question",
			questions: Questions{},
			expected:  Questions{},
		},
		{
			name: "push 1 question",
			questions: Questions{
				QuestionSumOfAll,
			},
			expected: Questions{
				QuestionSumOfAll,
			},
		},
		{
			name: "push 2 questions",
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
			name: "push 3 questions",
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
			name: "push 4 questions",
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
			name: "push 5 questions",
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
			name: "push 6 questions",
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
			questions.Push(tc.questions)
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
