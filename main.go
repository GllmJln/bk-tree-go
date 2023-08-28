package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	bktree "github.com/gllmjln/bk-tree-go/bk-tree"
	distance "github.com/gllmjln/bk-tree-go/distance"
)

func main() {

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("You must pipe data into the program.")
		os.Exit(1)
	}

	bytes, _ := io.ReadAll(os.Stdin)
	str := string(bytes)

	targetPtr := flag.String("s", "", "the string to find.")
	tolerancePtr := flag.Int("t", 10, "the tolerance of the search.")

	flag.Parse()

	if len(*targetPtr) == 0 {
		fmt.Println("You must provide a target string to find.")
		os.Exit(1)
	}

	bkt := bktree.NewBKTree(distance.CalculateDistance)

	for _, w := range strings.Split(str, "\n") {
		bkt.AddWord(w)
	}

	found := bkt.Search(*targetPtr, *tolerancePtr)

	sort.SliceStable(found, func(i, j int) bool { return found[i].Distance < found[j].Distance })

	sortedResult := make([]string, 0)
	for _, wrd := range found {
		sortedResult = append(sortedResult, wrd.Word)
	}

	format := strings.Join(sortedResult, "\n")

	fmt.Println(format)
}
