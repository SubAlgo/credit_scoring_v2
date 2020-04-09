package permissionSetting

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type getAccessPermissionRequest struct {
}

func getAccessPermission(ctx context.Context, req getAccessPermissionRequest) (res accessPermissionArgs, err error) {
	err = dbctx.QueryRow(ctx, `
			select 	accessShowLoanerNotMakeListAdmin,
					accessShowLoanerNotMakeListWorker,
					accessShowLoanerNewListAdmin, 
					accessShowLoanerNewListWorker, 
					accessShowLoanerInVerifyListAdmin,
					accessShowLoanerInVerifyListWorker,
					accessShowLoanerWaitApproveListAdmin,
					accessShowLoanerWaitApproveListWorker,
					accessShowLoanerHadApproveListAdmin,
					accessShowLoanerHadApproveListWorker,
					accessShowLoanerHadDenyListAdmin,
					accessShowLoanerHadDenyListWorker,
					verifyQuestionnaireByAdmin,
					verifyQuestionnaireByWorker,
					sendToApproveByAdmin,
					sendToApproveByWorker
			from permissionAccess
	`).Scan(&res.AccessShowLoanerNotMakeListAdmin,
		&res.AccessShowLoanerNotMakeListWorker,
		&res.AccessShowLoanerNewListAdmin,
		&res.AccessShowLoanerNewListWorker,
		&res.AccessShowLoanerInVerifyListAdmin,
		&res.AccessShowLoanerInVerifyListWorker,
		&res.AccessShowLoanerWaitApproveListAdmin,
		&res.AccessShowLoanerWaitApproveListWorker,
		&res.AccessShowLoanerHadApproveListAdmin,
		&res.AccessShowLoanerHadApproveListWorker,
		&res.AccessShowLoanerHadDenyListAdmin,
		&res.AccessShowLoanerHadDenyListWorker,
		&res.VerifyQuestionnaireByAdmin,
		&res.VerifyQuestionnaireByWorker,
		&res.SendToApproveByAdmin,
		&res.SendToApproveByWorker)

	if err != nil {
		return res, ErrAccessPermissionGetData
	}

	return
}
