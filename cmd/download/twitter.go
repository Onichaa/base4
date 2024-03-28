package cmd

import (
  x "mywabot/system"
  
    "fmt"
    "net/http"
    "net/url"
     "encoding/json"
    "io/ioutil"
    "strconv"
    "strings"
    "regexp"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(tt|ttnowm|tiktoknowm|tiktok|tiktokmp3|ttmp3)",
    Cmd:    []string{"tiktok", "tiktoknowm", "tiktokmp3"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    ValueQ: ".tt https://vt.tiktok.com/ZSFk78cDv/",
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      if !strings.Contains(m.Query, "tiktok") {
        m.Reply("Itu bukan link tiktok")
      return
      }  

      regex := regexp.MustCompile(`(https?:\/\/[^\s]+)`)
       newLink := regex.FindStringSubmatch(m.Query) 

      

    type Data struct {
      Result struct {
        ProcessedTime float64 `json:"processed_time"`
        Data          struct {
          Region          string `json:"region"`
          Title           string `json:"title"`
          Duration        int    `json:"duration"`
          Play            string `json:"play"`
          WmPlay          string `json:"wmplay"`
          Music           string `json:"music"`
          MusicInfo       struct {
            Title  string `json:"title"`
            Play   string `json:"play"`
            Author string `json:"author"`
          } `json:"music_info"`
          PlayCount     int `json:"play_count"`
          CommentCount  int `json:"comment_count"`
          ShareCount    int `json:"share_count"`
          DownloadCount int `json:"download_count"`
          Author        struct {
            Nickname  string `json:"nickname"`
            Avatar    string `json:"avatar"`
          } `json:"author"`
          Images []string `json:"images"`
        } `json:"data"`
      } `json:"result"`
    }

        url := "https://aemt.me/download/tiktokslide?url="+ url.QueryEscape(newLink[0])

      response, err := http.Get(url)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer response.Body.Close()

      body, err := ioutil.ReadAll(response.Body)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      var tiktok Data
      err = json.Unmarshal(body, &tiktok)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
             teks := `*TIKTOK NO WATERMARK*
             
*𖦹 Author:* ` + tiktok.Result.Data.Author.Nickname + `
*𖦹 Region:* ` + tiktok.Result.Data.Region + `
*𖦹 Judul:* ` + tiktok.Result.Data.Title + `
*𖦹 Durasi:* ` + strconv.Itoa(tiktok.Result.Data.Duration) + `
*𖦹 Info Musik:*
  *- Judul:* ` + tiktok.Result.Data.MusicInfo.Title + `
  *- Author:* ` + tiktok.Result.Data.MusicInfo.Author + `
*𖦹 Jumlah Dilihat:* ` + strconv.Itoa(tiktok.Result.Data.PlayCount) + ` kali
*𖦹 Jumlah Dikomentar:* ` + strconv.Itoa(tiktok.Result.Data.CommentCount) + `
*𖦹 Jumlah Dishare:* ` + strconv.Itoa(tiktok.Result.Data.ShareCount) + ` kali
*𖦹 Jumlah Didownload:* ` + strconv.Itoa(tiktok.Result.Data.DownloadCount) + ` kali`


      if tiktok.Result.Data.Duration == 0 {
        for _, i := range tiktok.Result.Data.Images {
          
          sock.SendImage(m.From, i, teks, *m)

        }   
      } else if reg, _ := regexp.MatchString(`(tiktokmp3|ttmp3)`, m.Text); reg {    
      sock.SendAudio(m.From, tiktok.Result.Data.MusicInfo.Play, false, *m) 
        } else {
      sock.SendVideo(m.From,tiktok.Result.Data.Play, teks, *m)
    }
      
      
      m.React("✅")
    },
  })
}
