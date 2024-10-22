package data

type Data interface {
	GetData() []UserModel
	GetDataById(id int) UserModel
	DeleteData(id int) []UserModel
	InsertData(UserModel) []UserModel
	UpdateDataById(id int, user UserModel) []UserModel
}

type data struct {
	userList []UserModel
}

func InitData() Data {
	return &data{
		userList: []UserModel{
			{
				ID:     1,
				Name:   "A",
				Age:    18,
				Gender: "Male",
			},
			{
				ID:     2,
				Name:   "Chi",
				Age:    17,
				Gender: "Female",
			},
			{
				ID:     3,
				Name:   "Shintaro",
				Age:    18,
				Gender: "Male",
			},
			{
				ID:     4,
				Name:   "Ayano",
				Age:    18,
				Gender: "Female",
			},
		},
	}
}

func (d *data) GetData() []UserModel {
	return d.userList
}

func (d *data) GetDataById(id int) UserModel {
	var foundIn int
	var isFound bool
	isFound = false
	for idx, val := range d.userList {
		if val.ID == id {
			foundIn = idx
			isFound = true
		}
	}

	if isFound {
		return d.userList[foundIn]
	}
	return UserModel{}
}

func (d *data) DeleteData(id int) []UserModel {
	var foundIn int
	var isFound bool
	isFound = false
	for idx, val := range d.userList {
		if val.ID == id {
			foundIn = idx
			isFound = true
		}
	}

	if isFound {
		d.userList[foundIn] = d.userList[len(d.userList)-1]
		d.userList[len(d.userList)-1] = UserModel{}
		d.userList = d.userList[:len(d.userList)-1]
		return d.userList
	}
	return []UserModel{}
}

func (d *data) InsertData(user UserModel) []UserModel {
	user.ID = len(d.userList)
	return append(d.userList, user)
}

func (d *data) UpdateDataById(id int, user UserModel) []UserModel {
	var foundIn int
	var isFound bool
	isFound = false
	for idx, val := range d.userList {
		if val.ID == id {
			foundIn = idx
			isFound = true
		}
	}

	if isFound {
		d.userList[foundIn].Name = user.Name
		d.userList[foundIn].Age = user.Age
		d.userList[foundIn].Gender = user.Gender
	}
	return d.userList
}
