package db

import (
	"encoding/json"
	"strconv"
)

// marshal converts a value into its storage representation.
// All values are expected to be marshallable so errors will panic.
func marshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	assert(err == nil, "marshal error: %s", err)
	return b
}

// unmarshal converts to a value from its storage representation.
// All data is expected to be unmarshallable so errors will panic.
func unmarshal(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	assert(err == nil, "unmarshal error: %s", err)
}

// itob converts an integer into a []byte representation.
func itob(i int) []byte {
	return []byte(strconv.Itoa(i))
}

// btoi converts to integer from a []byte representation.
func btoi(b []byte) int {
	i, err := strconv.Atoi(string(b))
	assert(err == nil, "btoi error: %s", err)
	return i
}
