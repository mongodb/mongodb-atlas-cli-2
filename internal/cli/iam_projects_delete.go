// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type iamProjectsDeleteOpts struct {
	projectID string
	confirm   bool
	store     store.ProjectDeleter
}

func (opts *iamProjectsDeleteOpts) init() error {
	s, err := store.New()

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *iamProjectsDeleteOpts) Run() error {
	if !opts.confirm {
		fmt.Println("Project not deleted")
		return nil
	}
	err := opts.store.DeleteProject(opts.projectID)

	if err != nil {
		return err
	}

	fmt.Printf("Project '%s' deleted\n", opts.projectID)
	return nil
}

func (opts *iamProjectsDeleteOpts) Confirm() error {
	if opts.confirm {
		return nil
	}
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete project '%s'?", opts.projectID),
	}
	return survey.AskOne(prompt, &opts.confirm)
}

// mcli iam project(s) delete [id] [--orgId orgId]
func IAMProjectsDeleteOpts() *cobra.Command {
	opts := new(iamProjectsDeleteOpts)
	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a project",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			opts.projectID = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	return cmd
}
