[![Coverage Status](https://coveralls.io/repos/github/wutzi15/levenshtein/badge.svg?branch=main)](https://coveralls.io/github/wutzi15/levenshtein?branch=main)
[![run tests](https://github.com/wutzi15/levenshtein/actions/workflows/go.yml/badge.svg)](https://github.com/wutzi15/levenshtein/actions/workflows/go.yml)

# Levenshtein
A very simple implementation of the Levenshtein distance in Go.
The levenshtein distance is a measure of the similarity between two strings. It is the minimum number of single-character edits (insertions, deletions or substitutions) required to change one string into the other.[[1](https://en.wikipedia.org/wiki/Levenshtein_distance)]

## Installation

```bash
go get github.com/wutzi15/levenshtein
```


## Usage

```go
func main() {
    fmt.Println(levenshtein.Levenshtein("kitten", "sitting"))
    fmt.Println(levenshtein.Levenshtein("Saturday", "Sunday"))
}

```

## License
MIT License

Copyright (c) 2022 Benedikt Bergenthal

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


## Contributing
Contributing is very welcome!

## References
Based on [this](https://stackabuse.com/levenshtein-distance-and-text-similarity-in-python/).
You can also read the wikipedia article on [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance).