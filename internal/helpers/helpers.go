package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

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

func GenerateMapFromContext(c echo.Context) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func GetModelName(clase interface{}) string {
	tipo := reflect.TypeOf(clase)
	var className string
	if tipo != nil {
		className = fmt.Sprintf("%s", tipo)
		className = strings.Split(className, "structs.")[1]
	}

	return className
}

func GenerateJSONFromModels(elements ...interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for _, el := range elements {
		if name := GetModelName(el); name != "" {
			res[name] = el
		}
	}

	return res
}
