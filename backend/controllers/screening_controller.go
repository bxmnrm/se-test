package controllers

import (
	"backend/config"
	"backend/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /Screening
func GetScreenings(ctx *gin.Context) {
	var screenings []models.Screening

	//.Preload("Criteria").Preload("Application") - รอเพิ่ม
	if err := config.DB.Preload("Status").Find(&screenings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, screenings)
}

// POST /Screening
func CreateScreening(ctx *gin.Context) {
	var inputValues models.Screening

	if err := ctx.ShouldBindJSON(&inputValues); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return //400 - Bad Request คือ ข้อมูลที่ client ส่งมาไม่ถูกต้อง
	}
	if err := config.DB.Create(&inputValues).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, inputValues) //201 - Created & 500 - Internal Server Error
}

// PUT - Update /Screening/:id
func UpdateScreening(ctx *gin.Context) {
	var screening models.Screening
	id := ctx.Param("id") //id ของ Screening ที่ต้องการแก้ไข
	
	if err := config.DB.First(&screening, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Screening not found"})
		return
	}
	var inputValues models.Screening
	if err := ctx.ShouldBindJSON(&inputValues); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	screening.StatusID = inputValues.StatusID
	screening.RejectReason = inputValues.RejectReason
	if err := config.DB.Save(&screening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, screening)
}

// DELETE /Screening/:id
func DeleteScreening(ctx *gin.Context){
	var screening models.Screening
	id := ctx.Param("id") //id ของ Screening ที่ต้องการลบ

	if err := config.DB.First(&screening, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Screening not found"})
		return
	}
	if err := config.DB.Delete(&screening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Screening deleted successfully"})
}

//ตัวอย่างฟังก์ชันรัน Logic การคัดกรอง
/*func RunScreeningLogic(ctx *gin.Context) {
    // 1. รับ ApplicationID ที่ต้องการตรวจสอบ
    var payload struct {
        ApplicationID string `json:"application_id"`
    }
    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // 2. ดึงข้อมูล Application และ Criteria ทั้งหมด
    var app models.Application
    if err := config.DB.First(&app, "id = ?", payload.ApplicationID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
        return
    }

    var criterias []models.Criteria
    config.DB.Find(&criterias) // ดึง Criteria ทั้งหมด (หรือเฉพาะที่เกี่ยวข้อง)

    // 3. วนลูปและรัน Logic การตรวจสอบ
    var newScreenings []models.Screening

    for _, criteria := range criterias {
        screening := models.Screening{
            ApplicationID: payload.ApplicationID,
            CriterialID:   criteria.ID.String(), // สมมติว่า Criteria มี gorm.Model
            StatusID:       "PASS",            // กำหนดค่าเริ่มต้น
            RejectReason:   "",
        }

        // Logic 1: ตรวจสอบ GPAX
        if app.GPAX < criteria.MinGPAX {
            screening.StatusID = "FAIL"
            screening.RejectReason = "GPAX is below the minimum required: " + fmt.Sprintf("%.2f", criteria.MinGPAX)
        }

        // Logic 2: ตรวจสอบรายได้ครอบครัว
        if criteria.MaxFamilyIncome > 0 && app.FamilyIncomePerYear > criteria.MaxFamilyIncome {
            screening.StatusID = "FAIL"
            screening.RejectReason = "Family income exceeds the maximum allowed: " + fmt.Sprintf("%.2f", criteria.MaxFamilyIncome)
        }

        // **Logic 3: ตรวจสอบชั้นปี (ถ้ายังไม่ FAIL)**
        // ถ้าต้องการตรวจสอบหลายเกณฑ์ ควรใช้ status ของ Screening นั้น ๆ เป็นตัวตัดสิน
        // ... เพิ่ม Logic อื่น ๆ ...

        newScreenings = append(newScreenings, screening)
    }

    // 4. บันทึกผลลัพธ์ทั้งหมดลงในตาราง Screening
    if err := config.DB.Create(&newScreenings).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Screening completed", "results": newScreenings})
}*/
