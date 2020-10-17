/*
Copyright 2020 The arhat.dev Authors.

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

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"arhat.dev/abbot/pkg/cmd"
	"arhat.dev/abbot/pkg/version"

	// Add network device drivers
	_ "arhat.dev/abbot/pkg/driver/driveradd"
)

// TODO: implement cni arg parse
func main() {
	rand.Seed(time.Now().UnixNano())

	rootCmd := cmd.NewAbbotCmd()
	rootCmd.AddCommand(version.NewVersionCmd())

	err := rootCmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to run abbot %v: %v\n", os.Args, err)
		os.Exit(1)
	}
}
