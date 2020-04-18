package handlers

import(
	"jwt-crud/models"
	"jwt-crud/configs"
	"jwt-crud/helpers"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func CreateTask() echo.HandlerFunc  {
	return func (c echo.Context) error {
		task := new(models.Task)
		c.Bind(task)
		task.Deadline = helpers.ConvertToTime(c.FormValue("deadline"))

		if err := c.Validate(task); err != nil {
			return err 
		}

		db := configs.InitGormDB()
		if err := db.Create(&task).Error; err != nil {
			return c.JSON(400, H{
				"message": "Something went wrong",
			})
		}

		return c.JSON(http.StatusOK, 	H{
			"message": "Task Added",
		})
	}
}

func Tasks() echo.HandlerFunc  {
	return func (c echo.Context) error {
		db := configs.InitGormDB()
		var task []models.Task

		db.Find(&task)

		return c.JSON(http.StatusOK, task)
	}
}

func TaskDetail() echo.HandlerFunc  {
	return func (c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		db := configs.InitGormDB()
		var task models.Task

		if db.First(&task, id).RecordNotFound() {
			return c.JSON(404, H{
				"message": "Task not found",
			})
		}
		
		return c.JSON(http.StatusOK, task)
	}
}

func TaskUpdate() echo.HandlerFunc  {
	return func (c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		task := new(models.Task)

		db := configs.InitGormDB()
		if db.First(&task, id).RecordNotFound() {
			return c.JSON(404, H{
				"message": "Task not found",
			})
		}

		c.Bind(task)
		task.Deadline = helpers.ConvertToTime(c.FormValue("deadline"))
		if err := c.Validate(task); err != nil {
			return err 
		}

		db.Save(&task)

		return c.JSON(http.StatusOK, H{
			"message": "Task Updated",
		})
	}
}

func TaskDelete() echo.HandlerFunc  {
	return func (c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		db := configs.InitGormDB()
		var task models.Task

		if db.First(&task, id).RecordNotFound() {
			return c.JSON(404, H{
				"message": "Task not found",
			})
		}

		db.Delete(&task)
		
		return c.JSON(http.StatusOK, H{
			"message": "Task Deleted",
		})
	}
}