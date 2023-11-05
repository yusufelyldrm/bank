package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
)

func TestMain(m *testing.M) {
	var err error

	dockerContainerNameCmd := exec.Command("docker", "ps", "--format", "{{.Names}}")
	dockerContainerNameOutput, err := dockerContainerNameCmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to execute 'docker ps' command:", err)
	}
	dockerContainerName := strings.TrimSpace(string(dockerContainerNameOutput))
	fmt.Println("Docker container name:", dockerContainerName)

	cmd := exec.Command("docker", "inspect", "-f", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", dockerContainerName)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to execute 'docker inspect' command:", err)
	}

	dbSourceIP := strings.TrimSpace(string(cmdOutput))
	if dbSourceIP == "" {
		log.Fatal("Failed to get PostgreSQL container IP")
	}

	dbSource := "postgresql://root:secret@" + dbSourceIP + ":5432/simple_bank?sslmode=disable"

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
