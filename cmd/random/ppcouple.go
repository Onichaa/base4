package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "math/rand"
  "net/http"
  "time"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:    "(ppcp|ppcouple|couple)",
    Cmd:     []string{"ppcouple"},
    Tags:    "random",
    Prefix:  true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      res := "https://raw.githubusercontent.com/Dksitepu/data/main/Image/ppcouple.json"

      resp, err := http.Get(res)
      if err != nil {
        fmt.Println(err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println(err)
        return
      }

      var pp []struct {
        Male   string `json:"male"`
        Female string `json:"female"`
      }
      err = json.Unmarshal(body, &pp)
      if err != nil {
        fmt.Println(err)
        return
      }

      rand.Seed(time.Now().UnixNano())
      index := rand.Intn(len(pp))

      sock.SendImage(m.From, pp[index].Male, "Cowo", *m)
      sock.SendImage(m.From, pp[index].Female, "Cewe", *m)
       m.React("✅")
    },
  })
}
