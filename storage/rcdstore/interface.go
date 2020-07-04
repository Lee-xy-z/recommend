/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:11
 */
package rcdstore

import "github.com/Lee-xy-z/recommend/pkg/model"

// Writer writes informations to storage.
type Writer interface {
	WriteRcd(rcd *model.Information) error
}

// Reader finds and loads informations and other data from storage.
type Reader interface {
}
