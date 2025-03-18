package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string 	`json:"coursename"`
	Price    int
	Platform string 	`json:"website"`
	Password string		`json:"-"`
	Tags     []string 	`json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "lco.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "lco.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "lco.com", "hit123", nil},
	}
	// finalJson,err := json.Marshal(lcoCourses)
	finalJson,err := json.MarshalIndent(lcoCourses, "","\t")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDatafromWeb := []byte(`
		{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "lco.com",
                "tags": ["web-dev","js"]
        }
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDatafromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDatafromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else{
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n",k,v,v)
	}

}