package action

import (
	"../vars"
	"bufio"
	"fmt"
	"github.com/conero/uymas/bin"
	"os"
)

/**
 * @DATE        2019/6/6
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

type IniAction struct {
	bin.Command
}

func (a *IniAction) Run() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for input.Scan() {
		text := input.Text()
		if text == vars.CmdExit {
			fmt.Println(" 欢迎下次使用 [inigo]， 再见！祝您愉快")
			break
		}
		fmt.Println(">>", text)
		fmt.Print("$ ")
	}
}
