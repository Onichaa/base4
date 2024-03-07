package cmd

import (
	"fmt"
	x "mywabot/system"

	"github.com/amiruldev20/waSocket"
	"github.com/fatih/color"
  "time"
)

func init() {

	x.ListenerAdd(
		func(c *waSocket.Client, m *x.IMsg) {

      wita, _ := time.LoadLocation("Asia/Makassar")
      currentTime := time.Now().In(wita)
      
			if m.IsGroup {
				get, err := c.GetGroupInfo(m.From)

				if err != nil {
					fmt.Println(err)
					return
				}
        magenta := color.New(color.FgGreen).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
        blue := color.New(color.FgBlue).SprintFunc()
        cyan := color.New(color.FgCyan).SprintFunc()
    
        fmt.Println(yellow("[GROUP]"), magenta(currentTime.Format("15:04:05")), blue(m.Text), yellow("dari"), yellow(m.PushName), cyan("Di Group"), cyan(get.Name))
			} else {
				magenta := color.New(color.FgGreen).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
        blue := color.New(color.FgBlue).SprintFunc()
     
        fmt.Println(yellow("[PRIVATE]"), magenta(currentTime.Format("15:04:05")), blue(m.Text), yellow("dari"), yellow(m.PushName))	
      }
		},
	)
}
