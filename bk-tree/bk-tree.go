package bktree

type DistanceFunction func(a string, b string) int

type Node struct {
	word     string
	children map[int]*Node
}

type BKTree struct {
	root             *Node
	distanceFunction DistanceFunction
}

func NewBKTree(d DistanceFunction) *BKTree {
	return &BKTree{
		distanceFunction: d,
	}
}

func (t *BKTree) AddWord(wrd string) {
	if t.root == nil {
		t.root = &Node{wrd, make(map[int]*Node)}
	} else {
		t.root.Add(wrd, t.distanceFunction)
	}
}

func (t *BKTree) Search(w string, tolerance int) []string {
	r := make([]string, 0)
	if t.root == nil {
		return r
	}

	candidates := []*Node{t.root}
	for len(candidates) != 0 {
		c := candidates[0]
		candidates = candidates[1:]
		dist := t.distanceFunction(w, c.word)
		if dist <= tolerance {
			r = append(r, c.word)
		}
		low, high := dist-tolerance, dist+tolerance
		for d, n := range c.children {
			if d >= low && d <= high {
				candidates = append(candidates, n)
			}
		}
	}

	return r

}

func (n *Node) Add(w string, d DistanceFunction) {
	dist := d(n.word, w)
	if child, ok := n.children[dist]; ok {
		child.Add(w, d)
	} else {
		n.children[dist] = &Node{w, make(map[int]*Node)}
	}
}
