package app

import (
	"github.com/yann0917/dedao/config"
	"github.com/yann0917/dedao/services"
)

// CourseType 课程分类
func CourseType() (list *services.CourseCategoryList, err error) {
	list, err = getService().CourseType()
	return
}

func CourseList(category, order string, page, limit int) (list *services.CourseList, err error) {
	list, err = getService().CourseList(category, order, page, limit)
	if err != nil {
		return
	}
	idMap := make(config.CourseIDMap, len(list.List))
	switch category {
	case CateCourse:
		for _, course := range list.List {
			idMap[course.ClassID] = GetCourseIDMap(&course)
		}
	case CateAudioBook, CateEbook:
		for _, course := range list.List {
			idMap[course.ID] = GetCourseIDMap(&course)
		}
	}

	config.Instance.SetIDMap(category, idMap)
	return
}

// CourseList 已购课程列表
func CourseListAll(category string) (list *services.CourseList, err error) {
	list, err = getService().CourseListAll(category, "study")
	if err != nil {
		return
	}
	idMap := make(config.CourseIDMap, len(list.List))
	switch category {
	case CateCourse:
		for _, course := range list.List {
			idMap[course.ClassID] = GetCourseIDMap(&course)
		}
	case CateAudioBook, CateEbook:
		for _, course := range list.List {
			idMap[course.ID] = GetCourseIDMap(&course)
		}
	}

	config.Instance.SetIDMap(category, idMap)
	return
}

// CourseInfo 已购课程列表
func CourseInfo(id int) (info *services.CourseInfo, err error) {
	idMap := config.Instance.GetIDMap(CateCourse, id)
	enID := ""
	if enid, ok := idMap["enid"].(string); ok {
		enID = enid
	}
	if enID == "" {
		courseDetail, err1 := CourseDetail(CateCourse, id)
		if err1 != nil {
			err = err1
			return
		}
		enID = courseDetail["enid"].(string)
	}
	info, err = getService().CourseInfo(enID)
	if err != nil {
		return
	}
	return
}

// CourseDetail 已购课程详情
func CourseDetail(category string, id int) (idMap map[string]interface{}, err error) {
	idMap = config.Instance.GetIDMap(category, id)
	enID := ""
	if enid, ok := idMap["enid"].(string); ok {
		enID = enid
	}
	if enID == "" {
		detail, err1 := getService().CourseDetail(category, id)
		if err1 != nil {
			err = err1
			return
		}
		idMap = GetCourseIDMap(detail)
	}
	return
}

func GetCourseIDMap(course *services.Course) map[string]interface{} {
	return map[string]interface{}{
		"enid":               course.Enid,
		"class_id":           course.ClassID,
		"title":              course.Title,
		"publish_num":        course.PublishNum,
		"has_play_auth":      course.HasPlayAuth,
		"audio_alias_id":     course.AudioDetail.AliasID,
		"audio_size":         course.AudioDetail.Size,
		"audio_mp3_play_url": course.AudioDetail.Mp3PlayURL,
	}
}

// CourseInfo 已购课程列表
func CourseInfoByEnid(enID string) (info *services.CourseInfo, err error) {
	info, err = getService().CourseInfo(enID)
	if err != nil {
		return
	}
	return
}
