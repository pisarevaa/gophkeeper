package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func CreateFormData(filepath string, name string) (*bytes.Buffer, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	values := map[string]io.Reader{
		"file": reader,
		"name": strings.NewReader(name),
	}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}
	w.Close()
	return &b, nil
}
