package file

import "os"

func RetrieveOrCreateFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0755)

	if err != nil {
		return nil, err
	}
	return file, nil
}
