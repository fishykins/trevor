package core

import "fmt"

const (
	color_none   string = "\033[0m"
	color_red    string = "\033[31m"
	color_green  string = "\033[32m"
	color_yellow string = "\033[33m"
	color_blue   string = "\033[34m"
	color_white  string = "\033[37m"
)

func Log(a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Println(append([]interface{}{color_white + trevor.name + color_yellow + " Log" + color_none + ":"}, a...)...)
		}
	} else {
		fmt.Println(a...)
	}
}

func Logf(format string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Printf("%s"+": "+format, append([]interface{}{color_white + trevor.name + color_yellow + " Log" + color_none}, a...)...)
		}
	} else {
		fmt.Printf(format, a...)
	}
}

func Logp(format string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Printf(format, a...)
		}
	} else {
		fmt.Printf(format, a...)
	}
}

func Error(a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Println(append([]interface{}{trevor.name + color_red + " Error:" + color_none}, a...)...)
		}
	} else {
		fmt.Println(a...)
	}
}
