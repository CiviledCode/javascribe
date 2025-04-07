package main

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/civiledcode/javascribe/dfa"
	"github.com/t14raptor/go-fast/parser"
)

var testsRan = []string{
	"010", "011", "012", "013", "014", "015", "016", // 01.
	"020", "021", "022", "023", "024", "025", "026", "027", "028", "029", // 02.
	"030", "031", "032", "033", "034", "035", // 03.
}

type testResult struct {
	Identifer string  `json:"id"`
	Assigns   []int64 `json:"assigns"`
}

type testResults struct {
	Expected []testResult `json:"expected"`
}

func TestDFA(t *testing.T) {
	testsRan = []string{"033"}
	for _, testName := range testsRan {
		testFile := testName + ".js"
		testOutput := testName + ".json"

		f, err := os.Open("./js_tests/" + testFile)
		if err != nil {
			panic(err)
		}

		jsCode, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		f.Close()

		f, err = os.Open("./js_tests/" + testOutput)
		if err != nil {
			panic(err)
		}

		results, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		f.Close()

		res := testResults{}

		err = json.Unmarshal(results, &res)
		if err != nil {
			panic(err)
		}

		a, err := parser.ParseFile(string(jsCode))
		if err != nil {
			panic(err)
		}

		rdaCtx := dfa.CreateContextRDA(256)
		//rdaCtx.Debug = true

		rdaCtx.Start(a)

		if len(res.Expected) == 0 {
			t.Fatalf("Test %s JSON invalid", testName)
		}

		for idx, ud := range rdaCtx.UseDefs {
			var nums []int64
			for _, def := range ud.Definitions {
				if def == nil {
					continue
				}

				nums = append(nums, def.Count)
			}

			if len(nums) == 0 {
				nums = []int64{-1}
			}

			expected := res.Expected[idx]

			if expected.Identifer != ud.Usage.Name {
				logFail(expected, testResult{Identifer: ud.Usage.Name, Assigns: nums}, t, testName)

			}

			if len(expected.Assigns) != len(nums) {
				logFail(expected, testResult{Identifer: ud.Usage.Name, Assigns: nums}, t, testName)
			} else {
				for x, num := range nums {
					if num != expected.Assigns[x] {
						logFail(expected, testResult{Identifer: ud.Usage.Name, Assigns: nums}, t, testName)
					}
				}
			}

			t.Logf("PASS: Identifier: %s   Assigns: %v", ud.Usage.Name, nums)
		}

		t.Logf("Test %s PASSED!\n\n", testName)
	}
}

func logFail(expected testResult, got testResult, t *testing.T, testname string) {
	t.Fatalf("incorrect result from test %s.js:\nexpected: id=%s assigns=%v\ngot:      id=%s assigns=%v", testname, expected.Identifer, expected.Assigns, got.Identifer, got.Assigns)
}
