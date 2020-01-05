package permissionSetting

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type getAccessPermissionRequest struct {
}

func getAccessPermission(ctx context.Context, req getAccessPermissionRequest) (res accessPermissionArgs, err error) {
	err = dbctx.QueryRow(ctx, `
			select 	accessShowLoanerNewListAdmin, 
					accessShowLoanerNewListWorker, 
					accessShowLoanerInVerifyListAdmin,
					accessShowLoanerInVerifyListWorker,
					accessShowLoanerWaitApproveListAdmin,
					accessShowLoanerWaitApproveListWorker,
					verifyQuestionnaireByAdmin,
					verifyQuestionnaireByWorker,
					sendToApproveByAdmin,
					sendToApproveByWorker
			from permissionAccess
	`).Scan(&res.AccessShowLoanerNewListAdmin,
		&res.AccessShowLoanerNewListWorker,
		&res.AccessShowLoanerInVerifyListAdmin,
		&res.AccessShowLoanerInVerifyListWorker,
		&res.AccessShowLoanerWaitApproveListAdmin,
		&res.AccessShowLoanerWaitApproveListWorker,
		&res.VerifyQuestionnaireByAdmin,
		&res.VerifyQuestionnaireByWorker,
		&res.SendToApproveByAdmin,
		&res.SendToApproveByWorker)

	if err != nil {
		return res, ErrAccessPermissionGetData
	}

	return
}
