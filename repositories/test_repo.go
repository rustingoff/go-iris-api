package repositories

type testRepo struct {
	db string
}

func NewTestRepo(db string) *testRepo {
	return &testRepo{db}
}

func (r *testRepo) Test() (string, error) {
	return "test", nil
}
