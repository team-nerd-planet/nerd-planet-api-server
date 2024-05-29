package subscription_dto

import "github.com/team-nerd-planet/api-server/internal/entity"

type ApplyReq struct {
	Email                   string                   `json:"email" binding:"required,email"`               // 이메일
	Name                    *string                  `json:"name" binding:"omitempty"`                     // 이름
	Division                *string                  `json:"division" binding:"omitempty"`                 // 소속
	PreferredCompanyArr     []int64                  `json:"preferred_company_arr" binding:"required"`     // 회사 DB ID 배열
	PreferredCompanySizeArr []entity.CompanySizeType `json:"preferred_companySize_arr" binding:"required"` // 회사 규모 배열 (0:스타트업, 1:중소기업, 2:중견기업, 3:대기업, 4:외국계)
	PreferredJobArr         []int64                  `json:"preferred_job_arr" binding:"required"`         // 직무 DB ID 배열
	PreferredSkillArr       []int64                  `json:"preferred_skill_arr" binding:"required"`       // 스킬 DB ID 배열
}

type ApplyRes struct {
	Ok bool `json:"ok"` // 구독 신청 메일 전송 결과
}

func (asr ApplyReq) NewSubscription() entity.Subscription {
	companySizeArr := make([]int64, len(asr.PreferredCompanySizeArr))
	for i, companySize := range asr.PreferredCompanySizeArr {
		companySizeArr[i] = int64(companySize)
	}

	return entity.Subscription{
		Email:                   asr.Email,
		Name:                    asr.Name,
		Division:                asr.Division,
		PreferredCompanyArr:     asr.PreferredCompanyArr,
		PreferredCompanySizeArr: companySizeArr,
		PreferredJobArr:         asr.PreferredJobArr,
		PreferredSkillArr:       asr.PreferredSkillArr,
	}
}
