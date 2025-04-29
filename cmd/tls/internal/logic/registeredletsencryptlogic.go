package logic

import (
	"context"

	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisteredLetsEncryptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisteredLetsEncryptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisteredLetsEncryptLogic {
	return &RegisteredLetsEncryptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisteredLetsEncryptLogic) RegisteredLetsEncrypt(req *types.PostRegisteredLetsEncryptReq) (resp *types.PostRegisteredLetsEncryptResp, err error) {
	privateKey, err := l.svcCtx.LegoSerivce.RegisteredLetsEncrypt(req.Email)
	if err != nil {
		return nil, err
	}
	resp = &types.PostRegisteredLetsEncryptResp{
		PrivateKeyStre: privateKey,
	}
	return resp, nil
}
