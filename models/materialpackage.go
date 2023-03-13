package models

type CashTaskMaterialPackage struct {
	Id             uint32 `gorm:"id" json:"id"`
	CashTaskId     uint32 `gorm:"cash_task_id" json:"cash_task_id"` // 变现任务id
	UserId         uint32 `gorm:"user_id" json:"user_id"`           // 用户id
	User           User   `gorm:"foreignKey:UserId;"`
	Cover          string `gorm:"cover" json:"cover"` // 封面
	Name           string `gorm:"name" json:"name"`   // 名称
	CatFoodNum     uint32 `gorm:"cat_food_num" json:"cat_food_num"`
	MaterialNum    int    `gorm:"material_num" json:"material_num"`         // 素材数量
	FolderNum      int    `gorm:"folder_num" json:"folder_num"`             // 文件夹数量
	MaterialUseNum int    `gorm:"material_use_num" json:"material_use_num"` // 素材使用数量
	ThirteenUseNum int    `gorm:"thirteen_use_num" json:"thirteen_use_num"` // 14天素材使用数量
	ImgNum         int    `gorm:"img_num" json:"img_num"`                   // 图片数量
	VideoNum       int    `gorm:"video_num" json:"video_num"`               // 视频数量
	MusicNum       int    `gorm:"music_num" json:"music_num"`               // 音乐数量
	Status         int    `gorm:"status" json:"status"`                     // 状态:0=审核中,1=审核通过,2=审核不通过,3=下架
	PublicStatus   uint32 `gorm:"public_status" json:"public_status"`       // 公域屏蔽状态:0=未屏蔽,1=屏蔽
	Reason         string `gorm:"reason" json:"reason"`
	ApproveTime    int    `gorm:"approve_time" json:"approve_time"`
	Quality        int    `gorm:":quality" json:"quality"`
	AvgScore       int    `gorm:"avg_score" json:"avg_score"`
	ScoreUserNum   uint32 `gorm:"score_user_num" json:"score_user_num"`
	CreateAt       uint32 `gorm:"autoCreateTime" json:"create_at,omitempty"` // 创建时间
	UpdateAt       uint32 `gorm:"autoUpdateTime" json:"update_at,omitempty"` // 更新时间
	IsFree         bool   `gorm:"is_free" json:"is_free"`
}

type CashTaskMaterialPackageBuyRecord struct {
	Id                        uint32                  `gorm:"id" json:"id"`
	UserId                    uint32                  `gorm:"user_id" json:"user_id"`                                             // 用户id
	CashTaskMaterialPackageId uint32                  `gorm:"cash_task_material_package_id" json:"cash_task_material_package_id"` // 变现任务素材包id
	PayOrderId                uint32                  `gorm:"pay_order_id" json:"pay_order_id"`                                   // 支付订单id
	CreatorUserId             uint32                  `gorm:"creator_user_id" json:"creator_user_id"`                             // 素材包所属的用户id
	CatFoodNum                int                     `gorm:"cat_food_num" json:"cat_food_num"`                                   // 猫粮购买数量
	ExpireTime                int                     `gorm:"expire_time" json:"expire_time"`                                     // 猫粮购买过期时间
	CreateAt                  uint32                  `gorm:"autoCreateTime" json:"create_at,omitempty"`                          // 创建时间
	UpdateAt                  uint32                  `gorm:"autoUpdateTime" json:"update_at,omitempty"`                          // 更新时间
	CashTaskMaterialPackage   CashTaskMaterialPackage `gorm:"cash_task_material_package" json:"cash_task_material_package"`
}
