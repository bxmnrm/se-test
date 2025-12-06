package controllers
import (
	"backend/config"
	"backend/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /student-fav/:studentID
func GetStudentFavsByStudent(ctx *gin.Context) {
    studentID := ctx.Param("studentID")
    var favs []models.StudentFav

    // Preload NewsPost เพื่อดึงรายละเอียดข่าวมาแสดง
    if err := config.DB.Preload("NewsPost").Where("student_id = ?", studentID).Find(&favs).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, favs)
}

//ฟิลด์ข้อมูลยังไม่ครบ - โค้ดเลยแดง
// POST /student-fav
/*func CreateStudentFav(ctx *gin.Context) {
    var inputValues models.StudentFav

    // 1. รับข้อมูล StudentID และ NewsID จาก Client
    if err := ctx.ShouldBindJSON(&inputValues); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 2. ตรวจสอบความซ้ำซ้อน (สำคัญมาก!)
    // ป้องกันไม่ให้ Fav ซ้ำ โดยการตรวจสอบว่ามี record ที่มี StudentID และ NewsID คู่เดียวกันอยู่แล้วหรือไม่
    var existingFav models.StudentFav
    check := config.DB.Where("student_id = ? AND news_id = ?", inputValues.StudentID, inputValues.NewsID).First(&existingFav)

    if check.RowsAffected > 0 {
        // 409 Conflict: ข่าวนี้ถูกบันทึก/ไลก์ไปแล้ว
        ctx.JSON(http.StatusConflict, gin.H{"error": "This news has already been favorited by the student."})
        return
    }

    // 3. สร้าง Record
    if err := config.DB.Create(&inputValues).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 201 Created
    ctx.JSON(http.StatusCreated, gin.H{"message": "News favorited successfully", "data": inputValues})
}

// DELETE /student-fav/:studentID/:newsID  (ใช้ Path parameter เพื่อความง่าย)
func DeleteStudentFav(ctx *gin.Context) {
    studentID := ctx.Param("studentID")
    newsID := ctx.Param("newsID")

    var studentFav models.StudentFav
    
    // 1. ค้นหา Record ที่ต้องการลบด้วยคู่ StudentID และ NewsID
    if err := config.DB.Where("student_id = ? AND news_id = ?", studentID, newsID).First(&studentFav).Error; err != nil {
        // 404 Not Found: ไม่พบรายการไลก์ที่ต้องการลบ
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Favorite record not found."})
        return
    }

    // 2. ลบ Record ที่พบ
    if err := config.DB.Delete(&studentFav).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // 200 OK
    ctx.JSON(http.StatusOK, gin.H{"message": "Favorite removed successfully"})
}*/