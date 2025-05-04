# cutestream 

A Go library for reading files written by Qt QDataStream

Based on pgaskin's gist https://gist.github.com/pgaskin/a41a61ffe6d70567a11dc481020c5290

## Supported data types

- All flavours of `int`: `int8_t`, `int16_t`, `int32_t`, `int64_t`, `uint8_t`, `uint16_t`, `uint32_t`, `uint64_t`
- Floating-points: `float`, `double`
- `QDate`, `QTime`, `QDateTime`
- `QUuid`

### Supported but pending tests

- `bool`
- `QString`
- `QChar`
- `std::string`
- `QVariantMap`
- `QVariantHash`
- `QVariantList`
- `QUrl`

### Supported `QDataStream` versions

- Version 19 (Qt 5.13) 
- Version 20 (Qt 6.0)

Note that currently Qt 6.1 - 6.4 also use `QDataStream` Version 20, so they are supported as well. 
Earlier (pre-19) versions will probably also work for most of the types, but it's not guaranteed.

Refer to https://doc.qt.io/qt-6/qdatastream.html#Version-enum for details on `QDataStream` versioning.

## Testing

- Add path tp folder with test data (by default `test` folder in this project root) to `CUTESTREAM_TEST_DIR`
  environment variable
- Run `go test`

Test data is generated using bundled generator app (see `generator` folder in the project root). 
You can also use it to regenerate test data on a fly.

## License

This code is distributed under MIT license

Copyright (c) 2023-2025 Dmitriy Linev

## TODO

- [ ] Usage examples
- [ ] Publish to pkg.go.dev
- [ ] More data types and tests