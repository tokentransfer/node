package block

import (
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"
)

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
	info := toDataInfo(s, libcrypto.RawBinary)
	data, err := core.Marshal(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func fromDataInfo(info *pb.DataInfo) *DataInfo {
	if info != nil {
		return &DataInfo{
			Hash:    libcore.Hash(info.Hash),
			Content: info.Content,
		}
	}
	return nil
}

func toDataInfo(info *DataInfo, rt libcrypto.RawType) *pb.DataInfo {
	if info != nil {
		switch rt {
		case libcrypto.RawBinary:
			return &pb.DataInfo{
				Hash:    []byte(info.Hash),
				Content: info.Content,
			}
		case libcrypto.RawIgnoreVariableFields:
			return &pb.DataInfo{
				Hash: []byte(info.Hash),
			}
		case libcrypto.RawIgnoreSigningFields:
			return &pb.DataInfo{
				Hash: []byte(info.Hash),
			}
		}

	}
	return nil
}

func (s *DataInfo) Raw(ignoreSigningFields bool) ([]byte, error) {
	if ignoreSigningFields {
		info := toDataInfo(s, libcrypto.RawIgnoreSigningFields)
		data, err := core.Marshal(info)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else { //ignore variable fields
		info := toDataInfo(s, libcrypto.RawIgnoreVariableFields)
		data, err := core.Marshal(info)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
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
		dataInfo := fromDataInfo(pbInfo)
		infos = append(infos, dataInfo)
	}
	s.Infos = infos
	return nil
}

func (s *PayloadInfo) MarshalBinary() ([]byte, error) {
	infos := make([]*pb.DataInfo, 0)
	for _, info := range s.Infos {
		pbInfo := toDataInfo(info, libcrypto.RawBinary)
		infos = append(infos, pbInfo)
	}
	info := &pb.PayloadInfo{
		Infos: infos,
	}
	data, err := core.Marshal(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func fromPayloadInfo(info *pb.PayloadInfo) *PayloadInfo {
	if info != nil {
		infos := make([]*DataInfo, 0)
		for _, pbInfo := range info.Infos {
			dataInfo := fromDataInfo(pbInfo)
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
			pbInfo := toDataInfo(info, rt)
			infos = append(infos, pbInfo)
		}
		return &pb.PayloadInfo{
			Infos: infos,
		}
	}
	return nil
}

func (s *PayloadInfo) Raw(ignoreSigningFields bool) ([]byte, error) {
	if ignoreSigningFields {
		infos := make([]*pb.DataInfo, 0)
		for _, info := range s.Infos {
			pbInfo := toDataInfo(info, libcrypto.RawIgnoreSigningFields)
			infos = append(infos, pbInfo)
		}
		info := &pb.PayloadInfo{
			Infos: infos,
		}
		data, err := core.Marshal(info)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else { //ignore variable fields
		infos := make([]*pb.DataInfo, 0)
		for _, info := range s.Infos {
			pbInfo := toDataInfo(info, libcrypto.RawIgnoreVariableFields)
			infos = append(infos, pbInfo)
		}
		info := &pb.PayloadInfo{
			Infos: infos,
		}
		data, err := core.Marshal(info)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}
