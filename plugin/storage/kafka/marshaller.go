/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 17:09
 */
package kafka

import "github.com/Lee-xy-z/recommend/pkg/model"

// Marshaller encodes a recommend into a byte array to be sent to Kafka
type Marshaller interface {
	Marshal(info *model.Information)
}
