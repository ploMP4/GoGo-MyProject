package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type DjangoCmd struct {
	restframework bool
	auth          bool
	jwt           bool
	cors          bool
	cmd           *cobra.Command
}

var django DjangoCmd

func init() {
	django.cmd = &cobra.Command{
		Use:   "django [appname]",
		Short: "Create Django application",
		Args:  NameExists,
		Run:   django.run,
	}

	rootCmd.AddCommand(django.cmd)

	django.cmd.Flags().BoolP("help", "h", false, "help for django")
	django.cmd.Flags().BoolVarP(&django.restframework, "restframework", "r", false, "Install and setup DjangoRestFramework")
	django.cmd.Flags().BoolVarP(&django.jwt, "jwt", "j", false, "Add JSON Web Tokens to use for user authentication")
	django.cmd.Flags().BoolVarP(&django.cors, "cors", "c", false, "Install django-cors-headers")
	django.cmd.Flags().BoolVarP(&django.auth, "auth", "a", false, "Create users django app with custom authentication (NOTE: Uses JWT)")
}

func (d *DjangoCmd) run(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup

	// Create app dir
	appName := args[0]

	// Create virtualenv
	color.Green("Creating python virtual environment")
	s := LoadSpinner()
	s.Start()

	d.createVirtualEnv()

	// Install django
	s.Restart()
	color.Green("Installng django")

	c := exec.Command("./env/bin/pip", "install", "django")
	err := c.Run()
	if err != nil {
		ExitGracefully(err)
	}

	// Initialize django project
	s.Restart()
	color.Green("Creating django project: " + appName)

	c = exec.Command("./env/bin/django-admin", "startproject", appName)
	err = c.Run()
	if err != nil {
		ExitGracefully(err)
	}

	if d.restframework {
		wg.Add(1)
		s.Restart()
		color.Green("Installing DjangoRestFramework")
		go d.installRestFramework(&wg, appName)
	}

	if d.cors {
		wg.Add(1)
		s.Restart()
		color.Green("Installing CORS Headers")
		go d.installCORS(&wg, appName)
	}

	if d.jwt {
		wg.Add(1)
		s.Restart()
		color.Green("Installing SimpleJWT")
		go d.installJWT(&wg, appName)
	}

	wg.Wait()
	s.Stop()
	ExitGracefully(nil, "Django project created successfully under name: "+appName)
}

func (d *DjangoCmd) createVirtualEnv() {
	c := exec.Command("virtualenv", "env")

	err := c.Run()
	if err != nil {
		ExitGracefully(err, "Couldn't Create virtual env, make sure you have virtual env installed (pip install virtualenv)")
	}
}

func (d *DjangoCmd) editSettings(appName, splitOn, toAppend string) {
	var settings string

	content, err := os.ReadFile(fmt.Sprintf("%s/%s/settings.py", appName, appName))
	if err != nil {
		ExitGracefully(err)
	}

	if splitOn != "" {
		s := strings.Split(string(content), splitOn)
		s[0] += toAppend

		settings = strings.Join(s, " ")
	} else {
		settings = string(content) + toAppend
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s/settings.py", appName, appName), []byte(settings), 0644)
	if err != nil {
		ExitGracefully(err)
	}
}

func (d *DjangoCmd) installRestFramework(wg *sync.WaitGroup, appName string) {
	defer wg.Done()

	c := exec.Command("./env/bin/pip", "install", "djangorestframework")

	err := c.Run()
	if err != nil {
		ExitGracefully(err, "Unable to install django rest framework")
	}

	// Add restframework to settings.py
	d.editSettings(appName, "'django.contrib.staticfiles',", "'django.contrib.staticfiles',\n\t'rest_framework',")
}

func (d *DjangoCmd) installCORS(wg *sync.WaitGroup, appName string) {
	defer wg.Done()

	c := exec.Command("./env/bin/pip", "install", "django-cors-headers")

	err := c.Run()
	if err != nil {
		ExitGracefully(err, "Unable to install django-cors-headers")
	}

	// Add cors headers to INSTALLED_APPS in settings.py
	d.editSettings(appName, "'django.contrib.staticfiles',", "'django.contrib.staticfiles',\n\t'corsheaders',")

	// Add cors headers to MIDDLEWARE in settings.py
	d.editSettings(appName, "MIDDLEWARE = [", "MIDDLEWARE = [\n\t'corsheaders.middleware.CorsMiddleware',")

	// Set cors origins
	d.editSettings(appName, "", "\n# Change in production\nCORS_ALLOW_ALL_ORIGINS = True")
}

func (d *DjangoCmd) installJWT(wg *sync.WaitGroup, appName string) {
	defer wg.Done()

	c := exec.Command("./env/bin/pip", "install", "djangorestframework-simplejwt")

	err := c.Run()
	if err != nil {
		ExitGracefully(err, "Unable to install django-simplejwt")
	}

	// Add JWT to settings.py
	appendString := `REST_FRAMEWORK = {
	'DEFAULT_AUTHENTICATION_CLASSES': (
		'rest_framework_simplejwt.authentication.JWTAuthentication',
	)
}
	
MIDDLEWARE = [`

	d.editSettings(appName, "MIDDLEWARE = [", appendString)
}

func (d *DjangoCmd) createAuth() {

}
