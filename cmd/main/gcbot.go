package cmd

import (
  x "mywabot/system"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(gcbot|grubbot|groupbot)",
    Cmd:    []string{"grubbot"},
    Tags:   "main",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {

      m.Reply(`*Jangan lupa join ke grub bot official ini:*

*Group Bot Official*
> • https://chat.whatsapp.com/DaBXFf82aqwHc03v22E09D


*Fearles*
> • https://chat.whatsapp.com/G78M55qPpAk03UCSBgkdN3`)

    },
  })
}
