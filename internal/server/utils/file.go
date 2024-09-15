package utils

import (
	"bytes"
	"io"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

// Чтение файла в MultipartForm.
func ReadBinaryData(r *http.Request) (model.UploadedFile, error) {
	var newFile model.UploadedFile
	for _, fheaders := range r.MultipartForm.File {
		for _, headers := range fheaders {
			file, err := headers.Open()
			if err != nil {
				return newFile, err
			}
			defer file.Close()

			var sizeBuff bytes.Buffer
			fileSize, err := sizeBuff.ReadFrom(file)
			if err != nil {
				return newFile, err
			}
			if _, errSeek := file.Seek(0, 0); errSeek != nil {
				return newFile, errSeek
			}
			newFile.Size = fileSize
			newFile.File = file
			newFile.FileName = headers.Filename
			contentBuf := bytes.NewBuffer(nil)
			if _, errCopy := io.Copy(contentBuf, file); errCopy != nil {
				return newFile, errCopy
			}
			newFile.FileContent = contentBuf.String()
			newFile.Data = contentBuf.Bytes()
			newFile.ContentType = http.DetectContentType(contentBuf.Bytes())
		}
	}
	return newFile, nil
}
