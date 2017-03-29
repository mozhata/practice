package main

import (
	"encoding/json"
	"fmt"
)

var P func(...interface{}) (int, error) = fmt.Println

type Response1 struct {
	Page   int
	Fruits []string
	Boo    bool
	Boo2   bool
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		Boo:    true,
	}
	P(res1D)
	res1B, err := json.Marshal(res1D)
	fmt.Println(string(res1B), err)
	v := Response1{}
	err = json.Unmarshal(res1B, &v)
	P("unmarshal: ", v, "err: ", err)
	{
		/*	res2D := &Response2{
				Page:   1,
				Fruits: []string{"apple", "peach", "pear"},
			}
			res2B, _ := json.Marshal(res2D)
			fmt.Println(string(res2B))

			byt := []byte(`{"num":6.13,"strs":["abs","b"]}`)
			var dat map[string]interface{}
			if err := json.Unmarshal(byt, &dat); err != nil {
				panic(err)
			}
			fmt.Println(byt)
			fmt.Println(dat)
			// num := dat["num"].(float64)
			// fmt.Println(num, dat["num"])
			// strs := dat["strs"].([]interface{})
			// str1 := strs[0].(string)
			// fmt.Println(str1, strs[0])
			str := `{"page": 1, "fruits": ["apple", "peach"]}`
			res := &Response2{}
			json.Unmarshal([]byte(str), res)
			fmt.Println(res, res.Fruits[0])

			enc := json.NewEncoder((os.Stdout))
			d := map[string]int{"apple": 5, "lettuce": 8}
			enc.Encode(d)
		*/
	}
}
