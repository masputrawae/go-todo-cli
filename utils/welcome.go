package utils

import (
	"fmt"
	"sort"
	"strings"

	"github.com/masputrawae/go-todo/pkg/store"
)

var art = `
   █████████                      ███████████              █████         
  ███░░░░░███                    ░█░░░███░░░█             ░░███          
 ███     ░░░   ██████            ░   ░███  ░   ██████   ███████   ██████ 
░███          ███░░███ ██████████    ░███     ███░░███ ███░░███  ███░░███
░███    █████░███ ░███░░░░░░░░░░     ░███    ░███ ░███░███ ░███ ░███ ░███
░░███  ░░███ ░███ ░███               ░███    ░███ ░███░███ ░███ ░███ ░███
 ░░█████████ ░░██████                █████   ░░██████ ░░████████░░██████ 
  ░░░░░░░░░   ░░░░░░                ░░░░░     ░░░░░░   ░░░░░░░░  ░░░░░░  
                                                                         
=========================================================================
Github: https://github.com/masputrawae/go-todo`

func Welcome() string {
	var prList []string

	keys := make([]string, 0, len(store.Priorities))
	for k := range store.Priorities {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		prList = append(prList, fmt.Sprintf(
			"  %s %s %s %s \033[0m",
			store.Priorities[k].Color,
			k,
			store.Priorities[k].Emoji,
			store.Priorities[k].Name,
		))
	}
	return fmt.Sprintf("%s\nUse Priorities\n%s", art, strings.Join(prList, "\n"))
}
