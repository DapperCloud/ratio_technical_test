package test

import (
	"testing"
	"runtime"
	"path"
	"path/filepath"
	"os"
	"log"
	"io"
	"bytes"
	"os/exec"
	"strings"
)

const (
	ScenariosRelativePath = "/scenarios/deterministic"
	CommandRelativePath = "/../bin/monsters"
)


func Test_Deterministic_Scenarios(t *testing.T) {
	baseDir := getBaseDir()
	scenarios := []string{}
    err := filepath.Walk(baseDir + ScenariosRelativePath, func(path string, f os.FileInfo, err error) error {
    	if err != nil {
    		panic(err)
		}
    	if f.IsDir() && path != baseDir + ScenariosRelativePath {
			scenarios = append(scenarios, path)
		}
        return nil
    })
    t.Logf("%v", scenarios)
    if err != nil {
    	panic(err)
	}

	for _, scenario := range scenarios {
		testScenario(t, baseDir + CommandRelativePath, scenario)
	}
}

func getBaseDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}

func testScenario(t *testing.T, commandPath string, scenarioDir string) {
	mapPath := scenarioDir + "/map.in"
	actualOutputPath := scenarioDir + "/actual.out"
	expectedOutputPath := scenarioDir + "/expected.out"
	scenarioName := fileNameFromPath(scenarioDir)

	cmd := exec.Command(commandPath, mapPath, "1")
	outputFile, err := os.Create(actualOutputPath)
	if err != nil {
		panic(err)
	}
	cmd.Stdout = outputFile

	err = cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			panic(err)
		}
	}

	if deepCompare(actualOutputPath, expectedOutputPath) {
		t.Logf("%v: OK", scenarioName)
	} else {
		t.Errorf("%v: KO - expected and actual outputs differ", scenarioName)
	}
}

func fileNameFromPath(path string) string {
	tokens := strings.Split(path, "/")
	return tokens[len(tokens) - 1]
}

const chunkSize = 64000

func deepCompare(file1, file2 string) bool {
    // Check file size ...

    f1, err := os.Open(file1)
    if err != nil {
        panic(err)
    }

    f2, err := os.Open(file2)
    if err != nil {
        panic(err)
    }

    for {
        b1 := make([]byte, chunkSize)
        _, err1 := f1.Read(b1)

        b2 := make([]byte, chunkSize)
        _, err2 := f2.Read(b2)

        if err1 != nil || err2 != nil {
            if err1 == io.EOF && err2 == io.EOF {
                return true
            } else if err1 == io.EOF || err2 == io.EOF {
                return false
            } else {
                log.Fatal(err1, err2)
            }
        }

        if !bytes.Equal(b1, b2) {
            return false
        }
    }
}