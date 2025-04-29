package logic

import (
	"context"
	"crypto/tls"
	"fmt"
	"sort"
	"strings"
	"time"

	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTlsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTlsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTlsLogic {
	return &GetTlsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTlsLogic) GetTls(req *types.GetTlsReq) (resp *types.GetTlsResp, err error) {
	sort.Slice(req.Domains, func(i, j int) bool { return req.Domains[i] > req.Domains[j] })
	domains := strings.Join(req.Domains, ",")
	mapDomain := l.analyseGetTheUpdateTime(req.Domains)
	if mapDomain == nil {
		return nil, fmt.Errorf("domains %s not found", domains)
	}

	newCreate := func() {}
	value, err := l.svcCtx.MyTlsCache.Take(domains, func() (any, error) {
		getTlsResource, err := l.svcCtx.LegoSerivce.GetTlsByDomain(req.Domains, mapDomain["DNS_CHALLENGE"], mapDomain)
		if err != nil {
			return nil, err
		}
		newCreate = func() {
			newCertificate, err := tls.X509KeyPair(getTlsResource.Certificate, getTlsResource.PrivateKey)
			if err != nil {
				l.Logger.Error("tls X509 key pair", err)
				return
			}
			updateTime := l.getTheUpdateTime(newCertificate)
			l.svcCtx.MyTlsCache.SetWithExpire(domains, getTlsResource, time.Until(updateTime))
		}
		return getTlsResource, nil
	})
	if err != nil {
		return nil, err
	}
	newCreate()

	resource, ok := value.(certificate.Resource)
	if !ok {
		return nil, fmt.Errorf("It's not ok for type string")
	}

	resp = &types.GetTlsResp{
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

func (l *GetTlsLogic) analyseGetTheUpdateTime(domains []string) map[string]string {
	domainMap := l.svcCtx.Config.LegoConf.DomainMap
	var mapDomain map[string]string
	for key, _ := range domainMap {
		for i := 0; i < len(domains); i++ {
			if strings.HasSuffix(domains[i], key) {
				mapDomain = domainMap[key]
			} else {
				mapDomain = nil
				break
			}
		}
		if mapDomain != nil {
			break
		}
	}
	return mapDomain
}

func (l *GetTlsLogic) getTheUpdateTime(cert tls.Certificate) time.Time {
	if cert.Leaf == nil {
		return time.Now()
	}
	return cert.Leaf.NotAfter.Add(-7 * 24 * time.Hour)
}
