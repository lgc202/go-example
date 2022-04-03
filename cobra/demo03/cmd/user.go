package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	name string
	list []string

	// userCmd 父命令
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "用户操作",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("用户列表: ", list)
		},
	}

	// 添加用户子命令
	addUserCmd = &cobra.Command{
		Use:   "add",
		Short: "添加用户: user add --name=?",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("添加用户: ", name)
		},
	}

	// 删除用户子命令
	delUserCmd = &cobra.Command{
		Use:   "del",
		Short: "删除用户: user del --name=?",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("删除用户: ", name)
		},
	}
)

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(addUserCmd)
	userCmd.AddCommand(delUserCmd)

	// 父命令接收持久标志
	userCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "用户名")

	// 父命令接收本地标志
	userCmd.Flags().StringSliceVarP(&list, "list", "l", []string{}, "用户列表")
}
