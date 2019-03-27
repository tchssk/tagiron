package main

import "testing"

func TestQuestionsPull(t *testing.T) {
	cases := []struct {
		name     string
		i        int
		j        int
		expected Questions
		pulled   Questions
	}{
		{
			name: "pull 1 question from the head",
			i:    0,
			j:    1,
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
			name: "pull 2 questions from the tail",
			i:    18,
			j:    20,
			expected: Questions{
				QuestionSixOrSeven,
				QuestionEightOrNine,
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
			},
		},
		{
			name: "pull 3 questions from 2nd from the head",
			i:    1,
			j:    4,
			expected: Questions{
				QuestionSumOfCenterThree,
				QuestionSumOfUpperThree,
				QuestionSumOfRed,
			},
			pulled: Questions{
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
			name: "pull 4 questions from 2nd from the tail",
			i:    10,
			j:    14,
			expected: Questions{
				QuestionBlueTiles,
				QuestionZero,
				QuestionOneOrTwo,
				QuestionThreeOrFour,
			},
			pulled: Questions{
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
			name: "pull 1 question from the middle",
			i:    2,
			j:    3,
			expected: Questions{
				QuestionDifference,
			},
			pulled: Questions{
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
			name: "pull 11 questions but remaining only 10 questions",
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
			pulled: Questions{},
		},
		{
			name:     "pull but no question",
			i:        0,
			j:        1,
			expected: Questions{},
			pulled:   Questions{},
		},
	}
	questions := NewQuestions()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := questions.Pull(tc.i, tc.j); len(actual) != len(tc.expected) {
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
