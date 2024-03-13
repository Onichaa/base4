package cmd

import (
  x "mywabot/system"

  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:    "(dellssampah|delsampah)",
    Cmd:     []string{"delsampah"},
    Tags:    "owner",
    Prefix:  true,
    Exec: func(client *x.Nc, m *x.IMsg) {
      m.React("⏱️")

    files, err := ioutil.ReadDir("./tmp")
    if err != nil {
      fmt.Println("Unable to scan directory:", err)
      return
    }

    filteredFiles := make([]string, 0)
    for _, file := range files {
      name := file.Name()
      if strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") || strings.HasSuffix(name, ".webp") || strings.HasSuffix(name, ".png") {
        filteredFiles = append(filteredFiles, name)
      }
    }

    for _, file := range filteredFiles {
      err := os.Remove(fmt.Sprintf("./tmp/%s", file))
      if err != nil {
        fmt.Println(err)
      }
    }

    m.Reply("Berhasil menghapus semua sampah di folder tmp")

      m.React("✅")

    },
  })
}
