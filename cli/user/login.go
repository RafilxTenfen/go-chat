package user

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/chat"
	"github.com/RafilxTenfen/go-chat/database"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/ibraimgm/libcmd"
	null "github.com/rhizomplatform/pg-null"
)

func (t *terminalHandler) login(cmd *libcmd.Cmd) error {
	email, pwd, err := emailPwdFromCmd(cmd, t)
	if err != nil {
		return err
	}

	db, err := database.DBConnect()
	if err != nil {
		return err
	}

	usr := app.User{
		Email:    null.S(email),
		Password: null.S(pwd),
	}

	usrDB := store.FindUser(db, usr)
	if usrDB == nil {
		return fmt.Errorf("User of email %s doesn't exists", usr.Email.String)
	}

	if err := usrDB.VerifyPwd(usr); err != nil {
		return err
	}

	shell := chat.CreateUserShell(usrDB, db)
	if shell == nil {
		return fmt.Errorf("Error on create user shell")
	}

	shell.Run()
	return nil
}
