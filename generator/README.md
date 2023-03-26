# QDataStream data generator

This is a tool to generate JSON-formatted representation of a `QDataStream`-serialized data. 
Generated JSONs can be used for testing purposes. 

## Usage

- Set `CUTESTREAM_TEST_DIR` to point to the specified you want to output test files to
- Build this tool
- Run
- Generated JSON files will be placed in the specified folder

## Data types 

Currently, the tool generated the following data types:

- All integer types (`int8_t`, `int16_t`, `int32_t`, `int64_t` and the corresponding unsigned versions)
- Float types (`float` and `double`, with both single-precision and double-precision. 
Refer to https://doc.qt.io/qt-6/qdatastream.html#FloatingPointPrecision-enum for details)
- `QUuid`
- `QData`, `QTime`, `QDateTime`

