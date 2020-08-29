package chatapi

import (
	"github.com/RafilxTenfen/go-chat/api"
	"github.com/ibraimgm/libcmd"
)

func (t *terminalHandler) run(cmd *libcmd.Cmd) error {
	port := cmd.GetUint16("port")

	return api.Run(*port)
}
