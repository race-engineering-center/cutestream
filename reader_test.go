package cutestream

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"reflect"
	"strconv"
	"testing"
)

const TESTDIR_ENV = "CUTESTREAM_TEST_DIR"

// helper functions -----------------------------------------------------------

func readFile(fileName string) ([]byte, error) {
	testFileName := os.Getenv(TESTDIR_ENV) + "/" + fileName
	file, err := os.Open(testFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func checkInt[T int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](result T, expected string, t *testing.T) {
	var expectedValue T
	switch reflect.ValueOf(result).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(expected, 10, int(reflect.TypeOf(result).Size())*8)
		assert.Nil(t, err)
		expectedValue = T(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := strconv.ParseUint(expected, 10, int(reflect.TypeOf(result).Size())*8)
		assert.Nil(t, err)
		expectedValue = T(v)
	}

	assert.Equal(t, expectedValue, result)
}

func checkFloat[T float32 | float64](result T, expected string, depth int, t *testing.T) {
	expectedValue, err := strconv.ParseFloat(expected, depth)
	assert.Nil(t, err)
	assert.Equal(t, T(expectedValue), result)
}

// tests ----------------------------------------------------------------------

func TestEnvSetup(t *testing.T) {
	testFolder := os.Getenv("CUTESTREAM_TEST_DIR")
	assert.Greater(t, len(testFolder), 0)
	assert.DirExists(t, testFolder)
}

func TestIntegerNumbers(t *testing.T) {
	j, err := readFile("generated_int.json")
	assert.Nil(t, err)

	type IntegerNumberData struct {
		Serialized string `json:"serialized"`
		Value      string `json:"value"`
	}

	type IntegerNumbersTestData map[string]map[string][]IntegerNumberData

	var data IntegerNumbersTestData
	err = json.Unmarshal(j, &data)
	assert.Nil(t, err)

	for version, versionData := range data {
		v, err := strconv.ParseInt(version, 10, 32)
		assert.Nil(t, err)
		for dataType, numbers := range versionData {
			for _, n := range numbers {
				b, err := base64.StdEncoding.DecodeString(n.Serialized)
				assert.Nil(t, err)

				reader, err := NewReaderWithVersion(bytes.NewReader(b), int(v))
				assert.Nil(t, err)

				switch dataType {
				case "int8":
					readValue, err := reader.ReadInt8()
					assert.Nil(t, err)
					checkInt[int8](readValue, n.Value, t)
				case "uint8":
					readValue, err := reader.ReadUint8()
					assert.Nil(t, err)
					checkInt[uint8](readValue, n.Value, t)
				case "int16":
					readValue, err := reader.ReadInt16()
					assert.Nil(t, err)
					checkInt[int16](readValue, n.Value, t)
				case "uint16":
					readValue, err := reader.ReadUint16()
					assert.Nil(t, err)
					checkInt[uint16](readValue, n.Value, t)
				case "int32":
					readValue, err := reader.ReadInt32()
					assert.Nil(t, err)
					checkInt[int32](readValue, n.Value, t)
				case "uint32":
					readValue, err := reader.ReadUint32()
					assert.Nil(t, err)
					checkInt[uint32](readValue, n.Value, t)
				case "int64":
					readValue, err := reader.ReadInt64()
					assert.Nil(t, err)
					checkInt[int64](readValue, n.Value, t)
				case "uint64":
					readValue, err := reader.ReadUint64()
					assert.Nil(t, err)
					checkInt[uint64](readValue, n.Value, t)
				default:
					t.Fatalf("Unsupported data type: %s", dataType)
				}
			}
		}
	}
}

func TestFloats(t *testing.T) {
	j, err := readFile("generated_float.json")
	assert.Nil(t, err)

	type FloatData struct {
		Serialized string `json:"serialized"`
		Value      string `json:"value"`
	}

	type FloatsTestData map[string]map[string][]FloatData

	var data FloatsTestData
	err = json.Unmarshal(j, &data)
	assert.Nil(t, err)

	for version, versionData := range data {
		v, err := strconv.ParseInt(version, 10, 32)
		assert.Nil(t, err)
		for dataType, floats := range versionData {
			for _, f := range floats {
				b, err := base64.StdEncoding.DecodeString(f.Serialized)
				assert.Nil(t, err)

				reader, err := NewReaderWithVersion(bytes.NewReader(b), int(v))
				assert.Nil(t, err)
				// float json contains separate values for
				// single- and double point precision
				// that's what _d and _s are for
				switch dataType {
				case "float_d":
					reader.DoublePrecision = true
					readValue, err := reader.ReadFloat()
					assert.Nil(t, err)
					checkFloat[float32](readValue, f.Value, 32, t)
				case "double_d":
					reader.DoublePrecision = true
					readValue, err := reader.ReadDouble()
					assert.Nil(t, err)
					checkFloat[float64](readValue, f.Value, 64, t)
				case "float_s":
					reader.DoublePrecision = false
					readValue, err := reader.ReadFloat()
					assert.Nil(t, err)
					checkFloat[float32](readValue, f.Value, 32, t)
				case "double_s":
					reader.DoublePrecision = false
					readValue, err := reader.ReadDouble()
					assert.Nil(t, err)
					checkFloat[float64](readValue, f.Value, 32, t)
				default:
					t.Fatalf("Unsupported data type: %s", dataType)
				}
			}
		}
	}
}

func TestDate(t *testing.T) {
	j, err := readFile("generated_date.json")
	assert.Nil(t, err)

	type Data struct {
		Version int `json:"version"`
		Dates   []struct {
			Serialized string `json:"serialized"`
			Value      struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"value"`
		} `json:"date"`
	}

	var data map[string]Data
	err = json.Unmarshal(j, &data)
	assert.Nil(t, err)

	for version, versionData := range data {
		v, err := strconv.ParseInt(version, 10, 32)
		assert.Nil(t, err)
		assert.Greater(t, len(versionData.Dates), 0)
		for _, d := range versionData.Dates {
			b, err := base64.StdEncoding.DecodeString(d.Serialized)
			assert.Nil(t, err)
			reader, err := NewReaderWithVersion(bytes.NewReader(b), int(v))
			assert.Nil(t, err)
			date, err := reader.ReadQDate()
			assert.Nil(t, err)
			assert.Equal(t, d.Value.Day, date.Day())
			assert.Equal(t, d.Value.Month, int(date.Month()))
			assert.Equal(t, d.Value.Year, date.Year())
		}
	}
}

func TestTime(t *testing.T) {
	j, err := readFile("generated_time.json")
	assert.Nil(t, err)

	type Data struct {
		Time []struct {
			Serialized string `json:"serialized"`
			Value      struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
				Ms     int `json:"ms"`
				Sec    int `json:"sec"`
			} `json:"value"`
		} `json:"time"`
	}

	var data map[string]Data
	err = json.Unmarshal(j, &data)
	assert.Nil(t, err)

	for version, versionData := range data {
		v, err := strconv.ParseInt(version, 10, 32)
		assert.Nil(t, err)
		assert.Greater(t, len(versionData.Time), 0)
		for _, tm := range versionData.Time {
			b, err := base64.StdEncoding.DecodeString(tm.Serialized)
			assert.Nil(t, err)
			reader, err := NewReaderWithVersion(bytes.NewReader(b), int(v))
			assert.Nil(t, err)
			time, err := reader.ReadQTime()
			assert.Nil(t, err)
			assert.Equal(t, tm.Value.Hour, int(time.Hours()))
			assert.Equal(t, tm.Value.Minute, int(time.Minutes())%60)
			assert.Equal(t, tm.Value.Sec, int(time.Seconds())%60)
			assert.Equal(t, tm.Value.Ms, int(time.Nanoseconds()/1000000)%1000)
		}
	}
}

func TestDateTime(t *testing.T) {
	j, err := readFile("generated_datetime.json")
	assert.Nil(t, err)

	type Data struct {
		Version   int `json:"version"`
		Datetimes []struct {
			Serialized string `json:"serialized"`
			Value      struct {
				Day    int `json:"day"`
				Month  int `json:"month"`
				Year   int `json:"year"`
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
				Sec    int `json:"sec"`
				Ms     int `json:"ms"`
			} `json:"value"`
		} `json:"datetime"`
	}

	var data map[string]Data
	err = json.Unmarshal(j, &data)
	assert.Nil(t, err)

	for version, versionData := range data {
		v, err := strconv.ParseInt(version, 10, 32)
		assert.Nil(t, err)
		assert.Greater(t, len(versionData.Datetimes), 0)
		for _, d := range versionData.Datetimes {
			b, err := base64.StdEncoding.DecodeString(d.Serialized)
			assert.Nil(t, err)
			reader, err := NewReaderWithVersion(bytes.NewReader(b), int(v))
			assert.Nil(t, err)
			datetime, err := reader.ReadQDateTime()
			assert.Nil(t, err)
			assert.Equal(t, d.Value.Day, datetime.Day())
			assert.Equal(t, d.Value.Month, int(datetime.Month()))
			assert.Equal(t, d.Value.Year, datetime.Year())
			assert.Equal(t, d.Value.Hour, datetime.Hour())
			assert.Equal(t, d.Value.Minute, datetime.Minute())
			assert.Equal(t, d.Value.Sec, datetime.Second())
			assert.Equal(t, d.Value.Ms, datetime.Nanosecond()/1000000)
		}
	}
}

func TestUuid(t *testing.T) {
	j, err := readFile("generated_uuid.json")
	assert.Nil(t, err)

	type QUuidData struct {
		Uuid []struct {
			Serialized string `json:"serialized"`
			Value      string `json:"value"`
		} `json:"uuid"`
	}

	type QUuidTestData map[string]QUuidData

	var data QUuidTestData
	err = json.Unmarshal(j, &data)
	assert.Nil(t, err)
	for version, versionData := range data {
		v, err := strconv.ParseInt(version, 10, 32)
		assert.Greater(t, len(versionData.Uuid), 0)
		assert.Nil(t, err)
		for _, u := range versionData.Uuid {
			b, err := base64.StdEncoding.DecodeString(u.Serialized)
			assert.Nil(t, err)
			reader, err := NewReaderWithVersion(bytes.NewReader(b), int(v))
			assert.Nil(t, err)
			uuid, err := reader.ReadQUuid()
			assert.Nil(t, err)
			assert.Equal(t, u.Value, uuid)
		}
	}
}
