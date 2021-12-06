package vk

type UserFull struct {
	*UserBase
	Counters     *UserCounters
	Education    *UserEducation
	LastSeen     *UserLastSeen `json:"last_seen"`
	Occupation   *UserOccupation
	Online       *UserOnline
	Permission   *UserPermission
	Photo        *UserPhoto
	Schools      *UserSchools
	Universities *UserUniversities
}

type UserBase struct {
	ID              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	CanAccessClosed bool   `json:"can_access_closed"`
	IsClosed        bool   `json:"is_closed"`
	Deactivated     string `json:"deactivated"`
	Hidden          bool   `json:"hidden"`
}

type UserCounters struct {
	Albums        int64 `json:"albums"`
	Videos        int64 `json:"videos"`
	Audios        int64 `json:"audios"`
	Notes         int64 `json:"notes "`
	Friends       int64 `json:"friends"`
	Groups        int64 `json:"groups"`
	OnlineFriends int64 `json:"online_friends"`
	MutualFriends int64 `json:"mutual_friends"`
	UserVideos    int64 `json:"user_videos"`
	Followers     int64 `json:"followers"`
	// Fields returned only for desktop applications:
	UserPhotos    int64 `json:"user_photos"`
	Subscriptions int64 `json:"subscriptions"`
}

type UserEducation struct {
	University     int64  `json:"university"`
	UniversityName string `json:"university_name"`
	Faculty        int64  `json:"faculty"`
	FacultyName    string `json:"faculty_name"`
	Graduation     int8   `json:"graduation"`
}

type UserFields struct {
	Sex      int8   `json:"sex"`
	Bdate    string `json:"bdate"`
	City     string `json:"city"`
	Country  int64  `json:"country"`
	HomeTown int64  `json:"home_town"`
}

type UserLastSeen struct {
	Platform *int8  `json:"platform"`
	Time     *int64 `json:"time"`
}

type UserOccupation struct {
	Type *string `json:"type"`
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
}

type UserOnline struct {
	Online       *bool  `json:"online"`
	OnlineMobile *bool  `json:"online_mobile"`
	OnlineApp    *int64 `json:"online_app"`
}

type UserPermission struct {
	CanPost                *bool `json:"can_post"`
	CanSeeAllPosts         *bool `json:"can_see_all_posts"`
	CanSeeAudio            *bool `json:"can_see_audio"`
	CanWritePrivateMessage *bool `json:"can_write_private_message"`
}

type UserPhoto struct {
	Photo50      *string `json:"photo_50"`
	Photo100     *string `json:"photo_100"`
	Photo200Orig *string `json:"photo_200_orig"`
	Photo200     *string `json:"photo_200"`
	Photo400Orig *string `json:"photo_400_orig"`
	PhotoMax     *string `json:"photo_max"`
	PhotoMaxOrig *string `json:"photo_max_orig"`
}

type UserSchools struct {
	Id            *int64  `json:"id"`
	Country       *int64  `json:"country"`
	City          *int64  `json:"city"`
	Name          *string `json:"name"`
	YearFrom      *int8   `json:"year_from"`
	YearTo        *int8   `json:"year_to"`
	YearGraduated *int8   `json:"year_graduated"`
	Class         *string `json:"class"`
	Speciality    *string `json:"speciality"`
	Type          *int64  `json:"type"`
	TypeStr       *string `json:"type_str"`
}

type UserUniversities struct {
	Id          *int64  `json:"id"`
	Country     *int64  `json:"country"`
	City        *int64  `json:"city"`
	Name        *string `json:"name"`
	Faculty     *int64  `json:"faculty"`
	FacultyName *string `json:"faculty_name"`
	Chair       *int64  `json:"chair"`
	ChairName   *string `json:"chair_name"`
	Graduation  *int8   `json:"graduation"`
}
