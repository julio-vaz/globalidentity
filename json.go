package globalidentity

import (
	"bytes"
	"encoding/json"
	"io"
)

func toJson(s interface{}) (io.Reader, error) {
	b, err := json.Marshal(s)
	return bytes.NewBuffer(b), err
}

func fromJson(s interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	return decoder.Decode(s)
}
