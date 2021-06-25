package client

import (
	pb "github.com/goblin2018/metcd/api/etcdserverpb"
)

type (
	Member                pb.Member
	MemberListResponse    pb.MemberListResponse
	MemberAddResponse     pb.MemberAddResponse
	MemberRemoveResponse  pb.MemberRemoveResponse
	MemberUpdateResponse  pb.MemberUpdateResponse
	MemberPromoteResponse pb.MemberPromoteResponse
)
