package main

import (
	"net/http"
)

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)

}

// func (app *application) decodePostForm(r *http.Request, dst any) error {
// 	err := r.ParseForm()
// 	if err != nil {
// 		return err
// 	}
// 	err = app.formDecoder.Decode(dst, r.PostForm)
// 	if err != nil {
// 		var invalidDecodeError *form.InvalidDecoderError
// 		if errors.As(err, &invalidDecodeError) {
// 			panic(err)
// 		}
// 		return err
// 	}
// 	return nil
// }
