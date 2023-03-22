package p2p

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"chainmaker.org/chainmaker/common/v2/helper"
	protocol "chainmaker.org/chainmaker/protocol/v2"
	config "github.com/tokentransfer/node/conf"

	log "github.com/caivega/glog"
)

func loadFile(filePath string) ([]byte, error) {
	if len(filePath) > 0 {
		var absFilePath string
		var err error
		if !filepath.IsAbs(filePath) {
			absFilePath, err = filepath.Abs(filePath)
		} else {
			absFilePath = filePath
			err = nil
		}
		if err != nil {
			return nil, err
		}
		fileBytes, err := ioutil.ReadFile(absFilePath)
		if err != nil {
			return nil, err
		}
		return fileBytes, nil
	}
	return nil, errors.New("empty file path")
}

func InitNet(c *config.Config, readyC chan struct{}) (protocol.Net, error) {
	var netType protocol.NetType
	var err error
	// load net type
	provider := c.NetConfig.Provider
	emptyProvider := ""
	log.Infof("load net provider: %s", provider)
	switch strings.ToLower(provider) {
	case "libp2p":
		netType = protocol.Libp2p
	case "liquid", emptyProvider:
		netType = protocol.Liquid
	default:
		return nil, errors.New("unsupported net provider")
	}

	authType := c.AuthType
	emptyAuthType := ""

	// load tls keys and cert path
	keyPath := c.NetConfig.TLSConfig.PrivKeyFile
	if len(keyPath) > 0 {
		if !filepath.IsAbs(keyPath) {
			keyPath, err = filepath.Abs(keyPath)
			if err != nil {
				return nil, err
			}
		}
		log.Infof("load net tls key file path: %s", keyPath)
	}

	var certPath string
	var pubKeyMod bool
	switch strings.ToLower(authType) {
	case protocol.PermissionedWithKey, protocol.Public:
		pubKeyMod = true
	case protocol.PermissionedWithCert, protocol.Identity, emptyAuthType:
		pubKeyMod = false
		certPath = c.NetConfig.TLSConfig.CertFile
		if len(certPath) > 0 {
			if !filepath.IsAbs(certPath) {
				certPath, err = filepath.Abs(certPath)
				if err != nil {
					return nil, err
				}
			}
			log.Infof("load net tls cert file path: %s", certPath)
		}
	default:
		return nil, errors.New("wrong auth")
	}

	keyBytes, _ := loadFile(keyPath)
	certBytes, _ := loadFile(certPath)
	//gmtls enc key/cert
	encKeyBytes, _ := loadFile(c.NetConfig.TLSConfig.PrivEncKeyFile)
	encCertBytes, _ := loadFile(c.NetConfig.TLSConfig.CertEncFile)

	listenAddr := c.NetConfig.ListenAddr
	if len(listenAddr) == 0 {
		address := c.GetAddress()
		if address == "localhost" {
			address = "127.0.0.1"
		}
		listenAddr = fmt.Sprintf("/ip4/%s/tcp/%d", address, c.GetPort())
	}
	bootstraps := append(c.GetBootstraps(), c.NetConfig.Seeds...)

	// new net
	var netFactory NetFactory
	n, err := netFactory.NewNet(
		netType,
		WithReadySignalC(readyC),
		WithListenAddr(listenAddr),
		WithCrypto(pubKeyMod, keyBytes, certBytes, encKeyBytes, encCertBytes),
		WithPeerStreamPoolSize(c.NetConfig.PeerStreamPoolSize),
		WithMaxPeerCountAllowed(c.NetConfig.MaxPeerCountAllow),
		WithPeerEliminationStrategy(c.NetConfig.PeerEliminationStrategy),
		WithSeeds(bootstraps...),
		WithBlackAddresses(c.NetConfig.BlackList.Addresses...),
		WithBlackNodeIds(c.NetConfig.BlackList.NodeIds...),
		WithMsgCompression(c.NetConfig.DebugConfig.UseNetMsgCompression),
		WithInsecurity(c.NetConfig.DebugConfig.IsNetInsecurity),
		WithStunClient(c.NetConfig.StunClient.ListenAddr,
			c.NetConfig.StunClient.StunServerAddr,
			c.NetConfig.StunClient.NetworkType,
			c.NetConfig.StunClient.Enabled),
		WithStunServer(c.NetConfig.StunServer.Enabled,
			c.NetConfig.StunServer.TwoPublicAddress,
			c.NetConfig.StunServer.OtherStunServerAddr,
			c.NetConfig.StunServer.LocalNotifyAddr,
			c.NetConfig.StunServer.OtherNotifyAddr,
			c.NetConfig.StunServer.ListenAddr1,
			c.NetConfig.StunServer.ListenAddr2,
			c.NetConfig.StunServer.ListenAddr3,
			c.NetConfig.StunServer.ListenAddr4,
			c.NetConfig.StunServer.NetworkType),
		WithHolePunch(c.NetConfig.EnablePunch),
	)
	if err != nil {
		errMsg := fmt.Sprintf("new net failed: %s", err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	privateKey, err := asym.PrivateKeyFromPEM(keyBytes, nil)
	if err != nil {
		return nil, err
	}
	nodeId, err := helper.CreateLibp2pPeerIdWithPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	c.SetNodeId(nodeId)

	// load custom chain trust roots
	for _, chainTrustRoots := range c.NetConfig.CustomChainTrustRoots {
		roots := make([][]byte, 0, len(chainTrustRoots.TrustRoots))
		for _, r := range chainTrustRoots.TrustRoots {
			rootBytes, err2 := ioutil.ReadFile(r.Root)
			if err2 != nil {
				log.Errorf("load custom chain trust roots failed: %s", err2.Error())
				return nil, err2
			}
			roots = append(roots, rootBytes)
		}
		n.SetChainCustomTrustRoots(chainTrustRoots.ChainId, roots)
		log.Infof("set custom trust roots for chain[%s] success.", chainTrustRoots.ChainId)
	}
	return n, nil
}
