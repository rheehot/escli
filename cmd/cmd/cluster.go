/*
copyright 2020 the Escli authors

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DevopsArtFactory/escli/cmd/cmd/builder"
	"github.com/DevopsArtFactory/escli/cmd/cmd/cluster"
)

func NewClusterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "get or set settings for cluster",
		Long:  "get or set settings for cluster",
	}

	clusterSettingsCommand := cluster.NewClusterSettingsCommand()
	clusterRerouteCommand := cluster.NewClusterRerouteCommand()

	cmd.AddCommand(clusterSettingsCommand)
	cmd.AddCommand(clusterRerouteCommand)

	builder.SetCommandFlags(clusterSettingsCommand)

	return cmd
}
