package main

import (
	"GoDisc/godisc"
	"fmt"
)

func main() {
	newChan, _ := godisc.GetChannel("711189062270189601", "NzE2NTI3MjE2NDU2MTA2MDM3.XtND-A.hQ59naWCjlorR_Uptb5BQ4hKdU4")
	fmt.Println(newChan.Name)
}