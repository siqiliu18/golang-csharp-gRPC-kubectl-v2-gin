package qhdata

import (
	"context"
	"fmt"
	"strconv"

	qhdata_proto "qhdataservice/proto"
)

//CallQil is rpc func
func (s *Service) CallQil(ctx context.Context, in *qhdata_proto.QILRequest) (*qhdata_proto.QILResponse, error) {
	fmt.Println("[CallQil] rpc func")
	out := qhdata_proto.QILResponse{
		Data: "QIL" + strconv.FormatInt(in.GetQilID(), 10) + "v" + strconv.FormatInt(in.GetQilVersion(), 10),
	}
	return &out, nil
}
