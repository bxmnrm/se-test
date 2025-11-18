package controllers
import (
	"backend/config"
	"backend/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /NewsPosts
func GetNewsPosts(ctx *gin.Context) {
	var newsPosts []models.NewsPost

	if err := config.DB.Preload("StatusNews").Find(&newsPosts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newsPosts)
}

// GET /NewsPosts/:id
func GetNewsPostByID(ctx *gin.Context) {
	var newsPost models.NewsPost
	id := ctx.Param("newsID")

	if err := config.DB.Preload("StatusNews").First(&newsPost, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "NewsPost not found"})
		return
	}
	ctx.JSON(http.StatusOK, newsPost)
}

// POST /news-posts
func CreateNewsPost(ctx *gin.Context) {
    var inputValues models.NewsPost

	// ⭐️ เพิ่ม .Preload("Scholarship") ⭐️ ดึงข้อมูลทุนทั้งหมดมา
    if err := ctx.ShouldBindJSON(&inputValues); err != nil {
        // 400 Bad Request: ข้อมูลที่ส่งมาไม่ถูกต้อง
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // บันทึกข้อมูลลงฐานข้อมูล
    if err := config.DB.Create(&inputValues).Error; err != nil {
        // 500 Internal Server Error: ข้อผิดพลาดในการบันทึกฐานข้อมูล
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // 201 Created: สร้างสำเร็จ
    ctx.JSON(http.StatusCreated, inputValues)
}

// PUT /news-posts/:newsID
func UpdateNewsPost(ctx *gin.Context) {
    var newsPost models.NewsPost
    newsID := ctx.Param("newsID")
    
    // 1. ค้นหาข่าวเดิมด้วย NewsID
    if err := config.DB.First(&newsPost, "news_id = ?", newsID).Error; err != nil {
        // 404 Not Found: ไม่พบข่าวที่ต้องการแก้ไข
        ctx.JSON(http.StatusNotFound, gin.H{"error": "NewsPost not found"})
        return
    }
    
    // 2. รับข้อมูลใหม่จาก Request Body
    var inputValues models.NewsPost
    if err := ctx.ShouldBindJSON(&inputValues); err != nil {
        // 400 Bad Request
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 3. อัปเดตฟิลด์ที่ต้องการ
    newsPost.Title = inputValues.Title
    newsPost.StatusNewsID = inputValues.StatusNewsID
    newsPost.FilePath = inputValues.FilePath

    // 4. บันทึกการเปลี่ยนแปลง
    if err := config.DB.Save(&newsPost).Error; err != nil {
        // 500 Internal Server Error
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // 200 OK: อัปเดตสำเร็จ
    ctx.JSON(http.StatusOK, newsPost)
}

// PATCH /news-posts/:newsID/hide
func HideNewsPost(ctx *gin.Context) {
    var newsPost models.NewsPost
    newsID := ctx.Param("newsID")
    
    // 1. ค้นหาข่าวเดิม
    if err := config.DB.First(&newsPost, "news_id = ?", newsID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "NewsPost not found"})
        return
    }
    
    // 2. กำหนด ID สถานะ "HIDE" (สมมติ ID สถานะซ่อนคือ "HIDE")
    // **ข้อควรระวัง:** คุณต้องมั่นใจว่า ID "HIDE" มีอยู่ในตาราง StatusNews
    const StatusHideID = "HIDE" 
    
    newsPost.StatusNewsID = StatusHideID

    // 3. บันทึกการเปลี่ยนแปลง
    if err := config.DB.Save(&newsPost).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"message": "NewsPost status updated to HIDE", "newsPost": newsPost})
}

// DELETE /news-posts/:newsID
func DeleteNewsPost(ctx *gin.Context) {
    var newsPost models.NewsPost
    newsID := ctx.Param("newsID")

    // 1. ค้นหาข่าวที่ต้องการลบ
    if err := config.DB.First(&newsPost, "news_id = ?", newsID).Error; err != nil {
        // 404 Not Found
        ctx.JSON(http.StatusNotFound, gin.H{"error": "NewsPost not found"})
        return
    }
    
    // 2. ลบข้อมูล
    if err := config.DB.Delete(&newsPost).Error; err != nil {
        // 500 Internal Server Error
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // 200 OK: ลบสำเร็จ
    ctx.JSON(http.StatusOK, gin.H{"message": "NewsPost deleted successfully"})
}