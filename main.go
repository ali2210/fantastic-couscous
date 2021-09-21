package main


import (
	dag "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-merkledag"
	format "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-block-format"
	ipld_format "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipld-format"
	"fmt"
	"context"
)

func main() {
	node := dag.NodeWithData([]byte("hello-world"))
	fmt.Println("Node: ", node)
	fmt.Printf("node type: %T\t", node)
	blocks := format.NewBlock([]byte("world"))
	fmt.Println("blocks cid: ", blocks.Cid())
	fmt.Printf("blocks type: %T\t", blocks)
	lnk, _:= ipld_format.MakeLink(node)
	fmt.Println(" Links:", lnk)
    // navi_obj := ipld_format.NavigableIPLDNode{}
	var node_getter ipld_format.NodeGetter = nil
	get_node, _ := node_getter.Get(context.Background(), node.Cid())
	fmt.Println(" node_getter", get_node)
	ilpd_node := ipld_format.NewNavigableIPLDNode(node, node_getter)
	fmt.Println(" ipld node:", ilpd_node)
	fmt.Println(" Total children:", ilpd_node.ChildTotal())
}