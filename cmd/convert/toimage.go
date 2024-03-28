package cmd

import (
  "fmt"
  x "mywabot/system"
  "os"
  "os/exec"
  "bytes"
  "regexp"
  "io/ioutil"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name: "(toimage|toimg|tovideo)",
    Cmd:  []string{"toimage", "tovideo"},
    Tags: "convert",
    IsMedia: true,
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
      m.React("⏱️")


      if m.IsQuotedSticker {

        if reg, _ := regexp.MatchString(`(toimage|toimg)`, m.Text); reg {    
        conjp := "./tmp/" + m.ID + ".jpg"
        conwb := "./tmp/" + m.ID + ".webp"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        err := os.WriteFile(conwb, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved webp")
          return
        }

        cmd := exec.Command("ffmpeg", "-i", conwb, conjp)
        var out bytes.Buffer
        var stderr bytes.Buffer
        cmd.Stdout = &out
        cmd.Stderr = &stderr
        err = cmd.Run()

        // Check error
        if err != nil {
           fmt.Println("Error:", err)
          //m.Reply("Jangan reply stiker yang bergerak, harus nya ketik .tovideo untuk stiker yang bergerak menjadi video")
          return
        }

        url, err := x.Upload(conjp)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        sock.SendImage(m.From, url, "", *m)

        os.Remove(conwb)
        os.Remove(conjp)

        } else {
        conwb := "./tmp/" + m.ID + ".webp"
        outMp4 := "./tmp/" + m.ID + ".mp4"
        outGif := "./tmp/" + m.ID + ".gif"
        byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        err := ioutil.WriteFile(conwb, byte, 0644)
        if err != nil {
          fmt.Println("Failed saved webp")
          return
        }
       
        cmd := exec.Command("convert", conwb, outGif)
         err = cmd.Run(); 
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        cmd = exec.Command("ffmpeg", "-i", outGif, "-vf", "crop=trunc(iw/2)*2:trunc(ih/2)*2", "-b:v", "0", "-crf", "25", "-f", "mp4", "-vcodec", "libx264", "-pix_fmt", "yuv420p", outMp4)
         err = cmd.Run(); 
          if err != nil {
            fmt.Println("Error:", err)
            return
        }


          url, err := x.Upload(outMp4)
          if err != nil {
              fmt.Println("Error:", err)
              return
          }
          sock.SendVideo(m.From, url, "Converted sticker to video", *m)
          os.Remove(outGif)
          os.Remove(outMp4)
          os.Remove(conwb)
        }
        m.React("✅")
      }
    },
  })
}
