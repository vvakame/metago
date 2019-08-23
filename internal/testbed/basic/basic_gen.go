// Code generated by metago. DO NOT EDIT.

//+build !metago

// basic usage.

package basic

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Foo struct {
	ID        int64
	Name      string `json:"nickname"`
	CreatedAt time.Time
}

func (obj *Foo) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	var i int
	{
		// continue first!

		if i != 0 {
			buf.WriteString(",")
		}

		propertyName := "ID"
		if v := ""; v != "" {
			propertyName = strings.SplitN(v, ",", 2)[0]
		}

		buf.WriteString(strconv.Quote(propertyName))
		buf.WriteString(":")

		{
			b, err := json.Marshal(obj.ID)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		}

		i++
	}
	{

		if i != 0 {
			buf.WriteString(",")
		}

		propertyName := "Name"
		if v := "nickname"; v != "" {
			propertyName = strings.SplitN(v, ",", 2)[0]
		}

		buf.WriteString(strconv.Quote(propertyName))
		buf.WriteString(":")

		{
			b, err := json.Marshal(obj.Name)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		}

		i++
	}
	{

		if obj.CreatedAt.IsZero() {
			goto metagoGoto0

		}

		if i != 0 {
			buf.WriteString(",")
		}

		propertyName := "CreatedAt"
		if v := ""; v != "" {
			propertyName = strings.SplitN(v, ",", 2)[0]
		}

		buf.WriteString(strconv.Quote(propertyName))
		buf.WriteString(":")

		{
			b, err := obj.CreatedAt.MarshalJSON()
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		}

		i++
	}
metagoGoto0:
	;

	buf.WriteString("}")

	return buf.Bytes(), nil
}