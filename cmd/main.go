// package main

// import "github.com/lzy/util"

// func main() {
// 	util.Util()
// }

package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "https://jx.2.akkcloud1.com/user/checkin"
   method := "POST"

   payload := strings.NewReader("")

   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Cookie", "uid=84480; email=1012938458%40qq.com; PHPSESSID=n5aun9b42as9m1btcl8f6c013l; key=1ef1600db8249c178068483ab1dc4682a330bb63c54b6; ip=afb13377a4f0a2c89e2aecb082454756; expire_in=1681139823; pop=yes; mtauth=5e2ceba40a7bb123052965a882244271")
   req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
   req.Header.Add("Accept", "*/*")
   req.Header.Add("Host", "jx.2.akkcloud1.com")
   req.Header.Add("Connection", "keep-alive")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(strings(body))
}