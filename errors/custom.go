package main

// type RequestError struct {
// 	StatusCode int

// 	Err error
// }

// func (r *RequestError) Error() string {
// 	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
// }

// func doRequest() error {
// 	return &RequestError{
// 		StatusCode: 503,
// 		Err:        errors.New("unavailable"),
// 	}
// }

// func main() {
// 	err := doRequest()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("success!")
// }

// type WrappedError struct {
// 	Context string
// 	Code    int
// 	Err     error
// }

// func (w *WrappedError) Error() string {
// 	return fmt.Sprintf("%d - %s: %v", w.Code, w.Context, w.Err)
// }

// func Wrap(info string, code int, err error) *WrappedError {
// 	return &WrappedError{
// 		Context: info,
// 		Code:    code,
// 		Err:     err,
// 	}
// }

// func main() {
// 	err := errors.New("boom!")
// 	err = Wrap("main", 400, err)

// 	//Exposing WrappedError methods
// 	if we, err2 := err.(*WrappedError); err2 {
// 		fmt.Println(we.Context, we.Code)
// 	}

// 	fmt.Println(err)
// }
