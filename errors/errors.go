//Package errors implements error utilities
package errors

//Ck panics if err != nil
func Ck(err error) {
	if err != nil {
		panic(err)
	}
}
