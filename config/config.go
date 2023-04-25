package config

import (
	"encoding/json"
	"os"

	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/account"
)

var as = account.NewAccountService()

type Config struct {
	// address 0xc287B1266732495Fe8c93CE3Ba631597153fdd91
	// secret 86d3350f255e5b6259d3d3a615b363f23c042971d89b7f9cb84aa7fadeeb2736
	GasAddress string `json:"gas_address"`
	gasAccount libcore.Address

	DataDir string `json:"data_dir"`
	Address string `json:"address"`
	Port    int64  `json:"port"`

	RPCAddress string `json:"rpc_address"`
	RPCPort    int64  `json:"rpc_port"`

	Mode string `json:"mode"`

	// address 0x768e56BCcb18a8622cc5BB5F6bfA6D82a255ab87
	// secret abb33a8c2bb48d3b1c2ce365685ac3b96563e6a17ebc4367cd637e33149f94ea
	Secret string `json:"secret"`

	Bootstraps []string `json:"bootstraps"`

	ChainId   string `json:"chain_id"`
	AuthType  string `json:"auth_type"`
	NodeId    string
	NetConfig NetConfig `json:"net_config"`
}

type NetConfig struct {
	Provider                string            `json:"provider"`
	ListenAddr              string            `json:"listen_address"`
	PeerStreamPoolSize      int               `json:"peer_stream_pool_size"`
	MaxPeerCountAllow       int               `json:"max_peer_count_allow"`
	PeerEliminationStrategy int               `json:"peer_elimination_strategy"`
	Seeds                   []string          `json:"seeds"`
	BlackList               BlackList         `json:"blocklist"`
	DebugConfig             DebugConfig       `json:"debug_config"`
	TLSConfig               TLSConfig         `json:"tls_config"`
	StunClient              StunClient        `json:"stun_client"`
	StunServer              StunServer        `json:"stun_server"`
	EnablePunch             bool              `json:"enable_punch"`
	CustomChainTrustRoots   []ChainTrustRoots `json:"custom_chain_trust_roots"`
}

type ChainTrustRoots struct {
	ChainId    string       `json:"chain_id"`
	TrustRoots []TrustRoots `json:"trust_roots"`
}

type TrustRoots struct {
	OrgId string `json:"org_id"`
	Root  string `json:"root"`
}

type BlackList struct {
	Addresses []string `json:"addresses"`
	NodeIds   []string `json:"node_ids"`
}

type DebugConfig struct {
	UseNetMsgCompression bool `json:"use_net_msg_compression"`
	IsNetInsecurity      bool `json:"is_net_insecurity"`
}

type StunClient struct {
	ListenAddr     string `json:"listen_address"`
	StunServerAddr string `json:"stun_server_address"`
	NetworkType    string `json:"network_type"`
	Enabled        bool   `json:"enabled"`
}

type StunServer struct {
	Enabled             bool   `json:"enabled"`
	TwoPublicAddress    bool   `json:"two_public_address"`
	OtherStunServerAddr string `json:"other_stun_server_address"`
	LocalNotifyAddr     string `json:"local_notify_address"`
	OtherNotifyAddr     string `json:"other_notify_address"`
	ListenAddr1         string `json:"listen_address_1"`
	ListenAddr2         string `json:"listen_address_2"`
	ListenAddr3         string `json:"listen_address_3"`
	ListenAddr4         string `json:"listen_address_4"`
	NetworkType         string `json:"network_type"`
}

type TLSConfig struct {
	PrivKeyFile string `json:"private_key_file"`
	CertFile    string `json:"cert_file"`

	PrivEncKeyFile string `json:"private_encode_key_file"`
	CertEncFile    string `json:"cert_encode_file"`
}

func NewConfig(configFile string) (*Config, error) {
	config := &Config{}

	cfg, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer cfg.Close()

	jsonParser := json.NewDecoder(cfg)
	err = jsonParser.Decode(config)
	if err != nil {
		return nil, err
	}

	gasAddress := config.GasAddress
	_, gasAccount, err := as.NewAccountFromAddress(gasAddress)
	if err != nil {
		return nil, err
	}
	config.gasAccount = gasAccount

	return config, nil
}

func (c *Config) GetId() string {
	return "test"
}

func (c *Config) GetNodeId() string {
	return c.NodeId
}

func (c *Config) SetNodeId(nodeId string) {
	c.NodeId = nodeId
}

func (c *Config) GetChainId() string {
	return c.ChainId
}

func (c *Config) GetType() int {
	return 0
}

func (c *Config) GetMode() string {
	return c.Mode
}

func (c *Config) SetMode(m string) {
	c.Mode = m
}

func (c *Config) GetGasAccount() libcore.Address {
	return c.gasAccount
}

func (c *Config) GetDataDir() string {
	return c.DataDir
}

func (c *Config) GetAddress() string {
	return c.Address
}

func (c *Config) GetPort() int64 {
	return c.Port
}

func (c *Config) GetSecret() string {
	return c.Secret
}

func (c *Config) GetBootstraps() []string {
	return c.Bootstraps
}

func (c *Config) GetThreadCount() uint32 {
	return 8
}

func (c *Config) GetTimeout() int64 {
	return 30
}

func (c *Config) GetRPCAddress() string {
	return c.RPCAddress
}

func (c *Config) GetRPCPort() int64 {
	return c.RPCPort
}

func (c *Config) GetSystemCode() string {
	return "TEST"
}

func (c *Config) GetBlockDuration() uint32 {
	return 10
}

func (c *Config) GetBackend() string {
	return ""
}
