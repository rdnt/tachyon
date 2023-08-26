package query

//func (s *Queries) handleEvent(e event.event) error {
//	switch e := e.(type) {
//	case event.SessionCreatedEvent:
//		err := s.sessions.CreateSession(session.Session{
//			Id:        e.Id,
//			Name:      e.Name,
//			ProjectId: e.ProjectId,
//		})
//		if err != nil {
//			return err
//		}
//	case event.UserCreatedEvent:
//		err := s.users.CreateUser(user.User{
//			Id:   e.Id,
//			Name: e.Name,
//		})
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
