package rpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spkg/zipfs"

	"github.com/tokentransfer/node/conf"
	"github.com/tokentransfer/node/consensus"
	"github.com/tokentransfer/node/util"

	libcore "github.com/tokentransfer/interfaces/core"
)

type RPCService struct {
	config *conf.Config
	node   *consensus.Node
}

func NewRPCService(n *consensus.Node) *RPCService {
	rpc := &RPCService{
		node: n,
	}
	return rpc
}

func (service *RPCService) Init(c libcore.Config) error {
	service.config = c.(*conf.Config)
	return nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to rpc service.")
}

func wrapResult(id interface{}, result interface{}, err error) interface{} {
	ret := map[string]interface{}{
		"id":      id,
		"jsonrpc": "2.0",
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["result"] = result
	}
	return ret
}

func writeResult(w http.ResponseWriter, id interface{}, result interface{}, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(wrapResult(id, result, err)); err != nil {
		log.Println(err)
	}
}

func (service *RPCService) rpcService(w http.ResponseWriter, r *http.Request) {
	m := &map[string]interface{}{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(m); err != nil {
		writeResult(w, 0, nil, err)
	} else {
		params := util.ToArray(m, "params")
		id := util.AsUint64(m, "id")
		method := util.ToString(m, "method")

		log.Println("rpc", id, r.RequestURI, method, len(params))
		result, err := service.node.Call(method, params)
		if err != nil {
			writeResult(w, id, nil, err)
		} else {
			writeResult(w, id, result, nil)
		}
		log.Println("response", id, err)
	}
}

func (service *RPCService) getRPCAddress() string {
	rpcAddress := service.config.GetRPCAddress()
	if rpcAddress == "localhost" || rpcAddress == "127.0.0.1" {
		return rpcAddress
	}
	return "0.0.0.0"
}

func (service *RPCService) Start() error {
	rpcAddress := service.getRPCAddress()
	rpcPort := service.config.GetRPCPort()

	address := fmt.Sprintf("%s:%d", rpcAddress, rpcPort)

	// http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/", zipfs.FileServerWith(service.node.LoadPage).ServeHTTP)
	http.HandleFunc("/v1/jsonrpc", service.rpcService)
	log.Println("rpc server started on", address)
	http.ListenAndServe(address, nil)
	return nil
}
