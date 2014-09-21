package utils

//201 - Created
//409 - Conflict
//400 - Bad Request
//401 - Unauthorized
//404 - Not Found

type Error struct {
	Status  int
	Message string
}

func (err Error) Error() string {
	return err.Message
}
