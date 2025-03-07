package admin

func (s *service) SignOut(adminId uint32) (err error) {
	return s.adminRepo.Update(&adminId, &map[string]any{"token": ""})
}
