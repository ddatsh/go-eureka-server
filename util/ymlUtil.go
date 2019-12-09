package util

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func ReadYml(file string) map[interface{}]interface{} {

	var decoder *yaml.Decoder
	var contentMap = map[interface{}]interface{}{}

	if _, err := os.Stat(file); !os.IsNotExist(err) {

		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder = yaml.NewDecoder(f)

		for {

			var overrideMap map[string]interface{}

			errorReading := decoder.Decode(&overrideMap)
			if errorReading == io.EOF {
				return contentMap
			} else {
				for k, v := range overrideMap {
					contentMap[k] = v
				}
			}
		}
	}

	return contentMap
}
