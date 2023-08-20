package helper

import (
	"embed"
	"fmt"
	"path/filepath"
	e "root/check_error"
	"root/constants"
)

//go:embed *.sql
var sqlFiles embed.FS

func init() {
	fileInfos, err := sqlFiles.ReadDir(constants.FILE_SQL_NAME)
	e.ErrorHandler(err)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		filePath := filepath.Join(constants.FILE_SQL_NAME, fileInfo.Name()) //sql + / + fileInfo.Name()
		sqlContent, err := sqlFiles.ReadFile(filePath)
		e.ErrorHandler(err)

		fmt.Printf("File: %s\nContent:\n%s\n", filePath, sqlContent)
	}
}
