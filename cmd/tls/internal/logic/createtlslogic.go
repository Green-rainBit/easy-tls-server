package logic

import (
	"context"

	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTlsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTlsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTlsLogic {
	return &CreateTlsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTlsLogic) CreateTls(req *types.PostCreateTlsReq) (resp *types.PostCreateTlsResp, err error) {
	resource, err := l.svcCtx.LegoSerivce.GetTlsByDomain(req.Domains, req.DnsChallenge, req.DnsConfig)
	if err != nil {
		return nil, err
	}
	resp = &types.PostCreateTlsResp{
		Domain:            resource.Domain,
		CertURL:           resource.CertURL,
		PrivateKey:        resource.PrivateKey,
		CertStableURL:     resource.CertStableURL,
		Certificate:       resource.Certificate,
		IssuerCertificate: resource.IssuerCertificate,
		CSR:               resource.CSR,
	}
	return resp, nil
}
