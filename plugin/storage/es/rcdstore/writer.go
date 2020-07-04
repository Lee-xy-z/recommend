/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:10
 */
package rcdstore

import "github.com/Lee-xy-z/recommend/pkg/model"

type RcdWriter struct {
}

// NewRcdWriter creates a new RcdWriter for use
func NewRcdWriter() *RcdWriter {
	return nil
}

// WriteRcd writes a rcd:operation in ElasticSearch
func (r *RcdWriter) WriteRcd(rcd *model.Information) error {
	return nil
}
