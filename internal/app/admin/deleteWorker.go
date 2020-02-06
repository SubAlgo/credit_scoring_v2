package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/user"
)

type deleteWorkerRequest struct {
	WorkerID int64 `json:"worker_id"`
}

type deleteWorkerResponse struct {
	Message string `json:"message"`
}

func deleteWorker(ctx context.Context, req deleteWorkerRequest) (res deleteWorkerResponse, err error) {
	// check permission access
	{
		userID := auth.GetUserID(ctx)
		userRole := auth.GetUserRole(ctx)
		if userID == 0 {
			return res, ErrNotSignIn
		}

		// ถ้าไม่ใช่ super admin จะไม่มีสิทธิใช้งาน
		if userRole != 1 {
			return res, ErrPermissionNotAllow
		}
	}

	// check ห้ามลบ user super admin
	{
		workerRole, err := user.GetUserRole(ctx, req.WorkerID)
		if err != nil {
			return res, ErrGetWorkerRole
		}

		if workerRole == 1 {
			return res, ErrDisableDeleteSuperUser
		}
	}

	// ลบข้อมูลพนักงาน
	{
		err = user.Delete(ctx, req.WorkerID)
		if err != nil {
			return res, ErrDeleteWorker
		}
	}

	res.Message = "ลบข้อมูลพนักงานสำเร็จ"
	return
}
