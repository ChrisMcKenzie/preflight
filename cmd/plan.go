package cmd

import (
	"fmt"
	"log"
	"os/exec"

	pfPlugin "github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/ChrisMcKenzie/preflight/preflight"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// bytes, err := ioutil.ReadFile(args[0])
		// if err != nil {
		// 	log.Println(err)
		// }
		//
		// file, err := hcl.ParseBytes(bytes)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		//
		// cl, err := preflight.LoadHcl(file)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// We're a host! Start by launching the plugin process.
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: pfPlugin.Handshake,
			Plugins:         pfPlugin.PluginMap,
			Cmd:             exec.Command(args[0]),
		})
		defer client.Kill()

		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
		}

		raw, err := rpcClient.Dispense("greeter")
		if err != nil {
			log.Println(err)
		}
		greeter := raw.(preflight.Greeter)
		fmt.Println(greeter.Greet())
		// cl.Plan()
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
