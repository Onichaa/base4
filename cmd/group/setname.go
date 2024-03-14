package cmd

import (
  x "mywabot/system"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(setname|setnama|setnamagrub)",
    Cmd:    []string{"setname"},
    Tags:   "group",
    Prefix: true,
    IsAdmin: true,
    IsBotAdm: true,
    IsGroup: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    err := sock.WA.SetGroupName(m.From, m.Query)
    if err != nil {
      m.Reply("Failed to change the group name")
      return
    }
    m.Reply("Berhasil mengubah nama group")
      
      m.React("✅")
    },
  })
}
