package clihandler

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

//Out is writer to standard output
var Out io.Writer = os.Stdout

// CreateCLIHandler returns *cli.App which handles CLI
func CreateCLIHandler() *cli.App {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var operation string
	app := &cli.App{
		Name:  "calc",
		Usage: "Perform addition, subtraction, multiplication and division on 2 numbers!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "operation",
				Aliases:     []string{"op"},
				Usage:       "choose operation",
				Destination: &operation,
				Required:    true,
				DefaultText: "add, sub, mul, div",
			},
		},
		Action: func(cCtx *cli.Context) error {
			var num1, num2, result float64
			var err1, err2 error
			if cCtx.NArg() == 2 {
				num1, err1 = strconv.ParseFloat(cCtx.Args().Get(0), 64)
				num2, err2 = strconv.ParseFloat(cCtx.Args().Get(1), 64)
				if err1 != nil || err2 != nil {
					err := errors.New("arguments should be numbers")
					log.Fatal().Err(err).Msg("")
				}
			} else {
				err := errors.New("you have to provide 2 arguments")
				log.Fatal().Err(err).Msg("")
			}
			switch operation {
			case "add":
				result = num1 + num2
				fmt.Fprintln(Out, result)
			case "sub":
				result = num1 - num2
				fmt.Fprintln(Out, result)
			case "mul":
				result = num1 * num2
				fmt.Fprintln(Out, result)
			case "div":
				if num2 == 0 {
					err := errors.New("you cannot divide by 0")
					log.Fatal().Err(err).Msg("")
				}
				result = num1 / num2
				fmt.Fprintln(Out, result)
			default:
				err := errors.New("unknown operation")
				log.Fatal().Err(err).Msg("")
			}

			log.Debug().
				Float64("Number 1", num1).
				Float64("Number 2", num2).
				Float64("Result", result).
				Msg("OK")

			return nil
		},
	}
	return app

}
