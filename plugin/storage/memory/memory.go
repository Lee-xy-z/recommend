/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 11:23
 */
package memory

import "github.com/Lee-xy-z/recommend/pkg/model"

// Store is an in-memory store of recommend
type Store struct {
}

// Writer writes informations to storage.

func (m *Store) WriteRcd(rcd *model.Information) error {
	return nil
}
