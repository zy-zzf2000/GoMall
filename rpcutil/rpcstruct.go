package rpcutil

//=================首页API部分=================

//推荐接口定义
type RecommendRequest struct {
	Cursor int64 `json:"cursor"`        //游标为每页最后一条数据的 RecommendTime
	Ps     int64 `form:"ps,default=20"` //取游标之后的多少条数据每页大小
}

type RecommendResponse struct {
	Products      []*Product `json:"products"`
	IsEnd         bool       `json:"is_end"`         // 是否最后一页
	RecommendTime int64      `json:"recommend_time"` // 商品列表最后一个商品的推荐时间
}

//产品结构体
type Product struct {
	ID          int64    `json:"id"`          // 商品ID
	Name        string   `json:"name"`        // 产品名称
	Images      []string `json:"images"`      // 图片
	Description string   `json:"description"` // 商品描述
	Price       float64  `json:"price"`       // 商品价格
	Stock       int64    `json:"stock"`       // 库存
	Category    string   `json:"category"`    // 分类
	Status      int64    `json:"status"`      // 状态：1-正常，2-下架
	CreateTime  int64    `json:"create_time"` // 创建时间
	UpdateTime  int64    `json:"update_time"` // 更新时间
}

//抢购接口定义
type FlashSaleResponse struct {
	StartTime int64      `json:"start_time"` // 抢购开始时间
	Products  []*Product `json:"products"`
}

//=================分类API部分=================

//分类接口定义
type CategoryListRequest struct {
	Cursor   int64  `form:"cursor"`        // 分页游标
	Ps       int64  `form:"ps,default=20"` // 每页大小
	Category string `form:"category"`      // 分类
	Sort     string `form:"sort"`          // 排序
}

type CategoryListResponse struct {
	Products []*Product `json:"products"`
	IsEnd    bool       `json:"is_end"`
	LastVal  int64      `json:"last_val"`
}

//=================购物车API部分=================
type CartListRequest struct {
	UID int64 `form:"uid"`
}

type CartListResponse struct {
	Products []*CartProduct `json:"products"`
}

type CartProduct struct {
	Product *Product `json:"product"`
	Count   int64    `json:"count"` // 购买数量
}

//================商品评价部分===================
type ProductCommentRequest struct {
	ProductID int64 `form:"product_id"`
	Cursor    int64 `form:"cursor"`
	Ps        int64 `form:"ps,default=20"`
}

type ProductCommentResponse struct {
	Comments    []*Comment `json:"comments"`
	IsEnd       bool       `json:"is_end"`       // 是否最后一页
	CommentTime int64      `json:"comment_time"` // 评论列表最后一个评论的时间
}

type Comment struct {
	ID         int64    `json:"id"`          // 评论ID
	ProductID  int64    `json:"product_id"`  // 商品ID
	Content    string   `json:"content"`     // 评论内容
	Images     []*Image `json:"images"`      // 评论图片
	User       *User    `json:"user"`        // 用户信息
	CreateTime int64    `json:"create_time"` // 评论时间
	UpdateTime int64    `json:"update_time"` // 更新时间
}

type User struct {
	ID     int64  `json:"id"`     // 用户ID
	Name   string `json:"name"`   // 用户名
	Avatar string `json:"avatar"` // 头像
}

type Image struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}
