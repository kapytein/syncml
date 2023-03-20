package main

import (
	"fmt"

	"github.com/kapytein/syncml/pkg/messages"
	"github.com/kapytein/syncml/pkg/syncml"
)

func main() {
	syncMLRequestParams := messages.SyncMLRequestParameters{DeviceID: "123", MsgID: 1, SessionID: 1}
	var commands []syncml.SyncMLCommands

	policy := messages.NewMessage("basic_info", syncMLRequestParams)
	commands = append(commands, policy.GetCommands())

	response, err := policy.BuildSyncMLResponseMessage(commands)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))
}
