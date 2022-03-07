package child

import "github.com/yamamoto-febc/module-parent/pkg/bar"

func FooBar() string {
	return "foo" + bar.Bar()
}
