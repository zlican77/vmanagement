package controller

import (
	"time"
	"vmbackend/config"
	"vmbackend/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"message": "无效的请求数据"})
		return
	}

	if loginData.Username == "admin" && loginData.Password == "admin" {
		token := "admin-token"
		c.JSON(200, gin.H{"message": "登录成功", "token": token})
	} else {
		c.JSON(401, gin.H{"message": "用户名或密码错误"})
	}
}

func GetGoodsList(c *gin.Context) {
	var goods []model.Good

	if err := config.DB.Find(&goods).Error; err != nil {
		c.JSON(500, gin.H{"message": "获取货物列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"message": "获取货物列表成功",
		"goods":   goods,
	})
}

func GetInboundList(c *gin.Context) {
	var inboundGoods []model.InboundGood

	if err := config.DB.Find(&inboundGoods).Error; err != nil {
		c.JSON(500, gin.H{"message": "获取入库列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"message":      "获取入库列表成功",
		"inboundGoods": inboundGoods,
	})
}

func GetOutboundList(c *gin.Context) {
	var outboundGoods []model.OutboundGood

	if err := config.DB.Find(&outboundGoods).Error; err != nil {
		c.JSON(500, gin.H{"message": "获取出库列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"message":       "获取出库列表成功",
		"outboundGoods": outboundGoods,
	})
}

func GetInventoryList(c *gin.Context) {
	var inventoryGoods []model.InventoryGood

	if err := config.DB.Find(&inventoryGoods).Error; err != nil {
		c.JSON(500, gin.H{"message": "获取库存列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"message":        "获取库存列表成功",
		"inventoryGoods": inventoryGoods,
	})
}

func PostInBound(c *gin.Context) {
	var inbound model.Inbound
	var inventory model.Inventory
	var inventoryGood model.InventoryGood
	var good model.Good

	//时间格式处理
	var frontData struct {
		GoodID      string  `json:"goodsId"`
		InboundTime string  `json:"inboundTime"` // 接收字符串格式的时间
		Quantity    float64 `json:"quantity"`
		Operator    string  `json:"operator"`
	}

	if err := c.ShouldBindJSON(&frontData); err != nil { //将拿到的表单数据映射于对象
		c.JSON(400, gin.H{"message": "无效的请求数据"})
		return

	}
	// 将时间字符串转换为time.Time
	parsedTime, err := time.Parse("2006-01-02T15:04", frontData.InboundTime)
	if err != nil {
		c.JSON(400, gin.H{"message": "无效的时间格式"})
		return
	}

	// 填充inbound结构体
	inbound.GoodID = frontData.GoodID
	inbound.InboundTime = parsedTime
	inbound.Quantity = frontData.Quantity
	inbound.Operator = frontData.Operator

	// 查找货物信息
	if err := config.DB.Where("货物号 = ?", inbound.GoodID).First(&good).Error; err != nil {
		c.JSON(400, gin.H{"message": "货物不存在"})
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	// 添加入库记录， 这里会发生一次触发器
	if err := tx.Create(&inbound).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "入库失败"})
		return
	}

	// 更新库存
	result := tx.Where("货物号 = ?", inbound.GoodID).First(&inventoryGood)
	if result.Error == nil {
		// 货物存在,更新库存量
		//inventoryGood.Quantity += inbound.Quantity
		//bug？ 加倍
		//fmt.Println(inbound.Quantity, inventoryGood.Quantity)
		//inventoryGood.Quantity已经是更新后的了
		//不是bug， 在sql文件中，存在`TT1`触发器，更新入库表则会自动更新库存量
		inventoryGood.UpdateTime = inbound.InboundTime
		if err := tx.Model(&inventory).Where("货物号 = ?", inbound.GoodID).Updates(map[string]interface{}{
			"库存量":  inventoryGood.Quantity,
			"更新日期": inventoryGood.UpdateTime,
		}).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "更新库存失败"})
			return
		}
	} else {
		// 货物不存在,创建新库存记录
		inventoryGood = model.InventoryGood{
			Inventory: model.Inventory{
				GoodID:     good.GoodID,
				UpdateTime: inbound.InboundTime,
				Quantity:   inbound.Quantity,
			},
			Name:  good.Name,
			Spec:  good.Spec,
			Model: good.Model,
		}
		if err := tx.Create(&inventoryGood).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "创建库存记录失败"})
			return
		}
	}

	tx.Commit()
	c.JSON(200, gin.H{"message": "入库成功"})
}

func PostOutBound(c *gin.Context) {
	var outbound model.Outbound
	var inventory model.Inventory
	var inventoryGood model.InventoryGood
	var good model.Good

	//时间格式处理
	var frontData struct {
		GoodID      string  `json:"goodsId"`
		InboundTime string  `json:"outboundTime"` // 接收字符串格式的时间
		Quantity    float64 `json:"quantity"`
		Operator    string  `json:"operator"`
	}

	if err := c.ShouldBindJSON(&frontData); err != nil { //将拿到的表单数据映射于对象
		c.JSON(400, gin.H{"message": "无效的请求数据"})
		return

	}
	// 将时间字符串转换为time.Time
	parsedTime, err := time.Parse("2006-01-02T15:04", frontData.InboundTime)
	if err != nil {
		c.JSON(400, gin.H{"message": "无效的时间格式"})
		return
	}

	// 填充inbound结构体
	outbound.GoodID = frontData.GoodID
	outbound.OutboundTime = parsedTime
	outbound.Quantity = frontData.Quantity
	outbound.Operator = frontData.Operator

	// 查找货物信息
	if err := config.DB.Where("货物号 = ?", outbound.GoodID).First(&good).Error; err != nil {
		c.JSON(400, gin.H{"message": "货物不存在"})
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	if err := tx.Create(&outbound).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "出库失败"})
		return
	}

	// 更新库存
	result := tx.Where("货物号 = ?", outbound.GoodID).First(&inventoryGood)
	if result.Error == nil {
		// 货物存在,更新库存量
		inventoryGood.Quantity -= outbound.Quantity
		//inventoryGood.Quantity已经是更新后的了
		//sql文件不存在出库表触发器
		inventoryGood.UpdateTime = outbound.OutboundTime
		if err := tx.Model(&inventory).Where("货物号 = ?", outbound.GoodID).Updates(map[string]interface{}{
			"库存量":  inventoryGood.Quantity,
			"更新日期": inventoryGood.UpdateTime,
		}).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "更新库存失败"})
			return
		}
	} else {
		// 货物不存在,创建新库存记录
		inventoryGood = model.InventoryGood{
			Inventory: model.Inventory{
				GoodID:     good.GoodID,
				UpdateTime: outbound.OutboundTime,
				Quantity:   outbound.Quantity,
			},
			Name:  good.Name,
			Spec:  good.Spec,
			Model: good.Model,
		}
		if err := tx.Create(&inventoryGood).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"message": "创建库存记录失败"})
			return
		}
	}

	tx.Commit()
	c.JSON(200, gin.H{"message": "出库成功"})
}
