// Package chooser 提供从候选项中等概率选择的能力。
package chooser

import (
	"errors"
	"math/rand/v2"
)

// ErrNoOptions 表示没有可选择的候选项。
var ErrNoOptions = errors.New("没有可选的选项")

// Pick 从候选项中等概率返回一个。
func Pick(options []string) (string, error) {
	if len(options) == 0 {
		return "", ErrNoOptions
	}
	return options[rand.IntN(len(options))], nil
}
