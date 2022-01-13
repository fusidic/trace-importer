package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jConfiguration struct {
	Uri      string
	Username string
	Password string
	Database string
}

func (nc Neo4jConfiguration) newDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.Uri, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func defaultHandler(http.ResponseWriter, *http.Request) {

}

func searcHanlerFunc(driver neo4j.Driver, database string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// TODO: finish search codes
	}
}

func main() {

	config := parseConfiguration()
	driver, err := config.newDriver()
	if err != nil {
		log.Fatal(err)
	}
	defer unsafeClose(driver)

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", defaultHandler)
	serverMux.HandleFunc("/search", searcHanlerFunc(driver, config.Database))

}

func parseConfiguration() *Neo4jConfiguration {
	database := lookupEnvOrGetDefault("NEO4J_DATABASE", "movies")
	if !strings.HasPrefix(lookupEnvOrGetDefault("NEO4j_VERSION", "4"), "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Uri:      lookupEnvOrGetDefault("NEO4J_URI", "neo4j+s://demo.neo4jlabs.com"),
		Username: lookupEnvOrGetDefault("NEO4J_USER", "alibaba_trace"),
		Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "alibaba_trace"),
		Database: database,
	}
}

func lookupEnvOrGetDefault(key, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
