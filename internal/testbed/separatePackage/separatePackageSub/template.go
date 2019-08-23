package separatePackageSub

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/vvakame/metago"
)

func MarshalJSONTemplate(mv metago.Value) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	var i int
	for _, mf := range mv.Fields() {
		if mf.Value().(time.Time).IsZero() {
			continue
		}

		if i != 0 {
			buf.WriteString(",")
		}

		propertyName := mf.Name()
		if v := mf.StructTagGet("json"); v != "" {
			propertyName = strings.SplitN(v, ",", 2)[0]
		}

		buf.WriteString(`"`)
		buf.WriteString(propertyName)
		buf.WriteString(`":`)

		switch v := mf.Value().(type) {
		case int64:
			buf.Write([]byte(strconv.FormatInt(v, 10)))
		case string:
			buf.Write([]byte(strconv.Quote(v)))
		case time.Time:
			b, err := v.MarshalJSON()
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		default:
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		}

		i++
	}

	buf.WriteString("}")

	return buf.Bytes(), nil
}
