package main

import (
	"flag"
	"log"
	"os"
	"go9p.googlecode.com/hg/p"
	"go9p.googlecode.com/hg/p/clnt"
)

var debuglevel = flag.Int("d", 0, "debuglevel")
var addr = flag.String("addr", "127.0.0.1:5640", "network address")

func main() {
	var user p.User
	var err os.Error
	var c *clnt.Clnt
	var file *clnt.File
	var d []*p.Dir

	flag.Parse()
	user = p.OsUsers.Uid2User(os.Geteuid())
	clnt.DefaultDebuglevel = *debuglevel
	c, err = clnt.Mount("tcp", *addr, "", user)
	if err != nil {
		goto error
	}

	if flag.NArg() != 1 {
		log.Println("invalid arguments")
		return
	}

	file, err = c.FOpen(flag.Arg(0), p.OREAD)
	if err != nil {
		goto error
	}

	for {
		d, err = file.Readdir(0)
		if d == nil || len(d) == 0 || err != nil {
			break
		}

		for i := 0; i < len(d); i++ {
			os.Stdout.WriteString(d[i].Name + "\n")
		}
	}

	file.Close()
	if err != nil && err != os.EOF {
		goto error
	}

	return

error:
	log.Println(err)
}
