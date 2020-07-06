/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 11:12
 */
package memory

import "github.com/Lee-xy-z/recommend/storage/rcdstore"

type Factory struct {
	options Options
	store   *Store
}

func NewFactory() *Factory {
	return &Factory{}
}

// CreateRcdWriter creates a rcdstore.Writer
func (f *Factory) CreateRcdWriter() (rcdstore.Writer, error) {
	return f.store, nil
}
