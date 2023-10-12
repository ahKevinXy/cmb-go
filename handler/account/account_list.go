package account

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/v1/cmb_errors"
	"github.com/ahKevinXy/go-cmb/v1/constants"
	"github.com/ahKevinXy/go-cmb/v1/help"
	"github.com/ahKevinXy/go-cmb/v1/models"

	"strconv"
	"time"
)

// MainAccountUsers
// @Description:   获取可经办列表
// @Author ahKevinXy
// @Date 2023-02-13 13:16:21
func MainAccountUsers(userId, sm4PrivateKey, userPrivateKey, buscod, busmod string) (*models.MainAccountUsersResponse, error) {
	reqData := new(models.MainAccountUsersRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountUserList
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Buscode = buscod
	reqData.Request.Body.Busmod = busmod
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo 优化 返回参数
	res, err := help.CmbSignRequest(string(req), constants.CmbAccountUserList, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {

		return nil, cmb_errors.SystemError
	}

	var resp models.MainAccountUsersResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
