package cmd

import (
	"os"
	"sync"

	"github.com/spf13/cobra"
)

type NextDjangoCmd struct {
	cmd *cobra.Command

	DjangoCmd
	NextCmd
}

var nextDjango NextDjangoCmd

func init() {
	nextDjango.cmd = &cobra.Command{
		Use:   "next-django [appname]",
		Short: "Create a fullstack app with a nextjs for the frontend and django as the backend",
		Args:  NameExists,
		Run:   nextDjango.run,
	}

	rootCmd.AddCommand(nextDjango.cmd)

	nextDjango.cmd.Flags().BoolP("help", "h", false, "help for django")
	nextDjango.cmd.Flags().BoolVarP(&nextDjango.auth, "auth", "a", false, "Setup authentication for frontend and backend")
	nextDjango.cmd.Flags().BoolVarP(&nextDjango.typescript, "typescript", "t", false, "Create a typescript project on the frontend")
	nextDjango.cmd.Flags().BoolVarP(&nextDjango.redux, "redux", "r", false, "Add redux to the frontend")
	nextDjango.cmd.Flags().BoolVarP(&nextDjango.materialUI, "mui", "m", false, "Add material-ui as a UI Library")

	nextDjango.cmd.Flags().BoolVarP(&nextDjango.restframework, "restframework", "f", false, "Install and setup DjangoRestFramework")
	nextDjango.cmd.Flags().BoolVarP(&nextDjango.jwt, "jwt", "j", false, "Add JSON Web Tokens to use for user authentication")
	nextDjango.cmd.Flags().BoolVarP(&nextDjango.cors, "cors", "c", false, "Install django-cors-headers")
}

func (nj *NextDjangoCmd) run(cmd *cobra.Command, args []string) {
	appDir := args[0]
	var wg sync.WaitGroup

	err := os.Mkdir(appDir, 0755)
	if err != nil {
		ExitGracefully(err)
	}

	err = os.Chdir(appDir)
	if err != nil {
		ExitGracefully(err)
	}

	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		nj.NextCmd.run(cmd, []string{"frontend"})
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		nj.DjangoCmd.run(cmd, []string{"backend"})
	}(&wg)

	wg.Wait()
}
