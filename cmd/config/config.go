/*
 * Copyright 2024 Jonas Kaninda
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "Goma Gateway configuration management",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		} else {
			fmt.Printf("config accepts no argument %q\n", args)
			os.Exit(1)
		}

	},
}

func init() {
	Cmd.AddCommand(InitConfigCmd)
	Cmd.AddCommand(CheckConfigCmd)
}
