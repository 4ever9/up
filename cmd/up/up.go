package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/4ever9/up"
	"github.com/spf13/cobra"
)

var root string
var port uint32

var logger = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	upCMD.Flags().StringVar(&root, "root", "./", "Show all version info")
	upCMD.Flags().Uint32Var(&port, "port", 8000, "Listening port")
}

var upCMD = &cobra.Command{
	Use:   "serve [option]",
	Short: "serve - Static file serving and directory listing",
	Run: func(cmd *cobra.Command, args []string) {
		s := up.NewServer()
		logger.Printf("listening on http://localhost:%d", port)
		logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s))
	},
}

func main() {
	upCMD.AddCommand(versionCMD)
	if err := upCMD.Execute(); err != nil {
		logger.Fatalln(err)
	}
}
