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

// PayMods
//
//	@Description:  获取支付模式
//	@param userId 用户ID
//	@param sm4PrivateKey 对称加密秘钥
//	@param userPrivateKey 用户私钥
//	@param buscode 业务模式
//	@Author  ahKevinXy
//	@Date 2023-04-06 19:54:15
func PayMods(userId, sm4PrivateKey, userPrivateKey, busCode string) (*models.QueryAccountTransCodeResponse, error) {

	reqData := new(models.QueryAccountTransCodeRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountCanPayMod
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Buscod = busCode

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), constants.CmbAccountCanPayMod, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryAccountTransCodeResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return &resp, nil
}
