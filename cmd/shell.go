/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"strings"

	"database/sql"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A Interactive input entry.",
	Long:  `A Interactive input entry.`,
	Run: func(cmd *cobra.Command, args []string) {
		var db *sql.DB
		dsn := viper.GetString("DBDSN")
		// pterm.Printfln("DBDSN: %s", dsn)
		if dsn != "" {
			var err error
			db, err = sql.Open("mysql", dsn)
			if err != nil {
				panic(err)
			}
			err = db.Ping()
			if err != nil {
				panic(err)
			}
			defer db.Close()
		}
		for {
			input, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("<--")
			if strings.ToLower(input) == "exit" {
				pterm.Println("\n-->: bye!")
				break
			}
			pterm.Println()
			rows, err := db.Query(input)
			if err != nil {
				pterm.Error.Printfln(err.Error())
				continue
			}

			cols, _ := rows.Columns()
			for rows.Next() {
				row := make([][]byte, len(cols))
				rowPtr := make([]interface{}, len(cols))
				for i := range row {
					rowPtr[i] = &row[i]
				}
				if err := rows.Scan(rowPtr...); err != nil {
					pterm.Error.Printfln(err.Error())
					continue
				}
				sep := []byte("\t")
				pterm.Printfln(string(bytes.Join(row, sep)))
			}
			if err := rows.Err(); err != nil {
				pterm.Error.Printfln(err.Error())
				continue
			}
			rows.Close()
		}

	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
