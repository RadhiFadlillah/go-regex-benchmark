# Go Regex Benchmark

This repo is a benchmark for various Golang's regular expressions library. Based on benchmark by Rustem Kamalov [here][original-benchmark].

For its input, this benchmark use [text](./input-text.txt) from concatenation of [Learn X in Y minutes][x-in-y] repository. As the original benchmark pointed out, it's not the best representative text for real world benchmark. So, as additional comparison I also made [another benchmark][benchmark-web] that uses around 1000 web pages for input, so feel free to check it out.

## Table of Contents

- [How to Run](#how-to-run)
- [Measurement](#measurement)
- [Regex Patterns](#regex-patterns)
  - [Short Regex](#short-regex)
  - [Long Regex](#long-regex)
- [Used Packages](#used-packages)
  - [Native Go Packages](#native-go-packages)
  - [Regex with CGO Binding](#regex-with-cgo-binding)
  - [Regex with Web Assembly Binding](#regex-with-web-assembly-binding)
  - [Regex Compiler](#regex-compiler)
- [Result](#result)
  - [Short Regex](#short-regex-1)
  - [Long Regex](#long-regex-1)
- [License](#license)

## How to Run

If you are using GNU Make, you can simply run `make` in the root directory of this benchmark:

```
make clean-all    # clean the old build
make build-all    # rebuild the benchmark executable
make              # run the benchmark
```

## Measurement

Unlike the original, measurement is done without including regex compilation, so it's only focused on pattern matching. The measurement are done 10 times, then the smallest durations are used as the final durations.

## Regex Patterns

In this benchmark, there are 2 kinds of regexes that will be tested: short regex and long regex.

### Short Regex

For the short regex, we use 3 patterns from the original benchmark:

- Email: `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`
- URI: `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`
- IPv4: `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`

### Long Regex

For the long regex, we use pattern for detecting multilingual long date texts. Given the following patterns:

- day:

  ```
  [0-3]?[0-9]
  ```

- year:

  ```
  199[0-9]|20[0-3][0-9]
  ```

- month:

  ```
  January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara
  ```

The final pattern is defined as:

```
(?i)({month})\s({day})(?:st|nd|rd|th)?,?\s({year})|({day})(?:st|nd|rd|th|\.)?\s(?:of\s)?({month})[,.]?\s({year})
```

## Used Packages

There are 12 regular expressions that used in this benchmark:

- 5 are regex packages in native Go code.
- 3 are regex packages with CGO binding.
- 1 is regex packages with WASM binding.
- 3 are regex compiler to compile regular expressions to native Go code.

### Native Go Packages

1. **Go** is the `regexp` package from Go's standard library.
2. **Grafana** is the package from [`github.com/grafana/regexp@speedup`][grafana] by Bryan Boreham that improve the standard library with several optimations. Actively maintained at the time this benchmark is written.
3. **Modernc** is the package from [`modernc.org/regexp`][modernc] that implements experimental DFA support. However, it doesn't use DFA implementation from `google/re2` and instead uses its own implementation.
4. **Regexp2** is the package from [`github.com/dlclark/regexp2`][regexp2] that ports regex engine from .NET frameworks.
5. **Code Search** is the package from [`github.com/google/codesearch/regexp`][codesearch]. It uses DFA algorithm with trigram indexing and was used by Google Code Search.

   Since it's used for search engine, currently its API only supports grep-like matching as in for every matching patterns it will returns the entire line instead of only returning the matching phrase.

   It's not thread safe, so every regex can only be used one goroutine at a time. And since its API only supports grep matching, currently it's not suitable for daily use. However it's still benchmarked here to glimpse the possible performance that can be achieved by native Go's regex engine in the future.

### Regex with CGO Binding

1. **RE2 CGO** is the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using cgo. Since Go's regex also use `google/re2` syntax, it can be used as drop in replacement for Go's native regex.
2. **Hyperscan** is the package from [`github.com/flier/gohs`][go-hyperscan] that binds Intel's [`hyperscan`][hyperscan] regex engine using cgo.
3. **PCRE** is the package from [`github.com/GRbit/go-pcre`][go-pcre] that binds PCRE regex engine using cgo.

### Regex with Web Assembly Binding

1. **RE2 WASM** is the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using Web Assembly. Like its cgo counterpart, it can be used as drop in replacement for Go's native regex.

### Regex Compiler

1. **re2go** is a regex compiler from [`re2c.org`][re2c] that compiles regular expressions into a native Go codes. Despite its name, it's not related with `google/re2` and uses its own lookahead TDFA algorithm.

   Since originally it only supports C, `re2go` use its own regex syntax which is not really compatible with Go's regex syntax. Fortunately it also supports [Flex regex syntax][flex] which is kinda similar with Go's syntax, so modifying the existing pattern is pretty easy.

   It's actively maintained and has been used in production and many open source projects, e.g. PHP and Ninja.

2. **Regexp2go** is a regex compiler from [`github.com/CAFxX/regexp2go`][regexp2go]. Despite its name, it's not related with `github.com/dlclark/regexp2` package mentioned above.

   It's similar in spirit to `re2go`, but aiming for compatibility with Go's regex syntax. At the time this benchmark written, this compiler hasn't been updated for 2 years and its documentation mentioned that it's not recommended to use in production.

   However for basic regex and with enough testing I reckon it should be good enough to use.

3. **Regexp2cg** is a regex compiler from [`github.com/dlclark/regexp2cg`][regexp2cg]. It's related with `github.com/dlclark/regexp2` package mentioned above.

   It will compile regular expressions into Go codes which can be used by the `dlclark/regexp2` package.

## Result

As reminder, as mentioned above this benchmark might not be the best representative text for real world. So make sure to check my [other benchmark][benchmark-web] that uses around 1000 web pages for its input.

With that out of the way, here is the benchmark result. The benchmark was run on my Linux PC with Intel i7-8550U and RAM 16 GB.

### Short Regex

|   Package   |     Type      | Email (ms) | URI (ms) | IP (ms) | Total (ms) |  Times |
| :---------: | :-----------: | ---------: | -------: | ------: | ---------: | -----: |
|   RE2 CGO   |      CGO      |       9.34 |    11.51 |    9.23 |      30.07 | 29.42x |
| Code Search | Native (Grep) |      11.50 |    12.58 |   11.82 |      35.90 | 24.65x |
|  Hyperscan  |      CGO      |      26.00 |    25.73 |    0.83 |      52.56 | 16.83x |
|    PCRE     |      CGO      |      22.84 |    23.51 |    8.67 |      55.01 | 16.08x |
|    re2go    |   Compiler    |      47.45 |    35.47 |    8.00 |      90.92 |  9.73x |
|  RE2 WASM   |     WASM      |      50.55 |    56.09 |   48.67 |     155.31 |  5.70x |
|  Regexp2Go  |   Compiler    |      67.56 |   556.80 |  168.23 |     792.59 |  1.12x |
|   Modernc   |    Native     |     243.07 |   236.55 |  380.77 |     860.39 |  1.03x |
|     Go      |    Native     |     241.85 |   248.85 |  394.04 |     884.74 |  1.00x |
|   Grafana   |    Native     |     261.77 |   251.27 |  402.06 |     915.10 |  0.97x |
|  Regexp2cg  |   Compiler    |    1948.30 |  1772.00 |   68.97 |    3789.26 |  0.23x |
|   Regexp2   |    Native     |    2325.68 |  2061.28 |   77.81 |    4464.77 |  0.20x |

Some interesting points:

- It's amazing to see how fast Code Search is. It's almost as fast as RE2 with cgo.
- The speed of the code generated by the re2go appears to be influenced by the length of its regex pattern. For the long-ish pattern like IP, its performance is even faster than RE2 with cgo. However, since some Go's regex syntaxes are not supported by re2go there are needs to modify the regex patterns before using it.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance albeit a bit slower than re2go.

### Long Regex

|   Package   |     Type      | Long Date (ms) |     Times |
| :---------: | :-----------: | -------------: | --------: |
|  Hyperscan  |      CGO      |           1.14 | 11332.36x |
|   RE2 CGO   |      CGO      |          13.09 |   988.39x |
|    re2go    |   Compiler    |          42.97 |   301.15x |
| Code Search | Native (Grep) |          45.67 |   283.34x |
|  RE2 WASM   |     WASM      |          59.29 |   218.22x |
|    PCRE     |      CGO      |         165.39 |    78.24x |
|   Grafana   |    Native     |        3386.20 |     3.82x |
|   Regexp2   |    Native     |        4302.07 |     3.01x |
|  Regexp2cg  |   Compiler    |        5250.58 |     2.46x |
|  Regexp2Go  |   Compiler    |        7451.24 |     1.74x |
|   Modernc   |    Native     |       12645.68 |     1.02x |
|     Go      |    Native     |       12939.35 |     1.00x |

Some interesting points:

- Hyperscan is really fast at handling long regex pattern. It's even faster than when it's used for short regex.
- PCRE's performance become a lot slower for long regex compared to RE2 and Hyperscan. It's even slower than RE2 WASM.
- For native Go code, Grafana is pretty fast at handling long regex pattern.
- For code without cgo, regex that compiled by re2go has the best performance. It has similar performance as Code Search.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance (since Code Search currently can't be used as daily regex engine).

## License

Like the original benchmark, this benchmark is also released under MIT license.

[original-benchmark]: https://github.com/karust/regex-benchmark
[benchmark-web]: https://github.com/RadhiFadlillah/go-regex-benchmark-web
[x-in-y]: https://github.com/adambard/learnxinyminutes-docs
[grafana]: https://github.com/grafana/regexp/tree/speedup?tab=readme-ov-file
[modernc]: https://gitlab.com/cznic/regexp
[regexp2]: https://github.com/dlclark/regexp2
[codesearch]: https://github.com/google/codesearch
[go-re2]: https://github.com/wasilibs/go-re2
[google-re2]: https://github.com/google/re2
[go-hyperscan]: https://github.com/flier/gohs
[hyperscan]: https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html
[go-pcre]: https://github.com/GRbit/go-pcre
[re2c]: https://re2c.org/manual/manual_go.html
[flex]: https://github.com/westes/flex
[regexp2go]: https://github.com/CAFxX/regexp2go
[regexp2cg]: https://github.com/dlclark/regexp2cg
