package yaml

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// thanks to github.com/ilyakaznacheev/cleanenv

func ReadConfig(path string, cfg interface{}) error {
	// open the configuration file
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".yml" {
		return fmt.Errorf("file format '%s' doesn't supported by the parser", ext)
	}

	err = parseYAML(file, cfg)

	if err != nil {
		return fmt.Errorf("config file parsing error: %s", err.Error())
	}
	return nil
}

func parseYAML(r io.Reader, str interface{}) error {
	return yaml.NewDecoder(r).Decode(str)
}
