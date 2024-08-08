package children

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	source  string
	number  int
)

// 这个是创建的最初始的命令
var rootCmd = &cobra.Command{
	Use:   "app_",
	Short: "a brief description of this app",
	Long:  `a brief description of this app`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("verbose")
		} else {
			fmt.Printf("cmd_app_ \n args: \n%v\n", args)
		}
	},
}

// 定义一个子命令，一般定义的另外的一个文件
var subCmd = &cobra.Command{
	Use:   "sub_",
	Short: "a brief description of this sub",
	Long:  `a long description of this sub`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cmd_sub_ \n args: \n%v\n", args)
	},
}

// 添加标志

// Execute 执行
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	fmt.Println("init1 add subCmd")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "启用详细输出")
	rootCmd.Flags().StringVarP(&source, "source", "s", "", "数据源")
	rootCmd.Flags().IntVarP(&number, "number", "n", 42, "一个数字")

	rootCmd.AddCommand(subCmd)
}
