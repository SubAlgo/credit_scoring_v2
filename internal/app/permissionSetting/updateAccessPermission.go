package permissionSetting

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type accessPermissionArgs struct {
	AccessShowLoanerNewListAdmin          bool `json:"accessShowLoanerNewListAdmin"`
	AccessShowLoanerNewListWorker         bool `json:"accessShowLoanerNewListWorker"`
	AccessShowLoanerInVerifyListAdmin     bool `json:"accessShowLoanerInVerifyListAdmin"`
	AccessShowLoanerInVerifyListWorker    bool `json:"accessShowLoanerInVerifyListWorker"`
	AccessShowLoanerWaitApproveListAdmin  bool `json:"accessShowLoanerWaitApproveListAdmin"`
	AccessShowLoanerWaitApproveListWorker bool `json:"accessShowLoanerWaitApproveListWorker"`
	AccessShowLoanerHadApproveListAdmin   bool `json:"accessShowLoanerHadApproveListAdmin"`
	AccessShowLoanerHadApproveListWorker  bool `json:"accessShowLoanerHadApproveListWorker"`
	AccessShowLoanerHadDenyListAdmin      bool `json:"accessShowLoanerHadDenyListAdmin"`
	AccessShowLoanerHadDenyListWorker     bool `json:"accessShowLoanerHadDenyListWorker"`
	VerifyQuestionnaireByAdmin            bool `json:"verifyQuestionnaireByAdmin"`
	VerifyQuestionnaireByWorker           bool `json:"verifyQuestionnaireByWorker"`
	SendToApproveByAdmin                  bool `json:"sendToApproveByAdmin"`
	SendToApproveByWorker                 bool `json:"sendToApproveByWorker"`
}

type permissionSettingResponse struct {
	Message string `json:"message"`
}

func updateAccessPermission(ctx context.Context, req accessPermissionArgs) (res permissionSettingResponse, err error) {
	userID := auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if userID == 0 {
		return res, ErrSignInRequired
	}

	if roleID != 1 {
		return res, ErrPermissionNotAllow
	}

	_, err = dbctx.Exec(ctx, `
			update permissionAccess
				set accessShowLoanerNewListAdmin = $1, accessShowLoanerNewListWorker = $2,
					accessShowLoanerInVerifyListAdmin = $3, accessShowLoanerInVerifyListWorker = $4,
					accessShowLoanerWaitApproveListAdmin = $5, accessShowLoanerWaitApproveListWorker = $6,
					accessShowLoanerHadApproveListAdmin = $7, accessShowLoanerHadApproveListWorker = $8,
					accessShowLoanerHadDenyListAdmin = $9, accessShowLoanerHadDenyListWorker = $10,
					verifyQuestionnaireByAdmin = $11, verifyQuestionnaireByWorker = $12,
					sendToApproveByAdmin = $13, sendToApproveByWorker = $14
	`, req.AccessShowLoanerNewListAdmin, req.AccessShowLoanerNewListWorker,
		req.AccessShowLoanerInVerifyListAdmin, req.AccessShowLoanerInVerifyListWorker,
		req.AccessShowLoanerWaitApproveListAdmin, req.AccessShowLoanerWaitApproveListWorker,
		req.AccessShowLoanerHadApproveListAdmin, req.AccessShowLoanerHadApproveListWorker,
		req.AccessShowLoanerHadDenyListAdmin, req.AccessShowLoanerHadDenyListWorker,
		req.VerifyQuestionnaireByAdmin, req.VerifyQuestionnaireByWorker,
		req.SendToApproveByAdmin, req.SendToApproveByWorker)

	if err != nil {
		return res, ErrUpdatePermission
	}
	res.Message = "บันทึกข้อมูลสำเร็จ"
	return
}
