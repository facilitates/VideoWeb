package videoweb

import (
	"errors"
	"log"
	"strconv"
	"time"
	"videoweb/dao"
	"videoweb/pkg/utils"

	"github.com/jinzhu/gorm"
)

func (req *RegisterRequest) UserRegister() error {
	var count int
	var user dao.UserDB
	dao.DB.Model(&dao.UserDB{}).Where("user_name=?", req.Auth.Username).First(&user).Count(&count)
	if count == 1 {
		return errors.New("用户已经存在")
	}
	digestPassword, err := utils.DigestPassword(req.Auth.Password)
	if err != nil {
		return err
	} else {
		user.UserName = req.Auth.Username
		user.Password = digestPassword
		user.UpdatedAt = time.Now()
		user.CreatedAt = time.Now()
		if err = dao.DB.Model(&dao.UserDB{}).Create(&user).Error; err != nil {
			return err
		} else {
			err := utils.CreateFolder(user.UserName)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func (req *LoginRequest) UserLogin() (dao.UserDB, string, uint, int, error) {
	var user dao.UserDB
	err := dao.DB.Model(&dao.UserDB{}).Where("user_name=?", req.Auth.Username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Println(err)
		return user, user.Password, user.ID, user.IsBindMfa, errors.New("用户不存在")
	}
	return user, user.Password, user.ID, user.IsBindMfa, nil
}

func (req *InfoRequest) UserInfo() (error, dao.UserDB) {
	userid, _ := strconv.ParseUint(req.UserID, 10, 64)
	var user dao.UserDB
	err := dao.DB.Where("id=?", userid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("用户不存在"), user
	} else if !gorm.IsRecordNotFoundError(err) {
		return err, user
	}
	return nil, user
}

func AvatarUpload(filepath string, userid uint) (dao.UserDB, error) {
	var user dao.UserDB
	err := dao.DB.Model(&dao.UserDB{}).Where("id=?", userid).First(&user).Error
	user.AvatarUrl = filepath
	user.UpdatedAt = time.Now()
	err = dao.DB.Model(&dao.UserDB{}).Save(&user).Error
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (req *MfaBindRequest) MfaBind(userid uint) error {
	var user dao.UserDB
	err := dao.DB.Model(&dao.UserDB{}).Where("id=?", userid).First(&user).Error
	if err != nil {
		return err
	}
	user.IsBindMfa = 1
	user.MfaSecret = req.Secret
	err = dao.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
