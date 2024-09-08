package handler

import (
	"compress/gzip"
	"io"
	"net/http"
)

type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

// Расжатие данных при получении.
func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c compressReader) Read(p []byte) (int, error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}

// Мидлвар по сжати и расжатию данных.
func (s *Handler) GzipMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") == "gzip" {
			cr, err := newCompressReader(r.Body)
			if err != nil {
				s.Logger.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = cr
			defer cr.Close()
		}
		h.ServeHTTP(w, r)
	})
}

// Мидллвар по расжатию данных запроса в GZIP.
// func GzipDecodeMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Header.Get("Content-Encoding") == "gzip" {
// 			// Create a Gzip reader
// 			reader, err := gzip.NewReader(r.Body)
// 			if err != nil {
// 				http.Error(w, "Failed to create gzip reader", http.StatusInternalServerError)
// 				return
// 			}
// 			defer reader.Close()

// 			// Read the decompressed body
// 			var buf bytes.Buffer
// 			if _, err := io.Copy(&buf, reader); err != nil {
// 				http.Error(w, "Failed to read gzip body", http.StatusInternalServerError)
// 				return
// 			}

// 			// Replace the request body with the decompressed data
// 			r.Body = io.NopCloser(&buf)
// 			// Reset the Content-Length header (optional)
// 			r.ContentLength = int64(buf.Len())
// 			// Remove the Content-Encoding header
// 			delete(r.Header, "Content-Encoding")
// 		}

// 		// Call the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }
