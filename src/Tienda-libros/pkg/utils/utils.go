package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}

	}

}
//NOTE: .Unmarshal() sirve para convertir un objeto JSON en un objeto Go 
