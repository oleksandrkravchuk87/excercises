package main

import (
	"fmt"
	"sort"
	"strings"
)

var commandResults map[string]result

type result struct {
	mp int
	w  int
	d  int
	l  int
	p  int
}

type kv struct {
	key string
	val result
}

func toRows(text string) []string {
	allRows := strings.Split(text, "\n")
	validRows := make([]string, 0)
	for _, v := range allRows {
		if v != "" && !strings.HasPrefix(v, "#") {
			validRows = append(validRows, v)
		}
	}
	return validRows
}

func convert(rows []string) map[string]result {
	commandResults := make(map[string]result)

	for _, v := range rows {
		tempSl := strings.Split(v, ";")

		switch {
		case tempSl[2] == "win":
			{
				winnerResult := commandResults[tempSl[0]]
				winnerResult.mp++
				winnerResult.w++
				winnerResult.p += 3
				commandResults[tempSl[0]] = winnerResult

				looserResult := commandResults[tempSl[1]]
				looserResult.mp++
				looserResult.l++
				commandResults[tempSl[1]] = looserResult
			}
		case tempSl[2] == "loss":
			{
				winnerResult := commandResults[tempSl[1]]
				winnerResult.mp++
				winnerResult.w++
				winnerResult.p += 3
				commandResults[tempSl[1]] = winnerResult

				looserResult := commandResults[tempSl[0]]
				looserResult.mp++
				looserResult.l++
				commandResults[tempSl[0]] = looserResult
			}
		case tempSl[2] == "draw":
			{
				firstCommandResult := commandResults[tempSl[0]]
				firstCommandResult.mp++
				firstCommandResult.d++
				firstCommandResult.p++
				commandResults[tempSl[0]] = firstCommandResult

				secondCommandResult := commandResults[tempSl[1]]
				secondCommandResult.mp++
				secondCommandResult.d++
				secondCommandResult.p++
				commandResults[tempSl[1]] = secondCommandResult
			}

		}

	}
	return commandResults

}

func printResults(results map[string]result) {
	var sl []kv

	for k, v := range results {

		sl = append(sl, kv{key: k, val: v})
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i].val.p > sl[j].val.p
	})

	for _, v := range sl {
		fmt.Printf("%-25s |%d|%d|%d|%d|%d|\n", v.key, v.val.mp, v.val.w, v.val.d, v.val.l, v.val.p)
	}

}

func main() {
	raw := `
Courageous Californians;Devastating Donkeys;win
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;loss
Courageous Californians;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;draw
Allegoric Alaskians;Courageous Californians;draw
`
	rows := toRows(raw)
	res := convert(rows)
	printResults(res)
}
