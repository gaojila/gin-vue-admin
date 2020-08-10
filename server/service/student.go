package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateStudent
// @description   create a Student
// @param     student               model.Student
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateStudent(student model.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

// @title    DeleteStudent
// @description   delete a Student
// @auth                     （2020/04/05  20:22）
// @param     student               model.Student
// @return                    error

func DeleteStudent(student model.Student) (err error) {
	err = global.GVA_DB.Delete(student).Error
	return err
}

// @title    DeleteStudentByIds
// @description   delete Students
// @auth                     （2020/04/05  20:22）
// @param     student               model.Student
// @return                    error

func DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Student{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateStudent
// @description   update a Student
// @param     student          *model.Student
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateStudent(student *model.Student) (err error) {
	err = global.GVA_DB.Save(student).Error
	return err
}

// @title    GetStudent
// @description   get the info of a Student
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Student        Student

func GetStudent(id uint) (err error, student model.Student) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// @title    GetStudentInfoList
// @description   get Student list by pagination, 分页获取Student
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetStudentInfoList(info request.StudentSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Student{})
    var students []model.Student
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return err, students, total
}