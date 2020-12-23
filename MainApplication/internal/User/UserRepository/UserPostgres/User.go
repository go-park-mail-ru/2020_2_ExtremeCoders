package UserPostgres

//type dataBase struct {
//	DB pgwrapper.DB
//}
//
//func New(db pgwrapper.DB) UserRepository.UserDB {
//	return dataBase{DB: db}
//}
//
//func (dbInfo dataBase) IsEmailExists(email string) error {
//	user := &UserModel.User{Email: email}
//	err := dbInfo.DB.Model(user).Where("email=?", email).Select()
//	if err != pg.ErrNoRows {
//		return UserRepository.EmailAlreadyExists
//	}
//	return nil
//}
//
//func (dbInfo dataBase) AddUser(user *UserModel.User) error {
//	user.Password = string(PasswordBcrypt([]byte(user.Password)))
//	_, err := dbInfo.DB.Model(user).Insert()
//	if err != nil {
//		return UserRepository.CantAddUser
//	}
//	return nil
//}
//
//func (dbInfo dataBase) AddSession(sid string, uid uint64, user *UserModel.User) error {
//	session := &UserModel.Session{Id: sid, UserId: int64(uid), User: user}
//	_, err := dbInfo.DB.Model(session).Insert()
//	if err != nil {
//		return UserRepository.CantAddSession
//	}
//	return nil
//}
//
//func (dbInfo dataBase) GenerateSID() ([]rune, error) {
//	var sid string
//	for {
//		for i := 0; i < config.SizeSID; i++ {
//			safeNum, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(config.SidRunes))))
//			sid += string(config.SidRunes[safeNum.Int64()])
//		}
//		fmt.Println(sid)
//		session := &UserModel.Session{Id: sid}
//		exist := dbInfo.DB.Model(session).Where("id=?", sid).Select()
//		if exist != nil {
//			break
//		}
//		sid = ""
//	}
//	return []rune(sid), nil
//}
//
//func (dbInfo dataBase) GenerateUID() (uint64, error) {
//	for {
//		uid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
//		user := UserModel.User{Id: uid.Uint64()}
//		exist := dbInfo.DB.Model(user).Where("id=?", uid.Int64()).Select()
//		if exist != nil {
//			return uid.Uint64(), nil
//		}
//	}
//}
//
//func (dbInfo dataBase) GetUserByEmail(email string) (*UserModel.User, error) {
//	user := &UserModel.User{Email: email}
//	err := dbInfo.DB.Model(user).Where("email=?", email).Select()
//	if err != nil {
//		return user, UserRepository.CantGetUserByEmail
//	}
//	return user, nil
//}
//
//func (dbInfo dataBase) GetUserByUID(uid uint64) (*UserModel.User, error) {
//	user := &UserModel.User{}
//	err := dbInfo.DB.Model(user).Where("id=?", uid).Select()
//	if err != nil {
//		return user, UserRepository.CantGetUserByUid
//	}
//	return user, nil
//}
//
//func (dbInfo dataBase) IsOkSession(sid string) (uint64, error) {
//	session := &UserModel.Session{Id: sid}
//	err := dbInfo.DB.Model(session).Where("id=?", sid).Select()
//	if err != nil {
//		return 0, UserRepository.InvalidSession
//	}
//	return uint64(session.UserId), nil
//}
//
//func (dbInfo dataBase) UpdateProfile(newUser UserModel.User, email string) error {
//	oldUser := &UserModel.User{Email: email}
//	err := dbInfo.DB.Model(oldUser).Where("email=?", email).Select()
//	if err != nil {
//		return UserRepository.CantGetUserOnUpdate
//	}
//	User := oldUser
//	User.Name = newUser.Name
//	User.Surname = newUser.Surname
//	User.Img = newUser.Img
//	_, err = dbInfo.DB.Model(User).Column("name", "surname", "img").Where("email=?", email).Update()
//	if err != nil {
//		return UserRepository.CantUpdateUser
//	}
//	return nil
//}
//
//func (dbInfo dataBase) RemoveSession(sid string) (error, uint64) {
//	session := &UserModel.Session{Id: sid}
//	err := dbInfo.DB.Model(session).Where("id=?", sid).Select()
//	_, err = dbInfo.DB.Model(session).Where("id=?", sid).Delete()
//	if err != nil {
//		return UserRepository.RemoveSessionError, 0
//	}
//	return nil, uint64(session.UserId)
//}
//
//func (dbInfo dataBase) GetSessionByUID(uid uint64) (string, error) {
//	session := &UserModel.Session{UserId: int64(uid)}
//	err := dbInfo.DB.Model(session).Where("user_id=?", uid).Select()
//	if err != nil {
//		return "", UserRepository.GetSessionError
//	}
//	return session.Id, nil
//}
//
//func PasswordBcrypt(plainPassword []byte) []byte {
//	passBcrypt, _ := bcrypt.GenerateFromPassword(plainPassword, 14)
//	return passBcrypt
//}
