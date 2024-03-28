package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "gemini",
    Cmd:    []string{"gemini"},
    Tags:   "ai",
    IsQuery:  true,
    Prefix: true,
    ValueQ: `Contoh penggunaan:
    
1. .gemini siap kamu?
2. Kirim/Reply gambar dengan caption .gemini gambar apakah itu?`,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")

    type Result struct {
      Status  string `json:"status"`
      Author  string `json:"author"`
      Code    int    `json:"code"`
      Result  string `json:"result"`
    }

      // quoted image
      if m.IsQuotedImage {
        conjp1 := "./tmp/" + m.ID + ".jpg"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
        err := os.WriteFile(conjp1, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        Url, err := x.Upload(conjp1)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        
      resp, err := http.Get("https://itzpire.site/ai/gemini-ai?q="+url.QueryEscape(m.Query)+"&url="+Url)
      if err != nil {
        fmt.Println("Error making request:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading response body:", err)
        return
      }

      var result Result
      err = json.Unmarshal(body, &result)
      if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
      }

      m.Reply(result.Result)
        os.Remove(conjp1)
        } else {

    resp, err := http.Get("https://itzpire.site/ai/gemini-ai?q="+url.QueryEscape(m.Query))
    if err != nil {
      fmt.Println("Error making request:", err)
      return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("Error reading response body:", err)
      return
    }

    var result Result
    err = json.Unmarshal(body, &result)
    if err != nil {
      fmt.Println("Error unmarshalling JSON:", err)
      return
    }

    m.Reply(result.Result)

      }


      // from image
      if m.IsImage {
        conjp1 := "./tmp/" + m.ID + ".jpg"
        byte, _ := sock.WA.Download(m.Media)
        err := os.WriteFile(conjp1, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved jpg")
          return
        }

        Url, err := x.Upload(conjp1)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        resp, err := http.Get("https://itzpire.site/ai/gemini-ai?q="+url.QueryEscape(m.Query)+"&url="+Url)
        if err != nil {
          fmt.Println("Error making request:", err)
          return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Error reading response body:", err)
          return
        }

        var result Result
        err = json.Unmarshal(body, &result)
        if err != nil {
          fmt.Println("Error unmarshalling JSON:", err)
          return
        }

        m.Reply(result.Result)
          os.Remove(conjp1)
          } else {

        resp, err := http.Get("https://itzpire.site/ai/gemini-ai?q="+url.QueryEscape(m.Query))
        if err != nil {
          fmt.Println("Error making request:", err)
          return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Error reading response body:", err)
          return
        }

        var result Result
        err = json.Unmarshal(body, &result)
        if err != nil {
          fmt.Println("Error unmarshalling JSON:", err)
          return
        }

        m.Reply(result.Result)
        
          }
      
      m.React("✅")
    },
  })
}

