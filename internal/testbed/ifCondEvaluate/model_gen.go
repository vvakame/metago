// Code generated by metago. DO NOT EDIT.

//+build !metago

// evaluate `if v, ok := mf.Value().(Foo); [ok] { ... }` case.

package ifCondEvaluate

import (
	"fmt"
	"time"
)

type Foo struct {
	ID        int64
	Name      string `json:"nickname"`
	CreatedAt time.Time
}

func (obj *Foo) MarshalJSON() ([]byte, error) {

	{

		{
			fmt.Println("a", obj.ID)
		}

		{
			fmt.Println("b", obj.ID)
		}

		{
			fmt.Println("a", obj.ID)
		}

	}
	{

		{
			fmt.Println("b", obj.Name)
		}

		{
			fmt.Println("a", obj.Name)
		}

		{
			fmt.Println("b", obj.Name)
		}
	}
	{

		{
			fmt.Println("b", obj.CreatedAt)
		}

		{
			fmt.Println("a", obj.CreatedAt)
		}

		{
			fmt.Println("b", obj.CreatedAt)
		}
	}

	panic("👀")
}
