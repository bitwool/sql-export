package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bitwool/sql-export/conn"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Usage:   "MySQL Server Address",
				Value:   "127.0.0.1",
				Aliases: []string{"h"},
				// Required: true,
			},
			&cli.StringFlag{
				Name:    "port",
				Usage:   "Port",
				Value:   "3306",
				Aliases: []string{"P"},
				// Required: true,
			},
			&cli.StringFlag{
				Name:    "user",
				Usage:   "MySQL User",
				Aliases: []string{"u"},
				// Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Usage:    "Password",
				Aliases:  []string{"p"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "dbname",
				Usage:    "Database Name",
				Aliases:  []string{"db"},
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			host := cmd.String("host")
			user := cmd.String("user")
			port := cmd.String("port")
			password := cmd.String("password")
			dbname := cmd.String("dbname")
			fmt.Printf("host: %s user: %s port: %s password: %s\n", host, user, port, password)
			conn.Connect(host, port, user, password, dbname)
			return nil
		},
		HideHelp: true,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
