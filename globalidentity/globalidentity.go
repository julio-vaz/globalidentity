package Globalidentity

import (
	"github.com/levigross/grequests"
	"fmt"
	"encoding/json"
)

type GlobalIdentityResponse struct{
	args string
	headers map[string] string
	origin string
	url string
}

func GlobalIdentity() bool{
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	var respMap GlobalIdentityResponse{}
	if err != nil{
		fmt.Println("Deu merda")
	}
	json.Unmarshal([]byte(resp.String()), &respMap)
	fmt.Println(respMap)
	return true
}