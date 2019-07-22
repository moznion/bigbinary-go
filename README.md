bigbinary-go [![CircleCI](https://circleci.com/gh/moznion/bigbinary-go/tree/master.svg?style=svg)](https://circleci.com/gh/moznion/bigbinary-go/tree/master)
==

bigbinary provides some utility functions to convert a byte slice to a big integer.

Synopsis
--

```go
bytes := []byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb, 0xfa, 0xf9, 0xf8, 0xf7, 0xf6, 0xf5, 0xf4, 0xf3, 0xf2, 0xf1, 0xf0}
bigInt := ConvertWithBigEndian(bytes)
fmt.Printf("%s\n", bigInt.String()) // 340277133820332220657323652036036850160
```

```go
bytes := []byte{0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff}
bigInt := ConvertWithBigEndian(bytes)
fmt.Printf("%s\n", bigInt.String()) // 340277133820332220657323652036036850160
```

Author
--

moznion (<moznion@gmail.com>)

License
--

MIT

