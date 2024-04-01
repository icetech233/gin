package gin

import (
	"log"

	"github.com/icetech233/gin/binding"
)

// BindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func (c *Context) BindWith(obj any, b binding.Binding) error {
	log.Println(`BindWith(\"any, binding.Binding\") error is going to
	be deprecated, please check issue #662 and either use MustBindWith() if you
	want HTTP 400 to be automatically returned if any error occur, or use
	ShouldBindWith() if you need to manage the error.`)
	return c.MustBindWith(obj, b)
}
