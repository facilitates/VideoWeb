package pack

import (
	"strconv"
	"videoweb/biz/model/videoweb"
	"videoweb/dao"
)

func BuildLoginResp(userDB dao.UserDB) videoweb.LoginResponse {
	return videoweb.LoginResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "登陆成功",
		},
		Data: &videoweb.User{
			ID:        strconv.FormatUint(uint64(userDB.ID), 10),
			Username:  userDB.UserName,
			AvatarURL: userDB.AvatarUrl,
			CreatedAt: userDB.CreatedAt.String(),
			UpdatedAt: userDB.UpdatedAt.String(),
		},
	}
}

func BuildRegisterResp() videoweb.RegisterResponse {
	return videoweb.RegisterResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "注册成功",
		},
	}
}

func BuildInfoResp(userDB dao.UserDB) videoweb.InfoResponse {
	return videoweb.InfoResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "查询成功",
		},
		Data: &videoweb.User{
			ID:        strconv.FormatUint(uint64(userDB.ID), 10),
			Username:  userDB.UserName,
			AvatarURL: userDB.AvatarUrl,
			CreatedAt: userDB.CreatedAt.String(),
			UpdatedAt: userDB.UpdatedAt.String(),
		},
	}
}

func BuildUploadAvatarResp(userDB dao.UserDB) videoweb.UploadAvatarResponse {
	return videoweb.UploadAvatarResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "头像上传成功",
		},
		Data: &videoweb.User{
			ID:        strconv.FormatUint(uint64(userDB.ID), 10),
			Username:  userDB.UserName,
			AvatarURL: userDB.AvatarUrl,
			CreatedAt: userDB.CreatedAt.String(),
			UpdatedAt: userDB.UpdatedAt.String(),
		},
	}
}
