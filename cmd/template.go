package cmd

import (
	"html/template"
	"os"
	"path"
	"strings"

	tmplFuncs "bitbucket.org/LunarWay/shuttle/pkg/templates"
	"github.com/spf13/cobra"
)

type context struct {
	Vars interface{}
	Args map[string]string
}

var templateCmd = &cobra.Command{
	Use:   "template [template]",
	Short: "Execute a template",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var templateName = args[0]
		loadedPlan := getPlan()

		namedArgs := map[string]string{}
		for _, arg := range args[1:] {
			parts := strings.SplitN(arg, "=", 2)
			namedArgs[parts[0]] = parts[1]
		}

		templatePath := path.Join(loadedPlan.PlanPath, templateName)

		tmpl, err := template.New(templateName).Funcs(tmplFuncs.GetFuncMap()).ParseFiles(templatePath)

		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(os.Stdout, context{
			Args: namedArgs,
			Vars: loadedPlan.ShuttleConfig.Variables,
		})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
