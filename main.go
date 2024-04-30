package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bitwool/sql-export/db"
	"github.com/bitwool/sql-export/export"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "Sql Export",
		Usage: "Export Data From MySQL Database By Select Query",
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
			&cli.StringFlag{
				Name:     "query",
				Usage:    "Select Query",
				Aliases:  []string{"q"},
				Required: true,
			},
			// &cli.BoolFlag{
			// 	Name:  "ignore-ai-pk",
			// 	Usage: "Ignore Auto Increment Primary Key",
			// },
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			host := cmd.String("host")
			user := cmd.String("user")
			port := cmd.String("port")
			password := cmd.String("password")
			dbname := cmd.String("dbname")
			query := cmd.String("query")

			// ignore_ai_pk := cmd.Bool("ignore-ai-pk")
			fmt.Printf("host: %s user: %s port: %s password: %s\n", host, user, port, password)
			db.Init(host, port, user, password, dbname, query)
			export.Query(query, false)
			db.Close()
			return nil
		},
		HideHelp: true,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
