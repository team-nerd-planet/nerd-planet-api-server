package tag

import (
	"log/slog"

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
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return &jobTags, true
}
