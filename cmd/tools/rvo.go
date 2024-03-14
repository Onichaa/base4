package cmd

import (
  x "mywabot/system"

  "context"
  "encoding/json"
  "fmt"
  waProto "github.com/amiruldev20/waSocket/binary/proto"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(rvo|readviewonce)",
    Cmd:    []string{"rvo"},
    Tags:   "tools",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    var yh map[string]interface{}
    quoted := m.Quoted.GetQuotedMessage()
    if quoted == nil {
      m.Reply("reply message of type viewonce.")
      return
    } else {
      if quoted.ViewOnceMessageV2.GetMessage() == nil && quoted.ViewOnceMessage.GetMessage() == nil &&  quoted.ViewOnceMessageV2Extension.GetMessage() == nil {
        m.Reply("the message is not viewonce.")
        return
      }
    }

    if quoted.ViewOnceMessageV2 != nil {
      quoted = quoted.ViewOnceMessageV2.Message
    } else if quoted.ViewOnceMessageV2Extension != nil{
      quoted = m.Quoted.GetQuotedMessage().ViewOnceMessageV2Extension.Message
    }else{
      quoted = m.Quoted.GetQuotedMessage().ViewOnceMessage.Message
    }
    ok, _ := json.Marshal(quoted)
    err := json.Unmarshal(ok, &yh)
    if err != nil {
      m.Reply(fmt.Sprintf("%v", err))
      return
    }
    for k := range yh {
      yh[k].(map[string]interface{})["viewOnce"] = false
    }
    var b waProto.Message
    ok, _ = json.Marshal(yh)
    err = json.Unmarshal(ok, &b)
    if err != nil {
      m.Reply(fmt.Sprintf("%v", err))
      return
    }
    _, err = sock.WA.SendMessage(context.Background(), m.From, &b)
    if err != nil {
      m.Reply(fmt.Sprintf("%v", err))
      return
    }

      m.React("✅")
    },
  })
}
