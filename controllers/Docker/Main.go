package Docker

import (
	"fmt"
	"os/exec"
)

func MakeContainer() (ID string) {
	out, err := exec.Command("ls", "-la").Output()
	if err != nil {
		fmt.Println("Command Exec Error.")
	}

	// 実行したコマンドの結果を出力
	fmt.Printf("ls result: \n%s", string(out))
	ID = string(out)
	return
}

/*
func DelContainer() {
}

func Run(cID string) {
}
*/
