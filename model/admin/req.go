package admin

type (
	ReqHeader struct {
		Authorization string `binding:"required"`
	}

	SeedReq struct {
		Email    string `validate:"email"`
		Password string `validate:"gte=8"`
	}

	SignInReqBody struct {
		Email	 string `binding:"email"`
		Password string `binding:"gte=8"`
	}
)
