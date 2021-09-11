package manager

import (
	"dlmm/command"
	"dlmm/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Manager struct {
	*command.Command
	Progress
	client *http.Client
}

func (m *Manager) Init(env *command.Command) {
	m.Command = env
	m.client = &http.Client{}
}

func (m *Manager) assertEnvIsSet() {
	if m.Command == nil || m.client == nil {
		log.Fatal("please set the env for manager by using 'Init' method.")
	}
}

func (m *Manager) Download() {
	m.assertEnvIsSet()

	resp, err := http.Get(*m.Command.Url)
	util.AssertNoError(err)

	defer resp.Body.Close()

	// chnage working directory
	err = os.Chdir(*m.Dir)
	util.AssertNoError(err)

	// create file
	file, err := os.Create(*m.FileName)
	util.AssertNoError(err)

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	util.AssertNoError(err)

	fmt.Println("Download success!!")
}
