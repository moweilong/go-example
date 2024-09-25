package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"go-plugin/shared"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/greeter"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("some error0")
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		fmt.Println("some error1")
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(shared.Greeter)
	fmt.Println(greeter.Greet())
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
//
// plugin.HandshakeConfig是客户端和服务器在启动插件连接之前用于握手的配置。
// 这是由ServeConfig和ClientConfig嵌入的。在实践中，插件主机会创建
// 一个HandshakeConfig，并将其导出，然后插件就可以轻松地使用它。
var handshakeConfig = plugin.HandshakeConfig{
	// ProtocolVersion是客户端必须匹配才能进行通信的版本。当使用插件时，
	// 这应该与ClientConfig上设置的ProtocolVersion相匹配。
	// 如果在客户端或服务器配置中使用VersionedPlugins，则不需要此字段。
	ProtocolVersion: 1,
	// MagicCookieKey和value用于验证插件是否要启动。这不是一种安全措施，
	// 只是一种用户体验功能。如果魔法饼干不匹配，我们将显示对人类友好的输出。
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
// pluginMap是我们可以分发的插件的映射。
var pluginMap = map[string]plugin.Plugin{
	"greeter": &shared.GreeterPlugin{},
}
