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

package msg

// import (
// 	"context"
// 	"encoding/json"
// 	"time"

// 	"github.com/openimsdk/open-im-server/v3/pkg/common/storage/model"

// 	"github.com/openimsdk/open-im-server/v3/pkg/authverify"
// 	"github.com/openimsdk/open-im-server/v3/pkg/common/servererrs"
// 	"github.com/openimsdk/protocol/constant"
// 	"github.com/openimsdk/protocol/msg"
// 	"github.com/openimsdk/protocol/sdkws"
// 	"github.com/openimsdk/tools/errs"
// 	"github.com/openimsdk/tools/log"
// 	"github.com/openimsdk/tools/mcontext"
// 	"github.com/openimsdk/tools/utils/datautil"
// )

// func (m *msgServer) EditMsg(ctx context.Context, req *msg.EditMsgReq) (*msg.EditMsgResp, error) {
// 	if req.UserID == "" {
// 		return nil, errs.ErrArgs.WrapMsg("user_id is empty")
// 	}
// 	if req.ConversationID == "" {
// 		return nil, errs.ErrArgs.WrapMsg("conversation_id is empty")
// 	}
// 	if req.Seq < 0 {
// 		return nil, errs.ErrArgs.WrapMsg("seq is invalid")
// 	}
// 	if req.NewContent == "" {
// 		return nil, errs.ErrArgs.WrapMsg("new_content is empty")
// 	}
// 	if err := authverify.CheckAccess(ctx, req.UserID); err != nil {
// 		return nil, err
// 	}
// 	user, err := m.UserLocalCache.GetUserInfo(ctx, req.UserID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, _, msgs, err := m.MsgDatabase.GetMsgBySeqs(ctx, req.UserID, req.ConversationID, []int64{req.Seq})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(msgs) == 0 || msgs[0] == nil {
// 		return nil, errs.ErrRecordNotFound.WrapMsg("msg not found")
// 	}
// 	if msgs[0].ContentType == constant.MsgRevokeNotification {
// 		return nil, servererrs.ErrMsgAlreadyRevoke.WrapMsg("msg already revoke")
// 	}

// 	data, _ := json.Marshal(msgs[0])
// 	log.ZDebug(ctx, "GetMsgBySeqs", "conversationID", req.ConversationID, "seq", req.Seq, "msg", string(data))
// 	var role int32
// 	if !authverify.IsAdmin(ctx) {
// 		sessionType := msgs[0].SessionType
// 		switch sessionType {
// 		case constant.SingleChatType:
// 			if err := authverify.CheckAccess(ctx, msgs[0].SendID); err != nil {
// 				return nil, err
// 			}
// 			role = user.AppMangerLevel
// 		case constant.ReadGroupChatType:
// 			members, err := m.GroupLocalCache.GetGroupMemberInfoMap(ctx, msgs[0].GroupID, datautil.Distinct([]string{req.UserID, msgs[0].SendID}))
// 			if err != nil {
// 				return nil, err
// 			}
// 			if req.UserID != msgs[0].SendID {
// 				switch members[req.UserID].RoleLevel {
// 				case constant.GroupOwner:
// 				case constant.GroupAdmin:
// 					if sendMember, ok := members[msgs[0].SendID]; ok {
// 						if sendMember.RoleLevel != constant.GroupOrdinaryUsers {
// 							return nil, errs.ErrNoPermission.WrapMsg("no permission")
// 						}
// 					}
// 				default:
// 					return nil, errs.ErrNoPermission.WrapMsg("no permission")
// 				}
// 			}
// 			if member := members[req.UserID]; member != nil {
// 				role = member.RoleLevel
// 			}
// 		default:
// 			return nil, errs.ErrInternalServer.WrapMsg("msg sessionType not supported", "sessionType", sessionType)
// 		}
// 	}
// 	now := time.Now().UnixMilli()
// 	err = m.MsgDatabase.EditMsg(ctx, req.ConversationID, req.Seq, &model.EditModel{
// 		Role:        role,
// 		UserID:      req.UserID,
// 		Nickname:    user.Nickname,
// 		Time:        now,
// 		NewContent:  req.NewContent,
// 		ContentType: req.ContentType,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	editorUserID := mcontext.GetOpUserID(ctx)
// 	var flag bool

// 	if len(m.config.Share.IMAdminUser.UserIDs) > 0 {
// 		flag = datautil.Contain(editorUserID, m.adminUserIDs...)
// 	}
// 	tips := sdkws.EditMsgTips{
// 		EditorUserID:   editorUserID,
// 		ClientMsgID:    msgs[0].ClientMsgID,
// 		EditTime:       now,
// 		Seq:            req.Seq,
// 		SessionType:    msgs[0].SessionType,
// 		ConversationID: req.ConversationID,
// 		IsAdminEdit:    flag,
// 		NewContent:     req.NewContent,
// 		ContentType:    req.ContentType,
// 	}
// 	var recvID string
// 	if msgs[0].SessionType == constant.ReadGroupChatType {
// 		recvID = msgs[0].GroupID
// 	} else {
// 		recvID = msgs[0].RecvID
// 	}
// 	m.notificationSender.NotificationWithSessionType(ctx, req.UserID, recvID, constant.MsgEditNotification, msgs[0].SessionType, &tips)
// 	// TODO: 添加webhook支持
// 	// m.webhookAfterEditMsg(ctx, &m.config.WebhooksConfig.AfterEditMsg, req)
// 	return &msg.EditMsgResp{}, nil
// }
