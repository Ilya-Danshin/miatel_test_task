package subscriber

type Sub struct {
	allowToRead  bool
	allowToWrite bool
}

func (s *Sub) IsReadAllowed() bool {
	return s.allowToRead
}

func (s *Sub) IsWriteAllowed() bool {
	return s.allowToWrite
}
