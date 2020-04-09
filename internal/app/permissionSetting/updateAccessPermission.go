package permissionSetting

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type accessPermissionArgs struct {
	AccessShowLoanerNotMakeListAdmin      bool `json:"accessShowLoanerNotMakeListAdmin"`
	AccessShowLoanerNotMakeListWorker     bool `json:"accessShowLoanerNotMakeListWorker"`
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
				set accessShowLoanerNotMakeListAdmin = $1, accessShowLoanerNotMakeListWorker = $2,
					accessShowLoanerNewListAdmin = $3, accessShowLoanerNewListWorker = $4,
					accessShowLoanerInVerifyListAdmin = $5, accessShowLoanerInVerifyListWorker = $6,
					accessShowLoanerWaitApproveListAdmin = $7, accessShowLoanerWaitApproveListWorker = $8,
					accessShowLoanerHadApproveListAdmin = $9, accessShowLoanerHadApproveListWorker = $10,
					accessShowLoanerHadDenyListAdmin = $11, accessShowLoanerHadDenyListWorker = $12,
					verifyQuestionnaireByAdmin = $13, verifyQuestionnaireByWorker = $14,
					sendToApproveByAdmin = $15, sendToApproveByWorker = $16
	`, req.AccessShowLoanerNotMakeListAdmin, req.AccessShowLoanerNotMakeListWorker,
		req.AccessShowLoanerNewListAdmin, req.AccessShowLoanerNewListWorker,
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
