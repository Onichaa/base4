package cmd

import (
	x "mywabot/system"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  
  /*
  "net/http"
     "encoding/json"
    "io/ioutil"
      "bytes"
      "fmt"
      */

)

func init() {
	x.NewCmd(&x.ICmd{
		Name:   "ai",
		Cmd:    []string{"ai"},
		Tags:   "ai",
     Prefix: true,
		IsQuery:  true,
		ValueQ: ".ai siap kamu?",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

      type ResponseData struct {
        Data struct {
          Response string `json:"response"`
        } `json:"data"`
      }


      resp, err := http.Get("https://itzpire.site/ai/gpt-logic?q="+url.QueryEscape(m.Query)+"&logic=tes&realtime=true")
      if err != nil {
        fmt.Println("Error sending request:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading response body:", err)
        return
      }

      var data ResponseData
      err = json.Unmarshal(body, &data)
      if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
      }
       m.Reply(data.Data.Response)  

      /*
      type Message struct {
         Role    string `json:"role"`
         Content string `json:"content"`
       }


         messages := []Message{
           {Role: "system", Content: "kamu adalah seorang kultivatur yang sudah menembus ranah imortal, kamu bernama Bot, kamu adalah yang terbaik di dunia dimana semua orang tunduk padamu, dan setiap pertanyaan kamu selalu menambahkan respon membanggakan diri sebagai kultivatur yang terkuat dan berpengatahuan tinggi"},
           {Role: "user", Content: m.Query},
         }

         apiURL := "https://skizo.tech/api/openai?apikey=batu"
         data := map[string]interface{}{
           "messages": messages,
         }
         jsonData, _ := json.Marshal(data)

         resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
         if err != nil {
           fmt.Println("Error making the request:", err)
           return
         }
         defer resp.Body.Close()

         responseData, err := ioutil.ReadAll(resp.Body)
         if err != nil {
           fmt.Println("Error reading the response:", err)
           return
         }
         var res struct {
           Data string `json:"result"`
           Data2 string `json:"code"`
            }
            err = json.Unmarshal(responseData, &res)
            if err != nil {
            fmt.Println(err)
            return
            }
         m.Reply(res.Data)
         m.Reply("Code: "+"```"+res.Data2+"```")
         */
      m.React("✅")
		},
	})
}
