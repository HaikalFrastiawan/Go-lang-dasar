package simple

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}