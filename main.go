package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thehxdev/tbc/client"
)

var (
	fs map[string]*flag.FlagSet = make(map[string]*flag.FlagSet)

	password string
	name     string
	txtid    string
	path     string
	conf     string
)

func main() {
	getEnvVars()
	configureCmdFlags()

	args := os.Args
	subcmd := args[1]
    if subcmd == "help" {
        usage()
        os.Exit(0)
    }

	u, err := client.Init(conf)
	if err != nil {
		log.Fatal(err)
	}


	err = fs[subcmd].Parse(args[2:])
	if err != nil {
		u.ErrLogger.Fatal(err)
	}

	var d []byte

	switch subcmd {
	case "useradd":
		err = u.UseraddHandler(password)

	case "userdel":
		err = u.UserdelHandler()

	case "tee":
		d, err = u.TeeHandler(name, path)

	case "ls":
		d, err = u.LsHandler()

	case "rm":
		err = u.RmHandler(txtid)
	
	case "chtxt":
		err = u.ChtxtHandler(txtid, path)

	case "rename":
		err = u.RenameHandler(name, txtid)
	case "mv":
		d, err = u.MvHandler(txtid)
	}

	if err != nil {
		u.ErrLogger.Fatal(err)
	}

	fmt.Print(string(d))
}

func getEnvVars() {
	var ok bool
	if conf, ok = os.LookupEnv("TBC_CONF"); !ok {
		conf = "./config.json"
	}
}

func configureCmdFlags() {
	fs["useradd"] = flag.NewFlagSet("useradd", flag.ExitOnError)
	fs["useradd"].StringVar(&password, "password", "", "Your Password")

	fs["userdel"] = flag.NewFlagSet("userdel", flag.ExitOnError)

	fs["ls"] = flag.NewFlagSet("ls", flag.ExitOnError)

	fs["tee"] = flag.NewFlagSet("tee", flag.ExitOnError)
	fs["tee"].StringVar(&name, "name", "", "New txt name")
	fs["tee"].StringVar(&path, "path", "", "File path")

	fs["rm"] = flag.NewFlagSet("rm", flag.ExitOnError)
	fs["rm"].StringVar(&txtid, "id", "", "Txt ID")

	fs["mv"] = flag.NewFlagSet("mv", flag.ExitOnError)
	fs["mv"].StringVar(&txtid, "id", "", "Txt ID")

	fs["chtxt"] = flag.NewFlagSet("chtxt", flag.ExitOnError)
	fs["chtxt"].StringVar(&txtid, "id", "", "Txt ID")
	fs["chtxt"].StringVar(&path, "path", "", "File path")

	fs["rename"] = flag.NewFlagSet("rename", flag.ExitOnError)
	fs["rename"].StringVar(&txtid, "id", "", "Txt ID")
	fs["rename"].StringVar(&name, "name", "", "Txt name")
}


func usage() {
    msg := `tbc Usage:
    useradd (Create new user)
        -password (New account password)

    tee (Create a new txt)
        -name (Name of new txt)
        -path (File path to read txt content)

    ls (Get all txts you created)

    rm (remove a txt)
        -id (Txt ID to remove)

    chtxt (Change a txt content)
        -id (Txt ID to change content)
        -path (File path to read txt content)

    rename (Change a txt name)
        -id (Txt id to change name)
        -name (Txt new name)
    

    mv (Change a txt id)
        -id (Txt id to change id)

===============
Env Vars:
    TBC_CONF (Path to config.json file)
`

    fmt.Println(msg)
}
