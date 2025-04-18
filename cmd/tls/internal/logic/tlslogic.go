package logic

import (
	"context"

	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TlsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTlsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TlsLogic {
	return &TlsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TlsLogic) Tls(req *types.Request) (resp *types.Response, err error) {
	mapDomain := l.svcCtx.Config.LegoConf.DomainMap[req.Domain]
	resource, err := l.svcCtx.LegoSerivce.GetTlsByDomain(req.Domain, mapDomain["DNS_CHALLENGE"], mapDomain)
	if err != nil {
		return nil, err
	}
	resp = &types.Response{
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
