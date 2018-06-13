// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unassignIpCmd represents the unassignIp command
var unassignIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Unassigns an IP address.",
	Long: `Example:
  packet unassign ip -a [assignment-UUID]
	`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := PacknGo.DeviceIPs.Unassign(assignmentID)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		fmt.Println("IP address unassigned successfully.")
	},
}

func init() {
	unassignCmd.AddCommand(unassignIPCmd)
	unassignIPCmd.Flags().StringVarP(&assignmentID, "assignment", "a", "", "--assignemnt or -a")
	unassignIPCmd.MarkFlagRequired("assignment")
}
