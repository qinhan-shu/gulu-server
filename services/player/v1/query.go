package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) Query(ctx context.Context, requestMessage interface{}) interface{} {
	req, ok := requestMessage.(*protocol.PlayerQuery)
	resp := protocol.PostResponse{Response: protocol.Response{ErrCode: errtable.OkCode}}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		resp.ErrCode = commonErrCode1
		resp.Msg = "convert Query request error"
		return resp
	}

	player, err := p.db.QueryPlayer(req)
	if err != nil {
		resp.ErrCode = commonErrCode2
		resp.Msg = "Invalid PlayerID"
		return resp
	}

	resp.Player.ID = req.ID
	resp.Player.Gold = player.Gold
	resp.Player.Diamond = player.Diamond
	resp.Player.Experience = player.Experience
	resp.Player.Strength = player.Strength

	return resp
}
