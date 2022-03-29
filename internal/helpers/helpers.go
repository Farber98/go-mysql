package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/labstack/echo"
)

func URLParamInt(c echo.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

//GetError genera un error HTTP que puede ser mostrado.
func GetError(err error) (string, string) {
	errorsFile, _ := ioutil.ReadFile("../errores_ES.json")
	var errorsMap map[string]string
	_ = json.Unmarshal([]byte(errorsFile), &errorsMap)
	errorCode := "ERROR_DEFAULT"
	errorMsg := errorsMap[errorCode]
	if err != nil {
		txt := err.Error()
		if txt[0:6] == "ERROR_" {
			if _, ok := errorsMap[err.Error()]; ok {
				errorCode = err.Error()
				errorMsg = errorsMap[err.Error()]
			} else {
				fmt.Println(err.Error())
			}
		}
	}
	return errorCode, errorMsg
}
