package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/echlebek/sensu-lite/types"
)

func namespaceExistsForResource(r types.MultitenantResource) clientv3.Cmp {
	key := getNamespacePath(r.GetNamespace())
	return clientv3.Compare(clientv3.Version(key), ">", 0)
}
