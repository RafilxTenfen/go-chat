package user

import (
	"fmt"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/RafilxTenfen/go-chat/database"
	"github.com/RafilxTenfen/go-chat/store"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/log"
)

func (t *terminalHandler) add(cmd *libcmd.Cmd) error {
	email, pwd, err := emailPwdFromCmd(cmd, t)
	if err != nil {
		return err
	}

	usr, err := app.NewUser(email, pwd)
	if err != nil {
		return err
	}

	db, err := database.DBConnect()
	if err != nil {
		return err
	}

	if err := store.InsertUser(db, usr); err != nil {
		return err
	}

	log.With(log.F{
		"email": usr.Email.String,
	}).Debug("User inserted")

	return nil
}

func emailPwdFromCmd(cmd *libcmd.Cmd, t *terminalHandler) (email, pwd string, err error) {
	email = *cmd.GetString("email")
	pwd = *cmd.GetString("password")

	if !app.ValidEmail(email) {
		return "", "", fmt.Errorf("Invalid Email '%s'", email)
	}

	if pwd == "" {
		t.term.Printf("A user must have a password\nType the new password:")
		pwd = t.term.ReadPassword()
	}

	return email, pwd, nil
}
