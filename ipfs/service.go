package ipfs

import (
	"context"
	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/multiformats/go-multiaddr"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"io"
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

func (s *Service) Upload(ctx context.Context, reader io.Reader) (string, error) {
	stat, err := s.api.Block().Put(ctx, reader)
	if err != nil {
		return "", err
	}
	return stat.Path().Cid().String(), nil
}

type Pinned struct {
	CID  string
	Size uint64
}

func (s *Service) Pin(ctx context.Context, CID string) (res Pinned, err error) {
	cCID, err := cid.Parse(CID)
	if err != nil {
		return res, errs.Wrap(err)
	}
	res.CID = cCID.String()

	err = s.api.Pin().Add(ctx, path.New("/ipfs/"+CID), options.Pin.Recursive(true))
	if err != nil {
		return res, errs.Wrap(err)
	}
	node, err := s.api.Dag().Get(ctx, cCID)
	if err != nil {
		return res, errs.Wrap(err)
	}

	res.Size, err = node.Size()
	if err != nil {
		return res, errs.Wrap(err)
	}

	return res, err
}

func (s *Service) IsPinned(ctx context.Context, CID string) (res bool, err error) {
	cCID, err := cid.Parse(CID)
	if err != nil {
		return res, errs.Wrap(err)
	}
	_, res, err = s.api.Pin().IsPinned(ctx, path.New("/ipfs/"+cCID.String()))
	if err != nil {
		return res, errs.Wrap(err)
	}
	return res, err
}
