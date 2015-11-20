package main

import "log"
import "os"
import "strconv"
import "io/ioutil"
import "path/filepath"
import "gopkg.in/yaml.v2"

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: " + os.Args[0] + " [config/storm.yml]")
	}
	filename, _ := filepath.Abs(os.Args[1])
	log.Println("Reading storm.yml file at: " + filename)
	yamlStr, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Saving a backup copy in: " + filename + "-orig")
	_ = ioutil.WriteFile(filename + "-orig", yamlStr, 0644)

	log.Println("Parsing YAML...")
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(yamlStr), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	delete(m, "storm.zookeeper.servers")

	zookeeperServers := []string{}

	serverNumber := 1
	stop := false
	for serverNumber < 15 && !stop {
		hostEnvVar := "ZK_SERVER_" + strconv.Itoa(serverNumber) + "_SERVICE_HOST"
		log.Println("Checking for Zookeeper host in environment variable: " + hostEnvVar)
		hostEnv := os.Getenv(hostEnvVar)
		if len(hostEnv) == 0 {
			log.Println("Not Found. Stopping here")
			stop = true
		} else {
			zookeeperServers = append(zookeeperServers, hostEnv)
			log.Printf("Found server: %v\n", hostEnv)
		}
		serverNumber += 1
	}

	if len(zookeeperServers) > 0 {
		m["storm.zookeeper.servers"] = zookeeperServers
	} else {
		log.Println("WARNING - No Zookeeper servers found. Keeping as the default localhost")
	}
	
	log.Println("Trying to save modified config")
	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

