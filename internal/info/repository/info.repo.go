package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// response structure ให้เหมือน TS
type ApiResponse struct {
	Status    string        `json:"status"`
	Message   string        `json:"message"`
	Timestamp string        `json:"timestamp"`
	Data      []interface{} `json:"data"`
}

type JobRepository struct {
	DB *pgxpool.Pool
}

func (r *JobRepository) GetTable() ApiResponse {
	query := `
		select
			id,
			create_date,
			update_date,
			hide_yn,
			fac_site,
			product_type,
			car_no,
			product_name,
			apn_no,
			defect_item,
			claim_failure,
			claim_date,
			rtv_date,
			pln_finish_date,
			customer,
			photo_att,
			sort_yn,
			sort_mark,
			qa_issue_by,
			qa_issue_date,
			proc_rc,
			proc_rc_desc,
			proc_rc_act_ca,
			proc_rc_act_pa,
			proc_rc_act_by,
			proc_rc_act_apprv,
			proc_esc,
			proc_esc_desc,
			proc_esc_act_ca,
			proc_esc_act_pa,
			proc_esc_act_by,
			proc_esc_act_apprv,
			qa_8d_report_att,
			qa_confirm_by,
			qa_confirm_date,
			qa_confirm_remark,
			qa_apprv_by,
			qa_apprv_date,
			qa_apprv_remark,
			hca_require,
			verify_ref_doc,
			verify_ref_date,
			verify_remark,
			area,
			inv_lot_claim,
			proc_rc_date,
			proc_esc_date,
			case
				when pln_finish_date::timestamp < now() 
				 and qa_apprv_date is null then 'Delay'
				when qa_apprv_date is not null 
				 and proc_rc_act_apprv is not null 
				 and proc_esc_act_apprv is not null 
				 and qa_confirm_date is not null then 'Finish'
				when qa_issue_date is not null 
				 and qa_apprv_date is null then 'Ongoing'
				else 'Open'
			end as status
		from qa.qa_pqe_claim_record;
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return ApiResponse{
			Status:    "Catch",
			Message:   "Failed to fetch data: " + err.Error(),
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			Data:      []interface{}{},
		}
	}
	defer rows.Close()

	// ดึง column names มา (สำคัญมาก ถ้าอยากเหมือน TS แบบ dynamic)
	fields := rows.FieldDescriptions()

	result := []interface{}{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return ApiResponse{
				Status:    "Catch",
				Message:   "Failed to read row: " + err.Error(),
				Timestamp: time.Now().Format("2006-01-02 15:04:05"),
				Data:      []interface{}{},
			}
		}

		rowMap := make(map[string]interface{})

		for i, field := range fields {
			rowMap[string(field.Name)] = values[i]
		}

		result = append(result, rowMap)
	}

	// เช็ค error หลัง loop
	if rows.Err() != nil {
		return ApiResponse{
			Status:    "Catch",
			Message:   rows.Err().Error(),
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			Data:      []interface{}{},
		}
	}

	// ไม่มี data
	if len(result) == 0 {
		return ApiResponse{
			Status:    "ERROR",
			Message:   "No data found",
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			Data:      []interface{}{},
		}
	}

	// success
	return ApiResponse{
		Status:    "OK",
		Message:   "Successfully fetched data",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      result,
	}
}