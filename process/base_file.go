package process

import "os"

func storeFile(fileName string, content string) {
	exists := fileExists(fileName)

	var f *os.File
	var err error
	if exists {
		os.Remove(fileName)
	}
	f, err = os.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
