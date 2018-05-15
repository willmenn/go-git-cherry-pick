package main

import (
	"fmt"
	"os/exec"
	"strings"
	"regexp"
)

func main() {

	out, err := exec.Command("git", "clone", "https://github.com/willmenn/zshell_pygmalion.git", "temp").Output()
	if err != nil {
		fmt.Print(err)
		fmt.Print(out)
	}
	//fmt.Printf("hello, world\n")

	_, logSplit := printLog()
	m := make(map[string]string)
	var hashes []string
	for _, l := range logSplit {
		var s = regexp.MustCompile("\\s").Split(l, 2)
		if len(s) > 1 {
			m[s[0]] = s[1]
			hashes = append(hashes, s[0])
		}
	}

	var hash string
	var bo bool
	for _, h := range hashes {
		if (bo) {
			hash = h
			bo = false
		}
		fmt.Println(h + " - " + m[h])
		b, err4 := regexp.MatchString("Feature.*", m[h])
		if b && err4 == nil {
			fmt.Println(h + " - " + m[h])
			bo = b

		}
	}

	fmt.Println(" ----- ")
	fmt.Println(hash)

	cmd1 := exec.Command("git", "checkout","-b","temp",string(hash))
	cmd1.Dir = "temp"
	_, err5 := cmd1.Output()
	if err5 != nil {
		fmt.Println(err5)
	}

	fmt.Println(" ----- ")

	var b1 bool
	for _, h := range hashes {
		b, err4 := regexp.MatchString("Feature.*", m[h])
		if(h == hash){
			b1 = true
		}
		if !b1 && !b && err4 == nil {
			fmt.Println(h + " - " + m[h])
			cmdC := exec.Command("git", "cherry-pick", h)
			cmdC.Dir = "temp"
			out6, err6 := cmdC.Output()
			if out6 != nil && err6 != nil {
				fmt.Println(err6)
			}

		}
	}

	_, logSplit1 := printLog()
	for _, l := range logSplit1 {
		fmt.Println(l)
	}

	//out1, err1 := exec.Command("rm","-rf","temp").Output()
	//if err != nil {
	//	fmt.Print(err1)
	//	fmt.Print(out1)
	//}
}

func printLog() (*exec.Cmd, []string) {
	cmd := exec.Command("git", "log", "--oneline", "--no-color")
	cmd.Dir = "temp"
	out3, err3 := cmd.Output()
	if err3 != nil {
		fmt.Print(err3)
	}
	log := string(out3)
	logSplit := strings.Split(log, "\n")
	return cmd, logSplit
}
