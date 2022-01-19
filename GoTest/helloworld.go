package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	// "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/httplib"
)

type Data []struct {

UserID string `json:"userId"`
ID string `json:"id"`
Title string `json:"title"`
Body string `json:"body"`

}

type Data1 []struct {
ID string `json:"id"`
Url string `json:"url"`
Width string `json:"width"`
Height string `json:"height"`
Mime_type string `json:"mime_type"`
Entities string `json:"entities"`
Breeds string `json:"breeds"`
Animals string `json:"animals"`
Categories string `json:"categories"`
}

func main() {

req := httplib.Get("https://api.thecatapi.com/v1/images/search?limit=6&page=3&breed=Abyssinian")
req.Header("x-api-key", `31cf5af0-bb4d-4275-971c-0e161cbdfa0b`)
// req.Param("username", "astaxie")
// req.Param("password", "123456")
resp, err := req.Response()
if err != nil {
log.Fatal(err)
}

body, _ := ioutil.ReadAll(resp.Body)

var f Data1

err = json.Unmarshal(body, &f)

//fmt.Printf("%+v", f[0].Body)

for i := 0; i < 6; i++ {

fmt.Println(f[i].Url)
}
// fmt.Printf("Body : %s", body)
    // web.Run()
}

