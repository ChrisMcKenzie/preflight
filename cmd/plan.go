package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/ChrisMcKenzie/preflight/config"
	pfPlugin "github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/ChrisMcKenzie/preflight/preflight"
	"github.com/fatih/color"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/hcl"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Println(err)
		}

		file, err := hcl.ParseBytes(bytes)
		if err != nil {
			log.Println(err)
			return
		}

		cl, err := config.LoadHcl(file)
		if err != nil {
			fmt.Println(err)
		}

		Plan(cl)
	},
}

func init() {
	RootCmd.AddCommand(planCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// planCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Plan ...
func Plan(cl *config.Config) {
	// We're a host! Start by launching the plugin process.
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	for _, task := range cl.Tasks {
		fmt.Printf("===== TASK: (%s) %s =====\n\n", green(task.Type), task.Name)
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: pfPlugin.Handshake,
			Plugins:         pfPlugin.PluginMap,
			Cmd:             exec.Command(fmt.Sprintf(".preflight/provisioner-%s", task.Type)),
		})
		defer client.Kill()

		log.SetOutput(ioutil.Discard)

		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			fmt.Printf(red("ERROR: %s\n\n"), err)
			break
		}

		raw, err := rpcClient.Dispense("provisioner")
		if err != nil {
			fmt.Println(err)
		}
		prov := raw.(preflight.Provisioner)

		var exists bool
		exists, err = prov.Exists(task)
		if err != nil {
			fmt.Printf(red("ERROR: %s\n\n"), err)
			break
		}

		if exists {
			fmt.Printf(yellow("- %s exists: no change needed\n\n"), task.RawConfig["name"])
		} else {
			fmt.Printf(green("+ %s is absent: installing\n\n"), task.RawConfig["name"].(string))
		}
	}
}
