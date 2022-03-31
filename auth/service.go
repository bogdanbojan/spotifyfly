package auth

import (
	"fmt"
	"github.com/fabioberger/chrome"
	"honnef.co/go/js/console"
)

func main() {
	c := chrome.NewChrome()
	color := "#3aa757"
	c.Runtime.OnInstalled(func(details map[string]string) {
		sync := chrome.Storage{}
		sync.Sync.Set(color,
			func() {
				console.Log(fmt.Sprintf("Default background color set to %v", color))
			})
	})
}
