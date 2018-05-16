package main

import (
	"fmt"
	"os/exec"
	"strings"
	"regexp"
)

func main() {

	gitClone()

	logSplit := getGitLog()

	m, hashes := createArrayOfHashAndMapOfHashAndCommits(logSplit)

	hash := getFirstHashForBranchCut(hashes, m)

	fmt.Println(" ----- ")
	fmt.Println(hash)

	createBranch(hash)

	fmt.Println(" ----- ")

	cherryPickOnlyCommitsThatDoesNotMatchRegex(hashes, m, hash)

	printGitLog()

	//deleteDir(err)
}

func cherryPickOnlyCommitsThatDoesNotMatchRegex(hashes []string, m map[string]string, hash string) {
	var b1 bool
	for _, h := range hashes {
		b, err4 := regexp.MatchString("Feature.*", m[h])
		if (h == hash) {
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
}

func createArrayOfHashAndMapOfHashAndCommits(logSplit []string) (map[string]string, []string) {
	m := make(map[string]string)
	var hashes []string
	for _, l := range logSplit {
		var s = regexp.MustCompile("\\s").Split(l, 2)
		if len(s) > 1 {
			m[s[0]] = s[1]
			hashes = append(hashes, s[0])
		}
	}
	return m, hashes
}

func gitClone() {
	out, err := exec.Command("git", "clone", "https://github.com/willmenn/zshell_pygmalion.git", "temp").Output()
	if err != nil {
		fmt.Print(err)
		fmt.Print(out)
	}
}

func getFirstHashForBranchCut(hashes []string, m map[string]string) string {
	var hash string
	var bo bool
	for _, h := range hashes {
		if (bo) {
			hash = h
			bo = false
		}
		fmt.Println(h + " - " + m[h])
		b, err := regexp.MatchString("Feature.*", m[h])
		if b && err == nil {
			fmt.Println(h + " - " + m[h])
			bo = b

		}
	}
	return hash
}

func createBranch(hash string) {
	cmd := exec.Command("git", "checkout", "-b", "temp", string(hash))
	cmd.Dir = "temp"
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func printGitLog() {
	logSplit1 := getGitLog()
	for _, l := range logSplit1 {
		fmt.Println(l)
	}
}

func deleteDir() {
	out, err := exec.Command("rm", "-rf", "temp").Output()
	if err != nil {
		fmt.Print(err)
		fmt.Print(out)
	}
}

func getGitLog() ([]string) {
	cmd := exec.Command("git", "log", "--oneline", "--no-color")
	cmd.Dir = "temp"
	out3, err3 := cmd.Output()
	if err3 != nil {
		fmt.Print(err3)
	}
	log := string(out3)
	logSplit := strings.Split(log, "\n")
	return logSplit
}
