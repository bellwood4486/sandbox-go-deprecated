package serialize

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func deepCopyJSON(src, dst interface{}) {
	buf, _ := json.Marshal(src)
	_ = json.Unmarshal(buf, dst)
}

func deepCopyGob(src, dst interface{}) {
	var buf bytes.Buffer
	_ = gob.NewEncoder(&buf).Encode(src)
	_ = gob.NewDecoder(&buf).Decode(dst)
}
