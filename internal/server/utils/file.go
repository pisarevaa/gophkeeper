package utils

import (
	"bytes"
	"io"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

// Чтение файла в MultipartForm
func ReadBinaryData(r *http.Request) (model.UploadedFile, error) {
	var newFile model.UploadedFile
	for _, fheaders := range r.MultipartForm.File {
		for _, headers := range fheaders {
			file, err := headers.Open()
			if err != nil {
				return newFile, err
			}
			defer file.Close()

			// detect contentType

			buff := make([]byte, 512)

			file.Read(buff)
			file.Seek(0, 0)
			contentType := http.DetectContentType(buff)
			newFile.ContentType = contentType

			// get file size

			var sizeBuff bytes.Buffer
			fileSize, err := sizeBuff.ReadFrom(file)
			if err != nil {
				return newFile, err
			}
			file.Seek(0, 0)
			newFile.Size = fileSize
			newFile.FileName = headers.Filename
			contentBuf := bytes.NewBuffer(nil)
			if _, err := io.Copy(contentBuf, file); err != nil {
				return newFile, err
			}
			newFile.FileContent = contentBuf.String()
		}
	}
	return newFile, nil
}
