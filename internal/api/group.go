// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/openimsdk/protocol/group"
	"github.com/openimsdk/tools/a2r"
)

type GroupApi struct {
	Client group.GroupClient
}

func NewGroupApi(client group.GroupClient) GroupApi {
	return GroupApi{client}
}

func (o *GroupApi) CreateGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.CreateGroup, o.Client)
}

func (o *GroupApi) SetGroupInfo(c *gin.Context) {
	a2r.Call(c, group.GroupClient.SetGroupInfo, o.Client)
}

func (o *GroupApi) SetGroupInfoEx(c *gin.Context) {
	a2r.Call(c, group.GroupClient.SetGroupInfoEx, o.Client)
}

func (o *GroupApi) JoinGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.JoinGroup, o.Client)
}

func (o *GroupApi) QuitGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.QuitGroup, o.Client)
}

func (o *GroupApi) ApplicationGroupResponse(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GroupApplicationResponse, o.Client)
}

func (o *GroupApi) TransferGroupOwner(c *gin.Context) {
	a2r.Call(c, group.GroupClient.TransferGroupOwner, o.Client)
}

func (o *GroupApi) GetRecvGroupApplicationList(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupApplicationList, o.Client)
}

func (o *GroupApi) GetUserReqGroupApplicationList(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetUserReqApplicationList, o.Client)
}

func (o *GroupApi) GetGroupUsersReqApplicationList(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupUsersReqApplicationList, o.Client)
}

func (o *GroupApi) GetSpecifiedUserGroupRequestInfo(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetSpecifiedUserGroupRequestInfo, o.Client)
}

func (o *GroupApi) GetGroupsInfo(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupsInfo, o.Client)
	//a2r.Call(c, group.GroupClient.GetGroupsInfo, o.Client, c, a2r.NewNilReplaceOption(group.GroupClient.GetGroupsInfo))
}

// @Summary		踢出群组成员
// @Description	踢出群组成员
// @Tags			group
// @Id	kickGroupMember
// @Accept			json
// @Produce		json
// @Param			data	body		group.KickGroupMemberReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.KickGroupMemberResp}
// @Router			/group/kick_group [post]
func (o *GroupApi) KickGroupMember(c *gin.Context) {
	a2r.Call(c, group.GroupClient.KickGroupMember, o.Client)
}

func (o *GroupApi) GetGroupMembersInfo(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupMembersInfo, o.Client)
	//a2r.Call(c, group.GroupClient.GetGroupMembersInfo, o.Client, c, a2r.NewNilReplaceOption(group.GroupClient.GetGroupMembersInfo))
}

// @Summary		获取群组成员列表
// @Description	获取群组成员列表
// @Tags			group
// @Id	getGroupMemberList
// @Accept			json
// @Produce		json
// @Param			data	body		group.GetGroupMemberListReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.GetGroupMemberListResp}
// @Router			/group/get_group_member_list [post]
func (o *GroupApi) GetGroupMemberList(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupMemberList, o.Client)
}

func (o *GroupApi) InviteUserToGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.InviteUserToGroup, o.Client)
}

func (o *GroupApi) GetJoinedGroupList(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetJoinedGroupList, o.Client)
}

// @Summary		解散群组
// @Description	解散群组
// @Tags			group
// @Id	dismissGroup
// @Accept			json
// @Produce		json
// @Param			data	body		group.DismissGroupReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.DismissGroupResp}
// @Router			/group/dismiss_group [post]
func (o *GroupApi) DismissGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.DismissGroup, o.Client)
}

func (o *GroupApi) MuteGroupMember(c *gin.Context) {
	a2r.Call(c, group.GroupClient.MuteGroupMember, o.Client)
}

func (o *GroupApi) CancelMuteGroupMember(c *gin.Context) {
	a2r.Call(c, group.GroupClient.CancelMuteGroupMember, o.Client)
}

// @Summary		禁言群组
// @Description	禁言群组
// @Tags			group
// @Id	muteGroup
// @Accept			json
// @Produce		json
// @Param			data	body		group.MuteGroupReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.MuteGroupResp}
// @Router			/group/mute_group [post]
func (o *GroupApi) MuteGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.MuteGroup, o.Client)
}

// @Summary		取消禁言群组
// @Description	取消禁言群组
// @Tags			group
// @Id	cancelMuteGroup
// @Accept			json
// @Produce		json
// @Param			data	body		group.CancelMuteGroupReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.CancelMuteGroupResp}
// @Router			/group/cancel_mute_group [post]
func (o *GroupApi) CancelMuteGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.CancelMuteGroup, o.Client)
}

func (o *GroupApi) SetGroupMemberInfo(c *gin.Context) {
	a2r.Call(c, group.GroupClient.SetGroupMemberInfo, o.Client)
}

func (o *GroupApi) GetGroupAbstractInfo(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupAbstractInfo, o.Client)
}

// func (g *Group) SetGroupMemberNickname(c *gin.Context) {
//	a2r.Call(c, group.GroupClient.SetGroupMemberNickname, g.userClient)
//}
//
// func (g *Group) GetGroupAllMemberList(c *gin.Context) {
//	a2r.Call(c, group.GroupClient.GetGroupAllMember, g.userClient)
//}

// @Summary		获取群组统计
// @Description	获取群组统计
// @Tags			statistics
// @Id	groupCreateCount
// @Accept			json
// @Produce		json
// @Param			data	body		group.GroupCreateCountReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.GroupCreateCountResp}
// @Router			/statistics/group/create [post]
func (o *GroupApi) GroupCreateCount(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GroupCreateCount, o.Client)
}

// @Summary		获取群组列表
// @Description	获取群组列表
// @Tags			group
// @Id	getGroups
// @Accept			json
// @Produce		json
// @Param			data	body		group.GetGroupsReq	true	"请求参数"
// @Success		200		{object}	apiresp.ApiResponse{data=group.GetGroupsResp}
// @Router			/group/get_groups [post]
func (o *GroupApi) GetGroups(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroups, o.Client)
}

func (o *GroupApi) GetGroupMemberUserIDs(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetGroupMemberUserIDs, o.Client)
}

func (o *GroupApi) GetIncrementalJoinGroup(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetIncrementalJoinGroup, o.Client)
}

func (o *GroupApi) GetIncrementalGroupMember(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetIncrementalGroupMember, o.Client)
}

func (o *GroupApi) GetIncrementalGroupMemberBatch(c *gin.Context) {
	a2r.Call(c, group.GroupClient.BatchGetIncrementalGroupMember, o.Client)
}

func (o *GroupApi) GetFullGroupMemberUserIDs(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetFullGroupMemberUserIDs, o.Client)
}

func (o *GroupApi) GetFullJoinGroupIDs(c *gin.Context) {
	a2r.Call(c, group.GroupClient.GetFullJoinGroupIDs, o.Client)
}
