package main

import 
	"strings"
	"sort"
)

func main() {

}

func ExecutePipeline(jobs []job) {
	for job := range jobs {
		go job()
	}
	CombineResults()

}

func SingleHash(in, out chan interface{}) {
	data <- in
	out <- data
}

func MultiHash(in, out chan interface{}) {
	data <- in
	out <- data
}

func CombineResults(in chan interface{}) {
	var results = make([]string)
	for _ = range in {
		results.append(<-in)
	sort.Sort(results)
	return strings.Join(results, "_")
}