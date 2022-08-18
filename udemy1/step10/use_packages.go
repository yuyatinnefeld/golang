package main

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"time"
)

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {

	// time
	fmt.Println("######### TIME ###########")
	t := time.Now()
	fmt.Println(t)
	rfc3339 := t.Format(time.RFC3339)
	fmt.Println(rfc3339)
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Second())

	// regex
	fmt.Println("######### REGEX ###########")
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([(a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss)
	fmt.Println(fss[0])
	fmt.Println(fss[1])
	fmt.Println(fss[2])

	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss)
	fmt.Println(fss[0])
	fmt.Println(fss[1])
	fmt.Println(fss[2])

	// sort
	fmt.Println("######### SORT ###########")
	i := []int{5, 3, 2, 8, 7}
	s := []string{"a", "d", "b", "c"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Smerdjakow", 28},
		{"Dmitri", 30},
		{"Iwan", 24},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	fmt.Println(i)
	sort.Strings(s)
	fmt.Println(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(p)

	// iota
	fmt.Println("######### IOTA ###########")
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	// context
	fmt.Println("######### CONTEXT ###########")
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go longProcess(ctx, ch)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-ch:
			fmt.Println("success")
			return
		}
	}

}
