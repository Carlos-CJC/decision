package main

import (
	"errors"
	"math/rand/v2"
)

var errNoOptions = errors.New("没有可选的选项")

// pick 从候选项中等概率返回一个。
func pick(options []string) (string, error) {
	if len(options) == 0 {
		return "", errNoOptions
	}
	return options[rand.IntN(len(options))], nil
}
