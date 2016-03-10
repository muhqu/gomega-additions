package worddiff

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func wordDiff(actual, expected string) (coloredDiff string, err error) {
	return gitDiff(actual, expected, "--word-diff=color")
}

func hasAnyPrefix(actual string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(actual, prefix) {
			return true
		}
	}
	return false
}

func gitDiff(actual, expected string, args ...string) (coloredDiff string, err error) {

	tmpFileActual, err := tempFileWithContents("", "actual", []byte(actual))
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFileActual.Name()) // clean up

	tmpFileExpected, err := tempFileWithContents("", "expected", []byte(expected))
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFileExpected.Name()) // clean up

	args = append(append([]string{"diff"}, args...), "--ignore-space-at-eol", "--",
		tmpFileActual.Name(), tmpFileExpected.Name())

	cmd := exec.Command("git", args...)
	b, _ := cmd.CombinedOutput()

	buf := new(bytes.Buffer)

	scanner := bufio.NewScanner(bytes.NewReader(b))
	var lineCnt = 0
	for scanner.Scan() {
		lineCnt++
		line := scanner.Text()
		if (lineCnt == 1 && !hasAnyPrefix(line, "diff ", "\x1b[1mdiff ")) ||
			(lineCnt == 2 && !hasAnyPrefix(line, "index ", "\x1b[1mindex ")) ||
			(lineCnt == 3 && !hasAnyPrefix(line, "--- ", "\x1b[1m--- ")) ||
			(lineCnt == 4 && !hasAnyPrefix(line, "+++ ", "\x1b[1m+++ ")) {
			return "", fmt.Errorf("unexpected %d. line of diff: \n%q", lineCnt, line)
		}
		if lineCnt > 4 {
			fmt.Fprintln(buf, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "\x1b[0m" + buf.String() + "\x1b[0m", nil
}

func tempFileWithContents(dir, file string, content []byte) (*os.File, error) {
	tmpfile, err := ioutil.TempFile(dir, file)
	if err != nil {
		return nil, err
	}
	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}
