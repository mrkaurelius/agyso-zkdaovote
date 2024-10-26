package zk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// const daovotersPath = "/home/mrk/devel/agyso/aligned24/agyso-daovote/sdk/daovote-rs/target/debug/daovote-rs"
const daovotersPath = "daovote-rs"

// Seriaslise and return calldata
func ExecAgysoDaoVoteRs() (calldata map[string]interface{}, err error) {

	out, err := exec.Command(daovotersPath).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("daovote-rs exec error")
	}
	fmt.Printf("daovote-rs output %s\n", string(out))

	callDataFile, err := os.Open(filepath.Join(proofBasePath, "calldata.json"))
	if err != nil {
		return nil, err
	}
	defer callDataFile.Close()

	callDataBytes, err := io.ReadAll(callDataFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(callDataBytes, &calldata); err != nil {
		return nil, err
	}

	return calldata, nil
}

func GetCallData() (calldata map[string]interface{}, err error) {
	callDataFile, err := os.Open(filepath.Join(proofBasePath, "calldata.json"))
	if err != nil {
		return nil, err
	}
	defer callDataFile.Close()

	callDataBytes, err := io.ReadAll(callDataFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(callDataBytes, &calldata); err != nil {
		return nil, err
	}

	return calldata, nil

}
