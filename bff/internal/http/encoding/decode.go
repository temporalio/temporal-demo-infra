package encoding

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	status int
	msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.msg
}

func (mr *MalformedRequest) StatusCode() int {
	return mr.status
}

func DecodeJSONResponse(r *http.Response, dst interface{}) error {
	if r.Body == nil {
		// nothing to do here
		return nil
	}
	defer func() {
		if r.Body == nil {
			return
		}
		// gulp
		_ = r.Body.Close()
	}()
	return DecodeReadCloser(r.Body, dst)
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if v := r.Header.Get("Content-Type"); v != "application/json" {
		msg := "Content-Type header is not application/json"
		return &MalformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read my body %w", err)
	}
	err = r.Body.Close() //  no memory leaks sucker
	if err != nil {
		return err
	}
	err = DecodeReadCloser(io.NopCloser(bytes.NewBuffer(bodyBytes)), dst)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return err
}

func DecodeReadCloser(r io.ReadCloser, dst interface{}) error {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Body contains badly-formed JSON")
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Body must not be empty"
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Body must not be larger than 1MB"
			return &MalformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Body must only contain a single JSON object"
		return &MalformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}
