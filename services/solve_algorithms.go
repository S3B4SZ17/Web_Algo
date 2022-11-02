package services

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/S3B4SZ17/Web_Algo/management"
	"go.uber.org/zap"
)

var (
	path string = "/tmp/web_algo/"
	file string = path + "algorithm.py"
)

func CreateDir() {
	// Create directory if it doesn't exist
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			management.Log.Fatal(err.Error())
		}

	} else {
		management.Log.Info("Directory already exists")
	}
}

func SaveFile(file_content *string) {

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	// f, err := os.Create(path + file)

	if err != nil {
		management.Log.Fatal(err.Error())
	}
	defer f.Close()

	_, err2 := f.WriteString(*file_content)

	if err2 != nil {
		management.Log.Fatal(err2.Error())
	}

	management.Log.Info("Algorithm saved to path = " + file)

}

func RunFile() (string, error) {
	prg := "python3"
	arg1 := file
	arg2 := "6589"

	cmd := exec.Command(prg, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		management.Log.Error(err.Error())
		return err.Error(), err
	}

	output := strings.TrimSuffix(string(stdout), "\n")
	management.Log.Info(output)

	return output, nil
}

func CompareResult(file_content string) (bool, error) {

	CreateDir()
	SaveFile(&file_content)
	got, err := RunFile()
	want := "9856"

	if err != nil {
		management.Log.Error("Your code is bad, please fix it.")
		return false, err
	}

	if got != want {
		management.Log.Error("Didnt return expected result", zap.String("got", got), zap.String("want", want))
		err := errors.New("Didnt return expected result, got: " + got + " and want: " + want)
		return false, err
	}
	return true, nil
}

func DeleteTempFile() {
	// Removing file from the directory
	err := os.Remove(file)
	if err != nil {
		management.Log.Fatal(err.Error())
	}
}

type AlgoResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

type AlgoFile struct {
	File string `json:"file"`
}
