package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rsiot/pkg/handlers/mark"
	"rsiot/pkg/handlers/subject"
	"rsiot/pkg/handlers/user"
	_ "rsiot/pkg/models"
	"strconv"
)

// AddUser godoc
// @Summary Retrieves user based on given JSON
// @Produce json
// @Param req body models.UserRequestBody true "InstanceCreateReq"
// @Tags user
// @Success 200 {string} User created successfully!
// @Failure 400 {string} Error
// @Router /users [post]
func addUser(c *gin.Context) {
	flag := user.CreateUser(c)

	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusCreated,
			"message": "User created successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// GetAllUser godoc
// @Summary Without any parameters
// @Produce json
// @Tags user
// @Success 200 {object} models.TransformedUser
// @Failure 404 {string} Error
// @Router /users [get]
func getAllUsers(c *gin.Context) {
	userList := user.GetAllUsers()

	if len(userList) > 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusOK,
			"message": userList,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Tags user
// @Success 200 {object} models.TransformedUser
// @Failure 404 {string} Error
// @Router /users/{id} [get]
func getUsersById(c *gin.Context) {
	id := getUIntId(c, "id")

	userById := user.GetUserById(id)

	if userById.Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": userById,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

func getUIntId(c *gin.Context, paramName string) uint {
	id, err := strconv.ParseUint(c.Param(paramName), 10, 64)
	if err != nil {
		panic("Invalid id")
	}

	return uint(id)
}

// UpdateUser godoc
// @Summary Retrieves user based on given ID and JSON
// @Produce json
// @Param id path integer true "User ID"
// @Param req body models.UserRequestBody true "InstanceCreateReq"
// @Tags user
// @Success 200 {string} User updated successfully!
// @Failure 400 {string} Error
// @Router /users/{id} [patch]
func updateUser(c *gin.Context) {
	id := getUIntId(c, "id")

	flag := user.UpdateUser(id, c)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User updated successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// DeleteUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Tags user
// @Success 200 {string} User deleted successfully!
// @Failure 400 {string} Error
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	id := getUIntId(c, "id")

	flag := user.DeleteUser(id)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User deleted successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// AddSubject godoc
// @Summary Retrieves subject based on given JSON
// @Produce json
// @Param req body models.SubjectRequestBody true "InstanceCreateReq"
// @Tags subject
// @Success 200 {string} Subject created successfully!
// @Failure 400 {string} Error
// @Router /subjects [post]
func addSubject(c *gin.Context) {
	flag := subject.CreateSubject(c)

	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusCreated,
			"message": "Subject created successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// GetAllSubjects godoc
// @Summary Without any parameters
// @Produce json
// @Tags subject
// @Success 200 {object} models.TransformedSubject
// @Failure 404 {string} Error
// @Router /subjects [get]
func getAllSubjects(c *gin.Context) {
	subjectList := subject.GetAllSubjects()

	if len(subjectList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": subjectList,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

// GetSubject godoc
// @Summary Retrieves subject based on given ID
// @Produce json
// @Param id path integer true "Subject ID"
// @Tags subject
// @Success 200 {object} models.TransformedSubject
// @Failure 404 {string} Error
// @Router /subjects/{id} [get]
func getSubjectById(c *gin.Context) {
	id := getUIntId(c, "id")

	subjectById := subject.GetSubjectById(id)

	if subjectById.Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": subjectById,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

// UpdateSubject godoc
// @Summary Retrieves subject based on given ID and JSON
// @Produce json
// @Param id path integer true "Subject ID"
// @Param req body models.SubjectRequestBody true "InstanceCreateReq"
// @Tags subject
// @Success 200 {string} Subject updated successfully!
// @Failure 400 {string} Error
// @Router /subjects/{id} [patch]
func updateSubject(c *gin.Context) {
	id := getUIntId(c, "id")

	flag := subject.UpdateSubject(id, c)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Subject updated successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// DeleteSubject godoc
// @Summary Retrieves subject based on given ID
// @Produce json
// @Param id path integer true "Subject ID"
// @Tags subject
// @Success 200 {string} Subject deleted successfully!
// @Failure 400 {string} Error
// @Router /subjects/{id} [delete]
func deleteSubject(c *gin.Context) {
	id := getUIntId(c, "id")

	flag := subject.DeleteSubject(id)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Subject deleted successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// AddMark godoc
// @Summary Retrieves mark based on given JSON
// @Produce json
// @Param req body models.MarkRequestBody true "InstanceCreateReq"
// @Tags mark
// @Success 200 {string} Mark created successfully!
// @Failure 400 {string} Error
// @Router /marks [post]
func addMark(c *gin.Context) {
	flag := mark.CreateMark(c)

	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusCreated,
			"message": "Marks created successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// GetAllMarks godoc
// @Summary Without any parameters
// @Produce json
// @Tags mark
// @Success 200 {object} models.TransformedMark
// @Failure 404 {string} Error
// @Router /marks [get]
func getAllMarks(c *gin.Context) {
	markList := mark.GetAllMarks()

	if len(markList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": markList,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

// GetMark godoc
// @Summary Retrieves mark based on given ID
// @Produce json
// @Param id path integer true "Mark ID"
// @Tags mark
// @Success 200 {object} models.TransformedMark
// @Failure 404 {string} Error
// @Router /marks/{id} [get]
func getMarkById(c *gin.Context) {
	id := getUIntId(c, "id")

	markById := mark.GetMarkById(id)

	if markById.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": markById,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

// UpdateMark godoc
// @Summary Retrieves mark based on given ID and JSON
// @Produce json
// @Param id path integer true "Mark ID"
// @Param req body models.MarkRequestBody true "InstanceCreateReq"
// @Tags mark
// @Success 200 {string} Mark updated successfully!
// @Failure 400 {string} Error
// @Router /marks/{id} [patch]
func updateMark(c *gin.Context) {
	id := getUIntId(c, "id")

	flag := mark.UpdateMark(id, c)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Mark updated successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// DeleteMark godoc
// @Summary Retrieves mark based on given ID
// @Produce json
// @Param id path integer true "Mark ID"
// @Tags mark
// @Success 200 {string} Mark deleted successfully!
// @Failure 400 {string} Error
// @Router /marks/{id} [delete]
func deleteMark(c *gin.Context) {
	id := getUIntId(c, "id")

	flag := mark.DeleteMark(id)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Mark deleted successfully!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
		})
	}
}

// GetMarkByUserId godoc
// @Summary Retrieves mark based on given user ID
// @Produce json
// @Param id path integer true "User ID"
// @Tags mark
// @Success 200 {object} models.TransformedMark
// @Failure 404 {string} Error
// @Router /marks/users/{id} [get]
func getMarkByUserId(c *gin.Context) {
	id := getUIntId(c, "id")

	marksList := mark.GetMarkByUserId(id)

	if len(marksList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": marksList,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

// GetMarkBySubjectId godoc
// @Summary Retrieves mark based on given subject ID
// @Produce json
// @Param id path integer true "Subject ID"
// @Tags mark
// @Success 200 {object} models.TransformedMark
// @Failure 404 {string} Error
// @Router /marks/subjects/{id} [get]
func getMarkBySubjectId(c *gin.Context) {
	id := getUIntId(c, "id")

	marksList := mark.GetMarkBySubjectId(id)

	if len(marksList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": marksList,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}

// GetMarkByUserAndSubjectId godoc
// @Summary Retrieves mark based on given subject user and subject ID
// @Produce json
// @Param id path integer true "User ID"
// @Param idSub path integer true "Subject ID"
// @Tags mark
// @Success 200 {object} models.TransformedMark
// @Failure 404 {string} Error
// @Router /marks/users/{id}/subjects/{idSub} [get]
func getMarkByUserAndSubjectId(c *gin.Context) {
	idUser := getUIntId(c, "id")
	idSubject := getUIntId(c, "idSub")

	marksList := mark.GetMarkByUserAndSubjectId(idUser, idSubject)

	if len(marksList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": marksList,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error",
		})
	}
}
