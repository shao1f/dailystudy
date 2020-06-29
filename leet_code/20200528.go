package main

import (
	"fmt"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

type ByteCount int

func (b *ByteCount) Write(p []byte) (int, error) {
	*b += ByteCount(len(p))
	return len(p), nil
}

func rob(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	}
	maxSum := 0
	for j := 0; j < l; j++ {
		sumTmp := nums[j]
		for i := j; i < l-2; i += 2 {
			sumTmp += nums[i+2]
		}
		if sumTmp > maxSum {
			maxSum = sumTmp
		}
	}
	return maxSum
}

func main() {
	// report := template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := report.Execute(os.Stdout, result); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
