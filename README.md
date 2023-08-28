# BK-tree in go

Implementation of a simple BK-tree in go, using the [Damerau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance) to calculate the distance between the strings.

A simple CLI tool was also made to use the program.

## Why

I recently wanted to find out how fuzzy string matching works.

I read [this article](https://medium.com/data-science-in-your-pocket/fuzzy-matching-algorithms-explained-e0ff30cc00ca) about fuzzy string matching. I then decided to try and implement a BK-tree in go, using the [Damerau-Levenshtein distance](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance) to calculate the distance.

Code for the Damerau-Levenshtein distance heavily inspired from:

- The [Damerau-Levenshtein go pkg](https://pkg.go.dev/github.com/lmas/Damerau-Levenshtein)
- The [Damerau-Levenshtein Wikipedia article](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)

## CLI

I then built a simple cli to call the program. In order to use the program, pipe data into it. You also must supply a search string. You can also optionally supply a tolerance. The default tolerance is 10.

A testfile is supplied in the repo to showcase example use.

### Usage

- **Without providing a tolerance**

```sh
cat testfile | ./bk-treego -s dawg
```

<details>
<summary>Output</summary>

```
dog
cat
lion
mouse
tiger
turtle
elephant
```

</details>

- **Providing a tolerance**:

```sh
cat testfile | ./bk-tree-go -s dawg -t 2
```

<details>
<summary>Output:</summary>
```sh
dog
```
 </details>

_Note the output is sorted by distance, therefore it is possible to get the closest match using `head`_

```sh
cat testfile | ./bk-tree-go -s dawg | head -1
```

<details>
<summary>Output:</summary>

```sh
dog
```

</details>
