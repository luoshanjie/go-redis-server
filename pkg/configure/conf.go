package configure

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ServerProperties struct {
	Bind string `cfg:"bind"`
	Port int    `cfg:"port"`
}

func NewServerProperties(configFileName string) *ServerProperties {
	f, err := os.Open(configFileName)
	if err != nil {
		panic(err)
	}
	return parse(f)
}

// ------------------------------------------------------------------------------
func parse(src io.Reader) *ServerProperties {
	config := &ServerProperties{}
	rawMap := make(map[string]string)
	scanner := bufio.NewScanner(src)

	// 解析配置
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == '#' {
			continue
		}
		pivot := strings.IndexAny(line, " ")
		if pivot > 0 && pivot < len(line)-1 {
			key := line[0:pivot]
			value := strings.Trim(line[pivot+1:], " ")
			rawMap[strings.ToLower(key)] = value
		}
	}

	// 绑定属性
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	n := t.Elem().NumField()
	for i := 0; i < n; i++ {
		field := t.Elem().Field(i)
		fieldVal := v.Elem().Field(i)
		key, ok := field.Tag.Lookup("cfg")
		if !ok {
			key = field.Name
		}
		value, ok := rawMap[strings.ToLower(key)]
		if ok {
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int:
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fieldVal.SetInt(intValue)
				}
			case reflect.Bool:
				boolValue := false
				if value == "yes" {
					boolValue = true
				}
				fieldVal.SetBool(boolValue)
			}
		}
	}

	return config
}
