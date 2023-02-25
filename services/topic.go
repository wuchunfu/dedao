package services

// TopicAll topic all
type TopicAll struct {
	HasMore bool         `json:"has_more"`
	List    []TopicIntro `json:"list"`
}

// TopicDetail topic detail
type TopicDetail struct {
	TopicIntro
	Presenters []TopicPresenter `json:"presenters"`
	TopArea    []TopicTopArea   `json:"top_area"`

	CreateTime       int    `json:"create_time"`
	UpdateTime       int    `json:"update_time"`
	State            int    `json:"state"`
	ShareURL         string `json:"share_url"`
	Ddurl            string `json:"ddurl"`
	TopicIDStr       string `json:"topic_id_str"`
	LastNotesID      string `json:"last_notes_id"`
	LastNotesIDHazy  string `json:"last_notes_id_hazy"`
	LastNotesUID     string `json:"last_notes_uid"`
	LastNotesUIDHazy string `json:"last_notes_uid_hazy"`
	LastNotesContent string `json:"last_notes_content"`
	LastUpdateTime   int    `json:"last_update_time"`
}

// TopicIntro topic intro
type TopicIntro struct {
	NotesTopicID string `json:"notes_topic_id"`
	TopicIDHazy  string `json:"topic_id_hazy"`
	Name         string `json:"name"`
	Img          string `json:"img"`
	Topmost      bool   `json:"topmost"`
	Tag          int    `json:"tag"`
	Intro        string `json:"intro"`
	ViewCount    int    `json:"view_count"`
	NotesCount   int    `json:"notes_count"`
	HasNewNotes  bool   `json:"has_new_notes"`
	UserState    int    `json:"user_state"`
	LogID        string `json:"log_id"`
	LogType      string `json:"log_type"`
}

// TopicPresenter topic presenter
type TopicPresenter struct {
	ID         int    `json:"id"`
	IDHazy     string `json:"id_hazy"`
	UID        int    `json:"uid"`
	UIDHazy    string `json:"uid_hazy"`
	IsV        int    `json:"isV"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	TopicCount int    `json:"topic_count"`
	Icon       string `json:"icon"`
	Relation   int    `json:"relation"`
}

// TopicTopArea Topic TopArea
type TopicTopArea struct {
	ID       int    `json:"id"`
	IDHazy   string `json:"id_hazy"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	State    int    `json:"state"`
	IndexNum int    `json:"index_num"`
}

// NotesList topic NotesList
type NotesList struct {
	HasMore          bool          `json:"has_more"`
	List             []interface{} `json:"list"`
	NoteDetailList   []NoteDetail  `json:"note_detail_list"`
	PresenterUID     []interface{} `json:"presenter_uid"`
	PresenterUIDHazy []string      `json:"presenter_uid_hazy"`
}

// NoteDetail note detail
type NoteDetail struct {
	DetailTitle     string        `json:"detail_title"`
	Comb            *[]Comb       `json:"comb"`
	State           int           `json:"state"`
	IsMine          bool          `json:"is_mine"`
	IsReposted      bool          `json:"is_reposted"`
	IsLike          bool          `json:"is_like"`
	OwnUIDHazy      string        `json:"own_uid_hazy"`
	Topic           NoteTopic     `json:"topic"`
	Tags            []interface{} `json:"tags"`
	Folders         interface{}   `json:"folders"`
	NoteCount       NoteCount     `json:"note_count"`
	FPart           NoteFPart     `json:"f_part"`
	SPart           *NoteFPart    `json:"s_part"`
	ShareURL        string        `json:"share_url"`
	Class           int           `json:"class"`
	Level           int           `json:"level"`
	LevelType       int           `json:"level_type"`
	LevelPermission bool          `json:"level_permission"`
	NotesType       int           `json:"notes_type"`
}

type NoteTopic struct {
	TopicID     int    `json:"topic_id"`
	TopicIDHazy string `json:"topic_id_hazy"`
	IsElected   bool   `json:"is_elected"`
	IsTopmost   bool   `json:"is_topmost"`
	TopicName   string `json:"topic_name"`
}

// NoteCount note count
type NoteCount struct {
	RepostCount  int `json:"repost_count"`
	CommentCount int `json:"comment_count"`
	LikeCount    int `json:"like_count"`
}

// NoteFPart Note FPart
type NoteFPart struct {
	UID           int            `json:"uid"`
	UIDHazy       string         `json:"uid_hazy"`
	NickName      string         `json:"nick_name"`
	Avatar        string         `json:"avatar"`
	Follow        int            `json:"follow"`
	IsV           int            `json:"is_v"`
	Slogan        string         `json:"slogan"`
	VInfo         string         `json:"v_info"`
	StudentID     int            `json:"student_id"`
	StudentIDHazy string         `json:"student_id_hazy"`
	IsPoster      bool           `json:"is_poster"`
	QrCode        string         `json:"qr_code"`
	Note          string         `json:"note"`
	TimeDesc      string         `json:"time_desc"`
	NoteTitle     string         `json:"note_title"`
	NoteScore     string         `json:"note_score"`
	NoteLine      string         `json:"note_line"`
	NoteID        string         `json:"note_id"`
	NoteIDHazy    string         `json:"note_id_hazy"`
	Tip           string         `json:"tip"`
	Images        []string       `json:"images"`
	BaseSource    NoteBaseSource `json:"base_source"`
	StyleNoteLine string         `json:"style_note_line"`
}

// NoteBaseSource Note BaseSource
type NoteBaseSource struct {
	Title          string   `json:"title"`
	SubTitle       string   `json:"sub_title"`
	Img            string   `json:"img"`
	PType          int      `json:"p_type"`
	PidStr         string   `json:"pid_str"`
	NotePtype      int      `json:"note_ptype"`
	IsPopLoginView bool     `json:"is_pop_login_view"`
	NeedCheckBuy   bool     `json:"need_check_buy"`
	URL1           string   `json:"url1"`
	URL2           string   `json:"url2"`
	CampHazy       CampHazy `json:"camp_hazy"`
}

type Comb struct {
	Uid     int    `json:"uid"`
	UidHazy string `json:"uid_hazy"`
	Name    string `json:"name"`
}

type CampHazy struct {
	SourceIdHazy string `json:"SourceIdHazy"`
	OriginIdHazy string `json:"OriginIdHazy"`
}

type NotesTimeline struct {
	IsMore bool         `json:"is_more"`
	MaxId  string       `json:"max_id"`
	Notes  []NoteDetail `json:"notes"`
}

// TopicAll topic all
func (s *Service) TopicAll(page, limit int) (list *TopicAll, err error) {

	body, err := s.reqTopicAll(page, limit)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

// TopicDetail topic detail by id hazy
func (s *Service) TopicDetail(id string) (detail *TopicDetail, err error) {

	body, err := s.reqTopicDetail(id)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &detail); err != nil {
		return
	}
	return
}

// TopicNotesList Topic NotesList
func (s *Service) TopicNotesList(id string) (list *NotesList, err error) {

	body, err := s.reqTopicNotesList(id)
	if err != nil {
		return
	}
	defer body.Close()

	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

// TopicNotesTimeline Topic timeline
func (s *Service) TopicNotesTimeline(maxID string) (list *NotesTimeline, err error) {
	body, err := s.reqTopicNotesTimeline(maxID)
	if err != nil {
		return
	}
	defer body.Close()

	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}
