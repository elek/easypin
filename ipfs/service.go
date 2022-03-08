package ipfs

import (
	"context"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/multiformats/go-multiaddr"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
)

type Config struct {
	Address string `help:"IPFS multiaddress of the IPFS node" releaseDefault:"" devDefault:""`
}

// Service for IPFS related operations.
//
// architecture: Service
type Service struct {
	log    *zap.Logger
	string core.Multiaddr
	api    *httpapi.HttpApi
}

func NewService(log *zap.Logger, address string) (*Service, error) {
	ma, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	api, err := httpapi.NewApi(ma)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &Service{
		log: log,
		api: api,
	}, nil
}

type Peer struct {
	ID      string
	Address string
}

func (s *Service) GetPeers(ctx context.Context) ([]Peer, error) {
	res := make([]Peer, 0)
	swarmPeers, err := s.api.Swarm().Peers(ctx)
	if err != nil {
		return res, errs.Wrap(err)
	}
	for _, p := range swarmPeers {
		res = append(res, Peer{
			ID:      p.ID().Pretty(),
			Address: p.Address().String(),
		})
	}
	return res, nil
}
