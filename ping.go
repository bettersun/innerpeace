package innerpeace

import (
	"log"

	"github.com/bettersun/innerpeace/common"
	"github.com/bettersun/moist"
)

func Ping() (interface{}, error) {
	var res common.Response

	m := make(map[string]interface{})
	m["hello"] = "world"

	res.Code = 100
	res.Message = "Ping"
	res.Data = m

	mm, err := moist.StructToMap(res)
	if err != nil {
		log.Println(err)
	}
	data, err := moist.ToIfKeyMap(mm)
	if err != nil {
		log.Println(err)
	}

	return data, err
}
