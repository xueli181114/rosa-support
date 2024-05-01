/*
Copyright (c) 2024 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package delete

import (
	"github.com/openshift-online/rosa-support/cmd/rosa-support/delete/tag"
	"github.com/openshift-online/rosa-support/cmd/rosa-support/delete/vpc"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a resource from stdin",
	Long:    "Delete a resource from stdin",
}

func init() {
	Cmd.AddCommand(vpc.Cmd)
	Cmd.AddCommand(tag.Cmd)
}
