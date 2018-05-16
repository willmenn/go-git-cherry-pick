package main

import (
	"fmt"
	"os/exec"
	"strings"
	"regexp"
	"github.com/labstack/echo"
	"net/http"
)

type (
	param struct {
		GitUrl     string `json:"gitUrl"`
		Regex      string `json:"regex"`
		BranchName string `json:"branchName"`
	}
)

func main() {
	e := echo.New()

	e.POST("/cherry-pick", cherryPick)

	e.Logger.Fatal(e.Start(":1323"))

}

func cherryPick(c echo.Context) error {
	p := new(param)
	if err := c.Bind(p); err != nil {
		return err
	}

	gitClone(p.GitUrl)

	logSplit := getGitLog()

	m, hashes := createArrayOfHashAndMapOfHashAndCommits(logSplit)

	hash := getFirstHashForBranchCut(hashes, m, p.Regex)

	fmt.Println(" ----- ")

	fmt.Println(hash)

	createBranch(hash, p.BranchName)

	fmt.Println(" ----- ")

	cherryPickOnlyCommitsThatDoesNotMatchRegex(hashes, m, hash, p.Regex)

	printGitLog()
	//deleteDir(err)

	return c.JSON(http.StatusOK, "Done!")
}

func cherryPickOnlyCommitsThatDoesNotMatchRegex(hashes []string, m map[string]string, hash string, regex string) {
	var b1 bool
	for _, h := range hashes {
		b, err4 := regexp.MatchString(regex, m[h])
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

func gitClone(url string) {
	out, err := exec.Command("git", "clone", url, "temp").Output()
	if err != nil {
		fmt.Print(err)
		fmt.Print(out)
	}
}

//"Feature.*"
func getFirstHashForBranchCut(hashes []string, m map[string]string, regex string) string {
	var hash string
	var bo bool
	for _, h := range hashes {
		if (bo) {
			hash = h
			bo = false
		}
		fmt.Println(h + " - " + m[h])
		b, err := regexp.MatchString(regex, m[h])
		if b && err == nil {
			fmt.Println(h + " - " + m[h])
			bo = b

		}
	}
	return hash
}

func createBranch(hash string, branchName string) {
	cmd := exec.Command("git", "checkout", "-b", branchName, string(hash))
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
