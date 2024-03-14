package cmd

import (
  x "mywabot/system"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(setdeskripsi|setbiogrub)",
    Cmd:    []string{"setbiogrub"},
    Tags:   "group",
    Prefix: true,
    IsAdmin: true,
    IsBotAdm: true,
    IsGroup: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      err := sock.WA.SetGroupTopic(m.From, "", "", m.Query)
      if err != nil {
        m.Reply("Failed to change the group description")
        return
      }
      m.Reply("Successfully change the group description")

      m.React("✅")
    },
  })
}
