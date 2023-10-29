package controller

import (
	"fmt"
	"os/exec"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/gin-gonic/gin"

	"github.com/hashicorp/go-plugin"
)

func InvokePlugin(c *gin.Context) {
	fmt.Println("a")

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &plug.PluginConnector{},
		},
		Cmd: exec.Command("./plugins/teatime-plugin-pinit/teatime-plugin-pinit"),
	})
	defer client.Kill()

	rpcc, err := client.Client()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	raw, err := rpcc.Dispense("main")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	info := raw.(plug.PluginInterface).Info()
	fmt.Println(info)
}
