## TimeID

_Generate unique IDs which contain a high-resolution timestamp._

Package to generate unique, 64bit integers based on the current nanosecond time. Safe for concurrent use by multiple goroutines. A monotonically increasing int64 is always returned.

## Usage

This package is designed for self-contained applications that need to generate unique identifiers without relying on an external ID source such as Snowflake (or a database like MySQL).

This may generate duplicates numbers if used by multiple application instances (such as a cluster) as the locks only work for a single application's threads.

## Warning

The integers generated have nanosecond resolution. Because the ID contains this time information, it could present an attack vector. If using these as public identifiers, know that the clients will know the exact time the ID was generated.

Most people should use UUID's (v4 also contains an embedded timestamp).

## Background

Go's `time` package has a [monotonic clock](https://github.com/golang/go/issues/12914#issuecomment-277335863) since Go 1.9. You can read the original  [Proposal here](https://github.com/golang/proposal/blob/master/design/12914-monotonic.md) and the [time package documentation](https://golang.org/pkg/time/#hdr-Monotonic_Clocks) on the subject.
