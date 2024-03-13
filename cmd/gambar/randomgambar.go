package cmd

import (
  x "mywabot/system"

  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "regexp"
  "math/rand"
  "time"
  "os"
  "os/exec"
  "bytes"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(art|asuna|aesthetic|akira|akiyama|ana|anjing|ayuzawa|boneka|boruto|bts|cecan|chiho|chitoge|cosplay|cogan|cosplayloli|cosplaysagiri|cyber|deidara|doraemon|eba|elaina|emilia|erza|exo|femdom|freefire|gamewallpaper|gremory|hekel|hestia|hijaber|hinata|husbu|inori|islamic|isuzu|itachi|itori|jahy|jeni|jiso|justina|kaga|kagura|kakasih|kaori|kartun|katakata|kaneki|kotori|kpop|kucing|kurumi|lisa|madara|megumin|mikasa|mikey|miku|minato|mobil|motor|mountain|naruto|nezuko|onepiece|pentol|pokemon|profil|pubg|randblackpink|rize|rose|ryujin|sagiri|sakura|sasuke|satanic|shina|shinka|shinomiya|shizuka|tatasurya|technology|tejina|toukachan|waifu|wallhp|wallml|wallmlnime|yotsuba|yuki|yulibocil|yumeko)",
    Cmd:    []string{"art", "asuna", "aesthetic", "akira", "akiyama", "ana", "anjing", "ayuzawa", "boneka", "boruto", "bts", "cecan", "chiho", "chitoge", "cosplay","cogan", "cosplayloli", "cosplaysagiri", "cyber", "deidara", "doraemon", "eba", "elaina", "emilia", "erza", "exo", "femdom", "freefire", "gamewallpaper", "gremory", "hekel", "hestia", "hijaber", "hinata", "husbu", "inori", "islamic", "isuzu", "itachi", "itori", "jahy", "jeni", "jiso", "justina", "kaga", "kagura", "kakasih", "kaori", "kartun", "katakata", "kaneki", "kotori", "kpop", "kucing", "kurumi", "lisa", "madara", "megumin", "mikasa", "mikey", "miku", "minato", "mobil", "motor", "mountain", "naruto", "nezuko", "onepiece", "pentol", "pokemon", "profil", "pubg", "randblackpink", "rize", "rose", "ryujin", "sagiri", "sakura", "sasuke", "satanic", "shina", "shinka", "shinomiya", "shizuka", "tatasurya", "technology", "tejina", "toukachan", "waifu", "wallhp", "wallml", "wallmlnime", "yotsuba", "yuki", "yulibocil", "yumeko"},
    Tags:   "gambar",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      var heyy string

      if reg, _ := regexp.MatchString(`(art)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/art.json"
        } else if reg, _ := regexp.MatchString(`(asuna)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/asuna.json"
        } else if reg, _ := regexp.MatchString(`(aesthetic)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/aesthetic.json"
        } else if reg, _ := regexp.MatchString(`(akira)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/akira.json"
        } else if reg, _ := regexp.MatchString(`(akiyama)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/akiyama.json"
        } else if reg, _ := regexp.MatchString(`(ana)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/ana.json"
        } else if reg, _ := regexp.MatchString(`(anjing)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/anjing.json"
        } else if reg, _ := regexp.MatchString(`(art)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/art.json"
        } else if reg, _ := regexp.MatchString(`(asuna)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/asuna.json"
        } else if reg, _ := regexp.MatchString(`(ayuzawa)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/ayuzawa.json"
        } else if reg, _ := regexp.MatchString(`(boneka)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/boneka.json"
        } else if reg, _ := regexp.MatchString(`(boruto)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/boruto.json"
        } else if reg, _ := regexp.MatchString(`(bts)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/bts.json"
        } else if reg, _ := regexp.MatchString(`(cecan)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/cecan.json"
        } else if reg, _ := regexp.MatchString(`(chiho)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/chiho.json"
        } else if reg, _ := regexp.MatchString(`(chitoge)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/chitoge.json"
        } else if reg, _ := regexp.MatchString(`(cogan)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/cogan.json"
        } else if reg, _ := regexp.MatchString(`(cosplay)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/cosplay.json"
        } else if reg, _ := regexp.MatchString(`(cosplayloli)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/cosplayloli.json" 
        } else if reg, _ := regexp.MatchString(`(cosplaysagiri)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/cosplaysagiri.json" 
        } else if reg, _ := regexp.MatchString(`(cyber)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/cyber.json" 
        } else if reg, _ := regexp.MatchString(`(deidara)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/deidara.json" 
        } else if reg, _ := regexp.MatchString(`(doraemon)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/doraemon.json" 
        } else if reg, _ := regexp.MatchString(`(eba)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/eba.json" 
        } else if reg, _ := regexp.MatchString(`(elaina)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/elaina.json" 
        } else if reg, _ := regexp.MatchString(`(emilia)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/emilia.json" 
        } else if reg, _ := regexp.MatchString(`(erza)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/erza.json"
        } else if reg, _ := regexp.MatchString(`(exo)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/exo.json" 
        } else if reg, _ := regexp.MatchString(`(femdom)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/femdom.json"
        } else if reg, _ := regexp.MatchString(`(freefire)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/freefire.json"
        } else if reg, _ := regexp.MatchString(`(gamewallpaper)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/gamewallpaper.json"
        } else if reg, _ := regexp.MatchString(`(gremory)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/gremory.json"
        } else if reg, _ := regexp.MatchString(`(hekel)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/hekel.json"
        } else if reg, _ := regexp.MatchString(`(hestia)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/hestia.json"
        } else if reg, _ := regexp.MatchString(`(hijaber)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/hijaber.json"
        } else if reg, _ := regexp.MatchString(`(hinata)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/hinata.json"
        } else if reg, _ := regexp.MatchString(`(husbu)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/husbu.json"
        } else if reg, _ := regexp.MatchString(`(inori)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/inori.json"
        } else if reg, _ := regexp.MatchString(`(islamic)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/islamic.json"
        } else if reg, _ := regexp.MatchString(`(isuzu)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/isuzu.json"
        } else if reg, _ := regexp.MatchString(`(itachi)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/itachi.json"
        } else if reg, _ := regexp.MatchString(`(itori)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/itori.json"
        } else if reg, _ := regexp.MatchString(`(jahy)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/jahy.json"
        } else if reg, _ := regexp.MatchString(`(jeni)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/jeni.json"
        } else if reg, _ := regexp.MatchString(`(jiso)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/jiso.json"
        } else if reg, _ := regexp.MatchString(`(justina)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/justina.json"
        } else if reg, _ := regexp.MatchString(`(kaga)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kaga.json"
        } else if reg, _ := regexp.MatchString(`(kagura)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kagura.json"
        } else if reg, _ := regexp.MatchString(`(kakasih)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kakasih.json"
        } else if reg, _ := regexp.MatchString(`(kaori)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kaori.json"
        } else if reg, _ := regexp.MatchString(`(kartun)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kartun.json"
        } else if reg, _ := regexp.MatchString(`(katakata)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/katakata.json"
        } else if reg, _ := regexp.MatchString(`(kaneki)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/keneki.json"
        } else if reg, _ := regexp.MatchString(`(kotori)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kotori.json"
        } else if reg, _ := regexp.MatchString(`(kpop)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kpop.json"
        } else if reg, _ := regexp.MatchString(`(kucing)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kucing.json"
        } else if reg, _ := regexp.MatchString(`(kurumi)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/kurumi.json"
        } else if reg, _ := regexp.MatchString(`(lisa)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/lisa.json"
        } else if reg, _ := regexp.MatchString(`(madara)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/madara.json"
        } else if reg, _ := regexp.MatchString(`(megumin)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/megumin.json"
        } else if reg, _ := regexp.MatchString(`(mikasa)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/mikasa.json"
        } else if reg, _ := regexp.MatchString(`(mikey)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/mikey.json"
        } else if reg, _ := regexp.MatchString(`(miku)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/miku.json"
        } else if reg, _ := regexp.MatchString(`(minato)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/minato.json"
        } else if reg, _ := regexp.MatchString(`(mobil)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/mobil.json"
        } else if reg, _ := regexp.MatchString(`(motor)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/motor.json"
        } else if reg, _ := regexp.MatchString(`(mountain)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/mountain.json"
        } else if reg, _ := regexp.MatchString(`(naruto)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/naruto.json"
        } else if reg, _ := regexp.MatchString(`(nezuko)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/nezuko.json"
        } else if reg, _ := regexp.MatchString(`(onepiece)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/onepiece.json"
        } else if reg, _ := regexp.MatchString(`(orgy)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/orgy.json"
        } else if reg, _ := regexp.MatchString(`(panties)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/panties.json"
        } else if reg, _ := regexp.MatchString(`(pentol)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/pentol.json"
        } else if reg, _ := regexp.MatchString(`(pokemon)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/pokemon.json"
        } else if reg, _ := regexp.MatchString(`(profil)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/profil.json"
        } else if reg, _ := regexp.MatchString(`(pubg)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/pubg.json"
        } else if reg, _ := regexp.MatchString(`(randblackpink)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/randblackpink.json"
        } else if reg, _ := regexp.MatchString(`(rize)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/rize.json"
        } else if reg, _ := regexp.MatchString(`(rose)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/rose.json"
        } else if reg, _ := regexp.MatchString(`(ryujin)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/ryujin.json"
        } else if reg, _ := regexp.MatchString(`(sagiri)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/sagiri.json"
        } else if reg, _ := regexp.MatchString(`(sakura)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/sakura.json"
        } else if reg, _ := regexp.MatchString(`(sasuke)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/sasuke.json"
        } else if reg, _ := regexp.MatchString(`(satanic)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/satanic.json"
        } else if reg, _ := regexp.MatchString(`(shina)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/shina.json"
        } else if reg, _ := regexp.MatchString(`(shinka)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/shinka.json"
        } else if reg, _ := regexp.MatchString(`(shinomiya)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/shinomiya.json"
        } else if reg, _ := regexp.MatchString(`(shizuka)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/shizuka.json"
        } else if reg, _ := regexp.MatchString(`(tatasurya)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/tatasurya.json"
        } else if reg, _ := regexp.MatchString(`(technology)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/technology.json"
        } else if reg, _ := regexp.MatchString(`(tejina)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/tejina.json"
        } else if reg, _ := regexp.MatchString(`(thighs)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/thighs.json"
        } else if reg, _ := regexp.MatchString(`(toukachan)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/toukachan.json"
        } else if reg, _ := regexp.MatchString(`(waifu)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/waifu.json"
        } else if reg, _ := regexp.MatchString(`(wallhp)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/wallhp.json"
        } else if reg, _ := regexp.MatchString(`(wallml)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/wallml.json"
        } else if reg, _ := regexp.MatchString(`(wallmlnime)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/wallnime.json"
        } else if reg, _ := regexp.MatchString(`(yotsuba)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/yotsuba.json"
        } else if reg, _ := regexp.MatchString(`(yuki)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/yuki.json"
        } else if reg, _ := regexp.MatchString(`(yulibocil)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/yulibocil.json"
        } else if reg, _ := regexp.MatchString(`(yumeko)`, m.Text); reg {
        heyy = "https://raw.githubusercontent.com/KirBotz/nyenyee/master/yumeko.json"
        }

      resp, err := http.Get(heyy)
      if err != nil {
        fmt.Println("Error downloading json:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading json response body:", err)
        return
      }

      var urls []string
      if err := json.Unmarshal(body, &urls); err != nil {
        fmt.Println("Error unmarshalling json:", err)
        return
      }

      rand.Seed(time.Now().UnixNano())
      randomIndex := rand.Intn(len(urls))
      randomUrl := urls[randomIndex]

      byte, err := x.ToByte(randomUrl)
      if err != nil {
         fmt.Println("Erroor:", err)
        return
      }
      
      conjp := "./tmp/" + m.ID + ".jpeg"
      conwb := "./tmp/" + m.ID + ".jpg"
      errr := os.WriteFile(conwb, byte, 0644)
      if errr != nil {
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
        return
      }

      url, err := x.Upload(conjp)
      if err != nil {
          fmt.Println("Error:", err)
          return
      }
      
      os.Remove(conwb)
      os.Remove(conjp)
      sock.SendImage(m.From, url, "", *m)
     
      m.React("✅")
    },
  })
}
