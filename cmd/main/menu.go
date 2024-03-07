package cmd

import (
	"fmt"
	x "mywabot/system"
	"sort"
	"strings"

  waProto "github.com/amiruldev20/waSocket/binary/proto"
  "google.golang.org/protobuf/proto"
)

type item struct {
	Name   []string
	Prefix bool
}

type tagSlice []string

func (t tagSlice) Len() int {
	return len(t)
}

func (t tagSlice) Less(i int, j int) bool {
	return t[i] < t[j]
}

func (t tagSlice) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

func menu(client *x.Nc, m *x.IMsg) {
	var str string
	str += fmt.Sprintf("Halo, %s 👋\nBot ini masih dalam tahap beta\n\nHost: Linux\nLibrary: waSocket\n\n", m.PushName)
	var tags map[string][]item
	for _, list := range x.GetList() {
		if tags == nil {
			tags = make(map[string][]item)
		}
		if _, ok := tags[list.Tags]; !ok {
			tags[list.Tags] = []item{}
		}
		tags[list.Tags] = append(tags[list.Tags], item{Name: list.Cmd, Prefix: list.Prefix})
	}

	var keys tagSlice
	for key := range tags {
		keys = append(keys, key)
	}

	sort.Sort(keys)

	for _, key := range keys {
		str += fmt.Sprintf("「 *%s MENU* 」\n", strings.ToUpper(key))
		for _, e := range tags[key] {
			var prefix string
			if e.Prefix {
				prefix = m.Prefix[:1]
			} else {
				prefix = ""
			}
			for _, nm := range e.Name {
				str += fmt.Sprintf("ゝ %s%s\n", prefix, nm)
			}
		}
		str += "\n"
	}
	txt := str + "\n© Powered by GoLang"
	//m.Reply(txt)
  
  var isImage = waProto.ContextInfo_ExternalAdReplyInfo_IMAGE
  client.SendText(m.From, strings.TrimSpace(txt), &waProto.ContextInfo{
    ExternalAdReply: &waProto.ContextInfo_ExternalAdReplyInfo{
      Title:                 proto.String("Bot WhatsApp 2023"),
      Body:                  proto.String("Simple WhatsApp Bot"),
      MediaType:             &isImage,
      ThumbnailUrl:          proto.String("https://telegra.ph//file/46181dd650c4927a3578a.jpg"),
      MediaUrl:              proto.String("https://wa.me/stickerpack/inc.dev"),
      SourceUrl:             proto.String("https://id.pinterest.com/pin/6403624461768243/"),
      ShowAdAttribution:     proto.Bool(true),
      RenderLargerThumbnail: proto.Bool(true),
    }})
}

func init() {
	x.NewCmd(&x.ICmd{
		Name:   "(menu|makan|help)",
		Cmd:    []string{"menu"},
		Tags:   "main",
		Prefix: true,
		Exec:   menu,
	})
}
