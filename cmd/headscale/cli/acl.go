package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	v1 "github.com/juanfont/headscale/gen/go/headscale/v1"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func init() {
	rootCmd.AddCommand(aclCmd)
	aclCmd.AddCommand(getAclCmd)

	setAclCmd.Flags().StringP("file", "f", "", "File to read the acl from")
	aclCmd.AddCommand(setAclCmd)
}

var aclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Manage the acl of Headscale",
}

var getAclCmd = &cobra.Command{
	Use:     "get",
	Short:   "get acl",
	Aliases: []string{"show"},
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")

		ctx, client, conn, cancel := getHeadscaleCLIClient()
		defer cancel()
		defer conn.Close()

		response, err := client.GetACL(ctx, &v1.GetACLRequest{})
		if err != nil {
			ErrorOutput(
				err,
				fmt.Sprintf("Cannot get acl: %s", status.Convert(err).Message()),
				output,
			)

			return
		}

		if output == "" {
			output = "json"
		}

		SuccessOutput(response.Policy, "", output)
	},
}

var setAclCmd = &cobra.Command{
	Use:     "set",
	Short:   "set acl",
	Aliases: []string{"import"},
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			ErrorOutput(
				err,
				fmt.Sprintf("Error getting file from flag: %s", err),
				output,
			)

			return
		}

		var reader io.Reader
		if file == "" {
			reader = os.Stdin
		} else {
			f, err := os.Open(file)
			if err != nil {
				ErrorOutput(
					err,
					fmt.Sprintf("Error opening file: %s", err),
					output,
				)
			}
			defer f.Close()
			reader = f
		}

		content, err := io.ReadAll(reader)
		if err != nil {
			ErrorOutput(
				err,
				fmt.Sprintf("Error reading file: %s", err),
				output,
			)

			return
		}

		log.Debug().Str("file", string(file)).Msg("Reading file")

		ctx, client, conn, cancel := getHeadscaleCLIClient()
		defer cancel()
		defer conn.Close()

		var m map[string]interface{}
		if err := json.Unmarshal(content, &m); err != nil {
			ErrorOutput(
				err,
				fmt.Sprintf("Error unmarshalling file: %s", err),
				output,
			)

			return
		}
		polPb, err := structpb.NewStruct(m)
		if err != nil {
			ErrorOutput(
				err,
				fmt.Sprintf("Error creating struct from map: %s", err),
				output,
			)

			return
		}

		if _, err := client.SetACL(ctx, &v1.SetACLRequest{
			Policy: polPb,
		}); err != nil {
			ErrorOutput(
				err,
				fmt.Sprintf("Cannot set acl: %s", status.Convert(err).Message()),
				output,
			)

			return
		}
	},
}
