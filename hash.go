package main

import (
	"fmt"
	"code.google.com/p/go.crypto/bcrypt"
)

var cmdHash = &Command{
	UsageLine: "hash [password]",
	Short:     "generate a password hash for inserting into the database",
	Long: `
The administrator password is kept in the datastore. To update that password you generate
a bcrypt hash using this command and then manually put it into the datastore.
	`,
}

func init() {
	cmdHash.Run = runHash // break init loop
}

func runHash(cmd *Command, args []string) {
	if len(args) != 0 {
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(args[0]), 10)
	fmt.Println(string(hash))
	return
}
