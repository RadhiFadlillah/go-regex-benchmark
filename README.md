# Go Regex Benchmark

This repo is a benchmark for various Golang's regular expressions library. Based on benchmark by Rustem Kamalov [here][original-benchmark].

Unlike the original repository, here I only focus on Go language without caring about its performance compared to the other languages.

## Table of Contents

- [Input Text](#input-text)
- [Regex Patterns](#regex-patterns)
  - [Short Regex](#short-regex)
  - [Long Regex](#long-regex)
- [Measure](#measure)
- [Used Packages](#used-packages)
- [Result](#result)
  - [Short Regex](#short-regex-1)
  - [Long Regex](#long-regex-1)

## Input Text

Like the original, the [input text](./input-text.txt) is a concatenation of [Learn X in Y minutes][x-in-y] repository.

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

## Measure

Unlike the original, measuring is done without including regex compilation. So the measurement only focused on pattern matching.

## Used Packages

There are 8 regular expression packages that used in this benchmark:

- **Go**: the `regexp` package from Go's standard library.
- **Grafana**: the package from [`github.com/grafana/regexp@speedup`][grafana] by Bryan Boreham that improve the standard library with several optimations. Actively maintained at the time this benchmark is written.
- **Modernc**: the package from [`modernc.org/regexp`][modernc] that implements experimental DFA support.
- **Regexp2**: the package from [`github.com/dlclark/regexp2`][regexp2] that ports regex engine from .NET frameworks.
- **RE2**: the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using WebAssembly or cgo.
- **Hyperscan**: the package from [`github.com/flier/gohs`][go-hyperscan] that binds Intel's [`hyperscan`][hyperscan] regex engine using cgo.
- **PCRE**: the package from [`github.com/GRbit/go-pcre`][go-pcre] that binds PCRE regex engine using cgo.
- **re2go**: unlike the others, this one use [`re2go`][re2go] to compile regular expressions into Go code. It uses TDFA algorithm.

## Result

The benchmark was run on Linux with Intel i7-8550U with RAM 16 GB.

### Short Regex

|  Package   | Use CGO | Email (ms) | URI (ms) | IP (ms) | Total (ms) |  Times |
| :--------: | :-----: | ---------: | -------: | ------: | ---------: | -----: |
|  RE2 CGO   |   Yes   |       9.27 |    11.72 |    9.49 |      30.48 | 29.86x |
| Hyperscan  |   Yes   |      25.91 |    24.75 |    0.84 |      51.49 | 17.68x |
|    PCRE    |   Yes   |      23.15 |    23.55 |    8.76 |      55.46 | 16.42x |
|   re2go    |         |      54.88 |    37.53 |    9.53 |     101.94 |  8.93x |
|  RE2 WASM  |         |      51.18 |    58.02 |   49.89 |     159.09 |  5.72x |
| Go std lib |         |     261.85 |   249.76 |  398.75 |     910.36 |  1.00x |
|  Modernc   |         |     257.54 |   252.04 |  404.09 |     913.66 |  1.00x |
|  Grafana   |         |     281.45 |   256.05 |  410.55 |     948.05 |  0.96x |
|  Regexp2   |         |    2422.29 |  2190.00 |   78.88 |    4691.17 |  0.19x |

- For short regex, RE2 with cgo is the fastest while Regexp2 is the slowest.
- Regex engines that utilize cgo are way faster than the ones without cgo.
- For code without cgo, re2go has the best performance. However some Go's regex syntaxes are not supported by re2go, so there are need to modify to the regex patterns.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance albeit a bit slower than re2go.
- Regex engine by Modernc that implements DFA is actually a bit slower than the engine from standard library.

### Long Regex

|  Package   | Use CGO | Long Date (ms) |     Times |
| :--------: | :-----: | -------------: | --------: |
| Hyperscan  |   Yes   |           1.13 | 11574.37x |
|  RE2 CGO   |   Yes   |          12.72 |  1031.22x |
|   re2go    |         |          64.69 |   202.76x |
|  RE2 WASM  |         |          64.93 |   202.01x |
|    PCRE    |   Yes   |         177.34 |    73.96x |
|  Grafana   |         |        3324.21 |     3.95x |
|  Regexp2   |         |        4386.10 |     2.99x |
|  Modernc   |         |       12954.90 |     1.01x |
| Go std lib |         |       13116.46 |     1.00x |

- For long regex, Hyperscan with cgo is the fastest while Go standard library is the slowest.
- For code without cgo, re2go has the best performance.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance.
- It's interesting how fast Hyperscan for handling long regex pattern.
- It's interesting to see PCRE's performance become slower for long regex compared to RE2 and Hyperscan. It's even slower than RE2 WASM.
- It's also interesting to see how fast Grafana compared to the Go's standard library.

[original-benchmark]: https://github.com/karust/regex-benchmark
[x-in-y]: https://github.com/adambard/learnxinyminutes-docs
[grafana]: https://github.com/grafana/regexp/tree/speedup?tab=readme-ov-file
[modernc]: https://gitlab.com/cznic/regexp
[regexp2]: https://github.com/dlclark/regexp2
[go-re2]: https://github.com/wasilibs/go-re2
[google-re2]: https://github.com/google/re2
[go-hyperscan]: https://github.com/flier/gohs
[hyperscan]: https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html
[go-pcre]: https://github.com/GRbit/go-pcre
[re2go]: https://re2c.org/manual/manual_go.html
