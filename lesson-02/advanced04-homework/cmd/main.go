package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/PraserX/fullstack-rookie/pkg/crypter"

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

func credentials() (string, error) {
	var err error

	fmt.Print("Enter password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic("Unhandled error")
	}
	password := string(bytePassword)
	fmt.Println()

	fmt.Print("Enter password again: ")
	bytePassword2, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic("Unhandled error")
	}
	password2 := string(bytePassword2)
	fmt.Println()

	if password != password2 {
		return "", fmt.Errorf("password do not match")
	}

	return strings.TrimSpace(password), nil
}

func main() {
	var err error

	app := &cli.App{
		Name:        "NoTSeCURe MessAGing",
		Usage:       "This is super secret not secure messaging system",
		Version:     "1.0.0",
		Description: "",
		Copyright:   "(c) 2021 National Cyber and Information Security Agency, Czech Republic",
		Commands: []*cli.Command{
			{
				Name:      "encrypt",
				Usage:     "encrypt file on input",
				ArgsUsage: "file",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						var err error
						var message string
						var password string

						var options []crypter.Option
						options = append(options, crypter.OptionFile(c.Args().Get(0)))

						c := crypter.New(options...)

						if password, err = credentials(); err != nil {
							return fmt.Errorf("%v", err)
						}

						c.SetPassword(password)

						if err = c.OpenFile(); err != nil {
							return fmt.Errorf("%v", err)
						}

						if message, err = c.Encrypt(); err != nil {
							return fmt.Errorf("%v", err)
						}

						fmt.Println(message)
						return nil
					}

					return fmt.Errorf("missing input file")
				},
			},
			{
				Name:      "decrypt",
				Usage:     "decrypt file on input",
				ArgsUsage: "file",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						var err error
						var message string
						var password string

						var options []crypter.Option
						options = append(options, crypter.OptionFile(c.Args().Get(0)))

						c := crypter.New(options...)

						if password, err = credentials(); err != nil {
							return fmt.Errorf("%v", err)
						}

						c.SetPassword(password)

						if err = c.OpenFile(); err != nil {
							return fmt.Errorf("%v", err)
						}

						if message, err = c.Decrypt(); err != nil {
							return fmt.Errorf("%v", err)
						}

						fmt.Println(message)
						return nil
					}

					return fmt.Errorf("missing input file")
				},
			},
		},
		Action: func(c *cli.Context) error {
			return fmt.Errorf("see help for more information")
		},
	}

	if err = app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error! %s\n", err.Error())
		os.Exit(1)
	}
}
