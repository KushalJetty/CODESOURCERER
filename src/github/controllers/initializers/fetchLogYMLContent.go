package initializers

import (
	"log"

	"gopkg.in/yaml.v3"
)

// Config holds the YAML configuration fields
type Config struct {
	TestDirectory    string `yaml:"test-directory"`
	Comments         string `yaml:"comments"`
	TestingBranch    string `yaml:"testing-branch"`
	TestingFramework string `yaml:"testing-framework"`
	WaterMark        string `yaml:"water-mark"`
	RedisCaching     string `yaml:"redis-caching"`
}

// FetchAndLogYAMLContents fetches and logs specific YAML contents
func FetchAndLogYAMLContents(owner, repo, commitSHA, filePath string) error {
	// Fetch the file content from GitHub
	content, err := FetchFileContentFromGitHub(owner, repo, commitSHA, filePath)
	if err != nil {
		return err
	}

	// Parse YAML content
	var config Config
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return err
	}

	// Log the contents
	log.Printf("YAML Contents:\n")
	log.Printf("Test Directory: %s", config.TestDirectory)
	log.Printf("Comments: %s", config.Comments)
	log.Printf("Testing Branch: %s", config.TestingBranch)
	log.Printf("Testing Framework: %s", config.TestingFramework)
	log.Printf("Water Mark: %s", config.WaterMark)
	log.Printf("Redis Caching: %s", config.RedisCaching)

	return nil
}
