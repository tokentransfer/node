package block

import (
	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"
)

type GroupInfo struct {
	Hash libcore.Hash
	List []*DataInfo
	Map  map[string]*DataInfo
}

func (s *GroupInfo) GetHash() libcore.Hash {
	return s.Hash
}

func (s *GroupInfo) SetHash(h libcore.Hash) {
	s.Hash = h
}

func (s *GroupInfo) AddData(info *DataInfo) *GroupInfo {
	s.List = append(s.List, info)

	return s
}

func (s *GroupInfo) PutData(name string, info *DataInfo) *GroupInfo {
	if s.Map == nil {
		s.Map = make(map[string]*DataInfo)
	}
	s.Map[name] = info

	return s
}

func (s *GroupInfo) GetData(name string) *DataInfo {
	if s.Map != nil {
		return s.Map[name]
	}
	return nil
}

func (s *GroupInfo) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_GROUP_INFO {
		return util.ErrorOfInvalid("group info", "data")
	}
	info := msg.(*pb.GroupInfo)

	s.Hash = libcore.Hash(info.Hash)
	s.List = fromDataList(info.List)
	s.Map = fromDataMap(info.Map)
	return nil
}

func fromDataList(infos []*pb.DataInfo) []*DataInfo {
	if infos == nil {
		return nil
	}
	l := len(infos)
	list := make([]*DataInfo, l)
	for i, info := range infos {
		list[i] = FromDataInfo(info)
	}
	return list
}

func fromDataMap(m map[string]*pb.DataInfo) map[string]*DataInfo {
	if m == nil {
		return nil
	}
	r := make(map[string]*DataInfo)
	for k, v := range m {
		r[k] = FromDataInfo(v)
	}
	return r
}

func (s *GroupInfo) MarshalBinary() ([]byte, error) {
	return s.Raw(libcrypto.RawBinary)
}

func (s *GroupInfo) Raw(rt libcrypto.RawType) ([]byte, error) {
	info := &pb.GroupInfo{
		Hash: []byte(s.Hash),
		List: toDataList(s.List, rt),
		Map:  toDataMap(s.Map, rt),
	}
	data, err := core.Marshal(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func toDataList(infos []*DataInfo, rt libcrypto.RawType) []*pb.DataInfo {
	if infos == nil {
		return nil
	}
	l := len(infos)
	list := make([]*pb.DataInfo, l)
	for i := 0; i < l; i++ {
		list[i] = ToDataInfo(infos[i], rt)
	}
	return list
}

func toDataMap(m map[string]*DataInfo, rt libcrypto.RawType) map[string]*pb.DataInfo {
	if m == nil {
		return nil
	}
	r := make(map[string]*pb.DataInfo)
	for k, v := range m {
		r[k] = ToDataInfo(v, rt)
	}
	return r
}

type DataInfo struct {
	Hash    libcore.Hash
	Content []byte
}

func (s *DataInfo) GetHash() libcore.Hash {
	return s.Hash
}

func (s *DataInfo) SetHash(h libcore.Hash) {
	s.Hash = h
}

func (s *DataInfo) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_DATA_INFO {
		return util.ErrorOfInvalid("data info", "data")
	}
	info := msg.(*pb.DataInfo)
	s.Hash = libcore.Hash(info.Hash)
	s.Content = info.Content
	return nil
}

func (s *DataInfo) MarshalBinary() ([]byte, error) {
	info := ToDataInfo(s, libcrypto.RawBinary)
	data, err := core.Marshal(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FromDataInfo(info *pb.DataInfo) *DataInfo {
	if info != nil {
		return &DataInfo{
			Hash:    libcore.Hash(info.Hash),
			Content: info.Content,
		}
	}
	return nil
}

func ToDataInfo(info *DataInfo, rt libcrypto.RawType) *pb.DataInfo {
	if info != nil {
		switch rt {
		case libcrypto.RawBinary:
			return &pb.DataInfo{
				Hash:    []byte(info.Hash),
				Content: info.Content,
			}
		case libcrypto.RawIgnoreVariableFields:
			return &pb.DataInfo{
				Hash:    []byte(info.Hash),
				Content: info.Content,
			}
		case libcrypto.RawIgnoreSigningFields:
			return &pb.DataInfo{
				Hash:    []byte(info.Hash),
				Content: info.Content,
			}
		case libcrypto.RawIgnoreContent:
			return &pb.DataInfo{
				Hash: []byte(info.Hash),
			}
		}
	}
	return nil
}

func (s *DataInfo) Raw(rt libcrypto.RawType) ([]byte, error) {
	info := ToDataInfo(s, rt)
	data, err := core.Marshal(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type PayloadInfo struct {
	Hash  libcore.Hash
	Infos []*DataInfo
}

func (s *PayloadInfo) GetHash() libcore.Hash {
	return s.Hash
}

func (s *PayloadInfo) SetHash(h libcore.Hash) {
	s.Hash = h
}

func (s *PayloadInfo) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_PAYLOAD_INFO {
		return util.ErrorOfInvalid("payload info", "data")
	}

	info := msg.(*pb.PayloadInfo)
	infos := make([]*DataInfo, 0)
	for _, pbInfo := range info.Infos {
		dataInfo := FromDataInfo(pbInfo)
		infos = append(infos, dataInfo)
	}
	s.Infos = infos
	return nil
}

func (s *PayloadInfo) MarshalBinary() ([]byte, error) {
	return s.Raw(libcrypto.RawBinary)
}

func fromPayloadInfo(info *pb.PayloadInfo) *PayloadInfo {
	if info != nil {
		infos := make([]*DataInfo, 0)
		for _, pbInfo := range info.Infos {
			dataInfo := FromDataInfo(pbInfo)
			infos = append(infos, dataInfo)
		}
		return &PayloadInfo{
			Infos: infos,
		}
	}
	return nil
}

func toPayloadInfo(info *PayloadInfo, rt libcrypto.RawType) *pb.PayloadInfo {
	if info != nil {
		infos := make([]*pb.DataInfo, 0)
		for _, info := range info.Infos {
			pbInfo := ToDataInfo(info, rt)
			infos = append(infos, pbInfo)
		}
		return &pb.PayloadInfo{
			Infos: infos,
		}
	}
	return nil
}

func (s *PayloadInfo) Raw(rt libcrypto.RawType) ([]byte, error) {
	info := &pb.PayloadInfo{
		Infos: toDataList(s.Infos, rt),
	}
	data, err := core.Marshal(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddressToByte(a libcore.Address) ([]byte, error) {
	if a == nil {
		return nil, nil
	}
	return a.MarshalBinary()
}

func ByteToAddress(b []byte) (libcore.Address, error) {
	if b == nil {
		return nil, nil
	}
	_, a, err := as.NewAccountFromBytes(b)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func ReadState(data []byte) (libblock.State, error) {
	if len(data) == 0 {
		return nil, util.ErrorOf("empty", "data", "entry")
	}
	meta := core.GetMeta(data)
	switch meta {
	case core.CORE_ACCOUNT_STATE:
		s := &AccountState{}
		err := s.UnmarshalBinary(data)
		if err != nil {
			return nil, err
		}
		return s, nil
	default:
		return nil, util.ErrorOfUnknown("data", "state")
	}
}

func CloneState(state libblock.State) (libblock.State, error) {
	data, err := state.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s, err := ReadState(data)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func ReadTransaction(data []byte) (libblock.Transaction, error) {
	if len(data) == 0 {
		return nil, util.ErrorOfInvalid("null", "transaction data")
	}
	meta := core.GetMeta(data)
	switch meta {
	case core.CORE_TRANSACTION:
		tx := &Transaction{}
		err := tx.UnmarshalBinary(data)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, util.ErrorOfInvalid("transaction", "data")
	}
}

func GetDataInfo(info *GroupInfo) (*DataInfo, error) {
	dataContent, err := info.Raw(libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}
	dataHash, err := cs.Hash(dataContent)
	if err != nil {
		return nil, err
	}
	return &DataInfo{
		Hash:    dataHash,
		Content: dataContent,
	}, nil
}

func GetGroupInfo(info *DataInfo) (*GroupInfo, error) {
	group := &GroupInfo{}
	err := group.UnmarshalBinary(info.Content)
	if err != nil {
		return nil, err
	}

	expectInfo, err := GetDataInfo(group)
	if err != nil {
		return nil, err
	}
	if info.Hash.String() != expectInfo.Hash.String() {
		return nil, util.ErrorOfUnmatched("hash", "group info", info.Hash.String(), expectInfo.Hash.String())
	}

	return group, nil
}
