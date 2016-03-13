# h**ID**e  [![Build Status](https://drone.io/github.com/c2h5oh/hide/status.png)](https://drone.io/github.com/c2h5oh/hide/latest)  [![GoDoc](https://godoc.org/github.com/c2h5oh/hide?status.svg)](https://godoc.org/github.com/c2h5oh/hide)  [![Go Report Card Badge](http://goreportcard.com/badge/c2h5oh/hide)](http://goreportcard.com/report/c2h5oh/hide)

Super easy ID obfuscation that actually works

## The why
Using auto-assigned IDs from database is far from ideal, because it provides a lot of information that can be exploited

* max value leaks how many items exist
* max value change leaks how fast new ones are added
* iterating over ID makes scraping all items easy


## The how

### Discarded ideas:
* base64 encode Ids (like Youtube ones) - it's obvious, super easy to reverse - even without coding: `10a` is followed by `10b`
* UUIDs - don't maintain order on sort, not the best pick for primary keys in relational databases - slower than integers, especially on joins, use more storage, partitioning is harder
* random IDs - don't maintain order on sort, the more values you get the more times you will have a duplicate and will have to generate another, results in sparse indexes, partitioning is harder


### So what do we want?
* well obfuscated - hard to figure out from outside even if you know the method
* integer IDs, at least in the database. Preferably still consecutive or at least not sparse
* maintained order (older ID < newer ID)
* as little overhead as possible
* as little code changes required


### Solution

Three words: `Modular multiplicative inverse`. Math warning: https://en.wikipedia.org/wiki/Modular_multiplicative_inverse

* To obfuscate ID calculate MMI of the ID using a large prime number
* To deobfuscate ID calculate MMI of the obfuscated ID using coprime of the previously used prime
* You can still use consecutive integers as IDs, with all of the benefits
* Obfuscated IDs look random
* Figuring which prime was used is not easy and brute-forcing it will be hard - even for `int32` there are close to **200.000.000** primes to choose from


# Usage example

Before:
```go
type User struct {
    ID int64 `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
}
```

After:
```go
import "github.com/c2h5oh/hide"

type User struct {
    ID hide.Int64 `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
}
```
That's it. Really. ID will be transparently obfuscated on `json.Marshal` and deobfuscated on `json.Unmarshal`

**(∩｀-´)⊃━☆ﾟ.*･｡ﾟ**

Also supported: `int32`, `uint32`, `int64`, `uint64`


# Remember to set your own primes!
Package comes with default primes set, but please pick your own. Good source: https://primes.utm.edu/lists/small/small.html
```go
hide.Default.SetInt32(myInt32Prime)   // set prime used for int32 obfuscation
hide.Default.SetUint32(myUint32Prime) // set prime used for uint32 obfuscation
hide.Default.SetInt64(myInt64Prime)   // set prime used for int64 obfuscation
hide.Default.SetUint64(myUint64Prime) // set prime used for uint64 obfuscation

```
