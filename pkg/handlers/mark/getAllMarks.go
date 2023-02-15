package mark

import (
	"rsiot/pkg/handlers"
	"rsiot/pkg/models"
)

func GetAllMarks() []models.TransformedMark {
	var mark []models.Mark
	var _marks []models.TransformedMark

	handlers.Database().Order("id asc").Find(&mark)

	return marks(mark, _marks)
}

func marks(mark []models.Mark, _marks []models.TransformedMark) []models.TransformedMark {
	for _, item := range mark {
		_marks = append(
			_marks, models.TransformedMark{
				Id:        item.ID,
				UserId:    item.UserId,
				SubjectId: item.SubjectId,
				Mark:      item.Mark,
			})
	}

	return _marks
}

func GetMarkByUserId(id uint) []models.TransformedMark {
	var mark []models.Mark
	var _marks []models.TransformedMark

	handlers.Database().Where("user_id = ?", id).Find(&mark)

	return marks(mark, _marks)
}

func GetMarkBySubjectId(id uint) []models.TransformedMark {
	var mark []models.Mark
	var _marks []models.TransformedMark

	handlers.Database().Where("subject_id = ?", id).Find(&mark)

	return marks(mark, _marks)
}

func GetMarkByUserAndSubjectId(idUser uint, idSubject uint) []models.TransformedMark {
	var mark []models.Mark
	var _marks []models.TransformedMark

	handlers.Database().Where("user_id = ? AND subject_id = ?", idUser, idSubject).Find(&mark)

	return marks(mark, _marks)
}
