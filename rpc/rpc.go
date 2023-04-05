package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caivega/glog"
	"github.com/spkg/zipfs"

	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/consensus"
	"github.com/tokentransfer/node/util"

	libcore "github.com/tokentransfer/interfaces/core"
)

type RPCService struct {
	config *config.Config
	node   *consensus.Node

	testMode bool
}

func NewRPCService(n *consensus.Node) *RPCService {
	rpc := &RPCService{
		node: n,
	}
	return rpc
}

func (service *RPCService) Init(c libcore.Config) error {
	service.config = c.(*config.Config)
	service.testMode = (service.config.GetMode() == "test")
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

func (service *RPCService) writeResult(w http.ResponseWriter, id interface{}, result interface{}, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if service.testMode {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(wrapResult(id, result, err)); err != nil {
		glog.Error(err)
	}
}

func (service *RPCService) rpcService(w http.ResponseWriter, r *http.Request) {
	m := &map[string]interface{}{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(m); err != nil {
		service.writeResult(w, 0, nil, err)
	} else {
		params := util.ToArray(m, "params")
		id := util.AsUint64(m, "id")
		method := util.ToString(m, "method")

		glog.Infoln("rpc", id, r.RequestURI, method, len(params))
		result, err := service.node.Call(method, params)
		if err != nil {
			service.writeResult(w, id, nil, err)
		} else {
			service.writeResult(w, id, result, nil)
		}
		glog.Infoln("response", id, err)
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
	http.HandleFunc("/v1/jsonrpc", service.rpcService)

	loadMap := make(map[string]func(name string) (*zipfs.FileSystem, error))
	loadMap["page"] = service.node.LoadPageByAddress
	loadMap["lib"] = service.node.LoadPageByName
	http.Handle("/", zipfs.FileServerWith(loadMap))
	if service.testMode {
		glog.Infof("rpc server started on %s, in %s mode\n", address, service.config.GetMode())
	} else {
		glog.Infoln("rpc server started on", address)
	}
	http.ListenAndServe(address, nil)
	return nil
}
