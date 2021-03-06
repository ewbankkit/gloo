package secret

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/argsutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"

	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/surveyutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/spf13/cobra"
)

func azureCmd(opts *options.Options) *cobra.Command {
	input := opts.Create.InputSecret.AzureSecret
	cmd := &cobra.Command{
		Use:   "azure",
		Short: `Create an Azure secret with the given name`,
		Long:  `Create an Azure secret with the given name`,
		RunE: func(c *cobra.Command, args []string) error {
			if err := argsutils.MetadataArgsParse(opts, args); err != nil {
				return err
			}
			if opts.Top.Interactive {
				// and gather any missing args that are available through interactive mode
				if err := AzureSecretArgsInteractive(&opts.Metadata, &input); err != nil {
					return err
				}
			}
			// create the secret
			if err := createAzureSecret(opts.Top.Ctx, opts.Metadata, input); err != nil {
				return err
			}
			fmt.Printf("Created Azure secret [%v] in namespace [%v]\n", opts.Metadata.Name, opts.Metadata.Namespace)
			return nil
		},
	}

	flags := cmd.Flags()
	flagutils.AddMetadataFlags(flags, &opts.Metadata)
	flags.StringSliceVar(&input.ApiKeys.Entries, "api-keys", []string{}, "comma-separated list of azure api key=value entries")

	return cmd
}

func AzureSecretArgsInteractive(meta *core.Metadata, input *options.AzureSecret) error {
	if err := surveyutils.InteractiveNamespace(&meta.Namespace); err != nil {
		return err
	}

	if err := cliutil.GetStringInput("name of secret", &meta.Name); err != nil {
		return err
	}
	if err := cliutil.GetStringSliceInput("Enter API key entry (key=value)", &input.ApiKeys.Entries); err != nil {
		return err
	}

	return nil
}

func createAzureSecret(ctx context.Context, meta core.Metadata, input options.AzureSecret) error {
	if input.ApiKeys.Entries == nil {
		return errors.Errorf("must provide azure api keys")
	}
	secret := &gloov1.Secret{
		Metadata: meta,
		Kind: &gloov1.Secret_Azure{
			Azure: &gloov1.AzureSecret{
				ApiKeys: input.ApiKeys.MustMap(),
			},
		},
	}

	secretClient := helpers.MustSecretClient()
	if _, err := secretClient.Write(secret, clients.WriteOpts{Ctx: ctx}); err != nil {
		return err
	}
	return nil
}
