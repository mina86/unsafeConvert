# Unsafe string↔[]byte conversion library

The library functions for unsafely convert between a string and
a slice of bytes.  You probably shouldn’t use it unless you need to
squeeze extra performance from your performance-critical code path.

See https://mina86.com/2017/golang-string-and-bytes/ for some more info.
