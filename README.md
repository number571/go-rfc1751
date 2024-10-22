<p align="center">
    <img src="images/go-rfc1751_logo.png" alt="go-rfc1751_logo.png"/>
</p>

<h2>
	<p align="center">
    	<strong>
	        Mnemonics with RFC1751 format
   		</strong>
	</p>
	<p align="center">
		<a href="https://github.com/topics/golang">
        	<img src="https://img.shields.io/github/go-mod/go-version/number571/go-rfc1751" alt="Go" />
		</a>
		<a href="https://github.com/number571/go-rfc1751/releases">
        	<img src="https://img.shields.io/github/v/release/number571/go-rfc1751.svg" alt="Release" />
		</a>
		<a href="https://github.com/number571/go-rfc1751/blob/master/LICENSE">
        	<img src="https://img.shields.io/github/license/number571/go-rfc1751.svg" alt="License" />
		</a>
		<a href="https://github.com/number571/go-rfc1751/actions">
        	<img src="https://github.com/number571/go-rfc1751/actions/workflows/go.yml/badge.svg" alt="Build" />
		</a>
		<a href="https://github.com/number571/go-rfc1751/blob/c5bd2b4f6efe50674c322175b0d2783f301e4732/Makefile#L43">
        	<img src="test/badge_coverage.svg" alt="Coverage" />
		</a>
	</p>
	<p align="center">
		<a href="https://goreportcard.com/report/github.com/number571/go-rfc1751">
        	<img src="https://goreportcard.com/badge/github.com/number571/go-rfc1751" alt="GoReportCard" />
		</a>
		<a href="https://github.com/number571/go-rfc1751/pulse">
        	<img src="https://img.shields.io/github/commit-activity/m/number571/go-rfc1751" alt="Activity" />
		</a>
		<a href="https://github.com/number571/go-rfc1751/commits/master">
        	<img src="https://img.shields.io/github/last-commit/number571/go-rfc1751.svg" alt="Commits" />
		</a>
		<a href="https://img.shields.io/github/downloads/number571/go-rfc1751/total.svg">
        	<img src="https://img.shields.io/github/downloads/number571/go-rfc1751/total.svg" alt="Downloads" />
		</a>
		<a href="https://godoc.org/github.com/number571/go-rfc1751">
        	<img src="https://godoc.org/github.com/number571/go-rfc1751?status.svg" alt="GoDoc" />
		</a>
	</p>
    About project
</h2>

The [RFC1751](https://datatracker.ietf.org/doc/html/rfc1751) encoding in Go language. The code is mostly rewritten from the Python [implementation](https://github.com/remram44/python-rfc1751). Tested on examples of repositories [remram44/python-rfc1751](https://github.com/remram44/python-rfc1751), [vmizg/rfc1751.js](https://github.com/vmizg/rfc1751.js), [arokettu/php-rfc1751](https://github.com/arokettu/php-rfc1751).

## Installation

```bash
$ go get github.com/number571/go-rfc1751
```

## Requirements

1. Go version `>= 1.16`

## Examples

### Functions

```go
r := rand.New(rand.NewSource(1)) // insecure: used math/rand
fmt.Println(NewMnemonic(r, 128))
// Output: BARK TROD AMY UP LUG KNOB GAS WHEN NEWT POT KEY MEAN
```

```go
b := []byte{204, 172, 42, 237, 89, 16, 86, 190}
fmt.Println(EncodeToString(b))
// Output: RASH BUSH MILK LOOK BAD BRIM
```

```go
s := "RASH BUSH MILK LOOK BAD BRIM"
fmt.Println(DecodeString(s))
// Output: [204 172 42 237 89 16 86 190] <nil>
```

### Applications

```bash
$ go run ./cmd/mnem -size=128
RACY GUN DUN LOP RUSS CODE LENT MOVE DUEL AID SIS BUST
```

## License

Licensed under the MIT License. See [LICENSE](LICENSE) for the full license text.

**[⬆ back to top](#installation)**
