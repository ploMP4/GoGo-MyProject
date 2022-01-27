package main

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

type CppProject struct {
	buildTool string
	dirs      []string
	projType  string
}

type ICppProject interface {
	parseArgs(args []string)
	create(args []string)
	createVanilla()
}

func (c *CppProject) parseArgs(args []string) {
	buildTools := []string{"cmake"}
	types := []string{"opengl", "vanilla"}

	if len(args) > 0 {
		for _, arg := range args {
			for _, tool := range buildTools {
				if strings.ToLower(arg) == tool {
					c.buildTool = arg
					break
				}
			}

			for _, t := range types {
				if strings.ToLower(arg) == t {
					c.projType = arg
					break
				}
			}
		}
	}
}

func (c *CppProject) create(args []string) {
	c.parseArgs(args)

	switch c.buildTool {
	case "cmake":
		color.Yellow("CMake")

	default:
		color.Green("Creating C++ application: %s", ggp.appname)
		c.createVanilla()
	}
}

func (c *CppProject) createVanilla() {
	c.dirs = []string{"bin", "includes", "src"}

	err := os.Mkdir(ggp.appname, 0755)
	if err != nil {
		exitGracefully(err)
	}

	err = os.Chdir(ggp.appname)
	if err != nil {
		exitGracefully(err)
	}

	color.Green("Creating Project Structure...")
	for _, dir := range c.dirs {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			exitGracefully(err)
		}
	}

	main := ggp.rootPath + "/src/main.cpp"
	err = copyFileFromTemplate("templates/cpp/main.cpp.txt", main)
	if err != nil {
		exitGracefully(err)
	}

	color.Green("Creating Makefile...")
	makefile := ggp.rootPath + "/Makefile"
	err = copyFileFromTemplate("templates/cpp/Makefile", makefile)
	if err != nil {
		exitGracefully(err)
	}
}
