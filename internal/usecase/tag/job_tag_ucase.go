package tag

import (
	"fmt"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type JobTagUsecase struct {
	jobTagRepo entity.JobTagRepo
}

func NewJobTagUsecase(jobTagRepo entity.JobTagRepo) JobTagUsecase {
	return JobTagUsecase{
		jobTagRepo: jobTagRepo,
	}
}

func (stu JobTagUsecase) FindAll() (*[]entity.JobTag, bool) {
	jobTags, err := stu.jobTagRepo.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}

	return &jobTags, true
}
