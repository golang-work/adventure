package support

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Bcrypt(password, comparePassword string) (string, error) {
	if comparePassword == "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}

		return string(hash), nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(comparePassword), []byte(password))
	if err != nil {
		return "", err
	} else {
		return "", nil
	}
}

func Padding(template string, vars ...string) string {
	if len(vars)%2 != 0 {
		return template
	}

	varMap := map[string]string{}
	for i := 0; i < len(vars); i += 2 {
		varMap[vars[i]] = vars[i+1]
	}

	// :abc 或 {{abc}} 相对单纯方案
	valid := regexp.MustCompile(`:(\w+)|\{\{(\w+)\}\}`)
	return valid.ReplaceAllStringFunc(template, func(matched string) string {
		mapKey := strings.TrimLeft(matched, ":")
		mapKey = strings.TrimLeft(mapKey, "{{")
		mapKey = strings.TrimRight(mapKey, "}}")
		if matched, ok := varMap[mapKey]; ok {
			return matched
		}
		return matched
	})
}
