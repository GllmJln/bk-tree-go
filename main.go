package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

	bytes, _ := ioutil.ReadAll(os.Stdin)
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

	format := strings.Join(found, "\n")

	fmt.Println(format)
}
