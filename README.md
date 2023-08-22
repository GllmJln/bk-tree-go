# BK-tree in go

I recently wanted to find out how fuzzy string matching works.

I read [this article](https://medium.com/data-science-in-your-pocket/fuzzy-matching-algorithms-explained-e0ff30cc00ca) about fuzzy string matching. I then decided to try and implement a BK-tree in go, using the [Damerau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance) to calculate the distance.

Code heavily inspired/stolen from:

- The [Demerau-Levenshtein go pkg](https://pkg.go.dev/github.com/lmas/Damerau-Levenshtein)
- The [Demerau-Levenshtein Wikipedia article](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)
