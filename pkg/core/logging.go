package core

import "fmt"

func Log(a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Println(append([]interface{}{trevor.name + ":"}, a...)...)
		}
	} else {
		fmt.Println(a...)
	}
}

func Logf(format string, a ...interface{}) {
	trevor := App()
	if trevor != nil {
		if trevor.Debug {
			fmt.Printf("%s: "+format, append([]interface{}{trevor.name}, a...)...)
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
