package main

import (
	"os"
	"os/exec"
	"sync"

	"github.com/fatih/color"
)

type ReactProject struct {
	command    []string
	redux      bool
	materialUI bool
	bootstrap  bool
}

func (r *ReactProject) parseArgs(args []string) {
	r.command = append(r.command, "npx")
	r.command = append(r.command, "create-react-app")
	r.command = append(r.command, ggp.appname)

	if len(args) > 0 {
		for _, arg := range args {
			switch arg {
			case "typescript", "ts":
				r.command = append(r.command, "--template")
				r.command = append(r.command, "typescript")
				color.Blue("Template: Typescript")
				continue

			case "redux":
				r.redux = true
				color.Blue("State management: Redux")
				continue

			case "materialUI", "material-ui", "mui":
				r.materialUI = true
				color.Blue("UI Library: Material-UI")
				continue

			case "bootstrap":
				r.bootstrap = true
				color.Blue("UI Library: Bootstrap")
				continue

			default:
				color.Red("ERROR: invalid argument: " + arg)
				color.Yellow("Skipping argument [" + arg + "]...")
			}
		}
	}
}

func (r *ReactProject) create(args []string) {
	var wg sync.WaitGroup

	color.Green("Creating React app: " + ggp.appname)

	r.parseArgs(args)

	cmd := exec.Command(r.command[0], r.command[1:]...)

	// https://blog.bitsrc.io/build-command-line-spinners-in-node-js-3e432d926d56
	// var i int
	// go func() {
	// 	for {
	// 		i++
	// 		if i%2 == 0 {
	// 			fmt.Println("\x1B[?25l/")
	// 			time.Sleep(1 * time.Second)
	// 		} else {
	// 			fmt.Println("-")
	// 			time.Sleep(1 * time.Second)
	// 		}
	// 	}
	// }()

	err := cmd.Run()
	if err != nil {
		exitGracefully(err)
	}

	// TODO: create redux boilerplate from template
	if r.redux {
		wg.Add(1)

		go func() {
			defer wg.Done()
			color.Green("Installing Redux...")

			os.Chdir(ggp.appname)
			cmd := exec.Command("npm", "install", "redux")

			err := cmd.Run()
			if err != nil {
				exitGracefully(err, "Unable to install redux")
			}
		}()
	}

	if r.materialUI {
		wg.Add(1)

		go func() {
			defer wg.Done()
			color.Green("Installing Material UI...")

			os.Chdir(ggp.appname)
			cmd := exec.Command("npm", "install", "@material-ui/core")

			err := cmd.Run()
			if err != nil {
				exitGracefully(err, "Unable to install material-ui")
			}
		}()
	}

	wg.Wait()
}
