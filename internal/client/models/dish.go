package models

type Catalogs []*Catalog

type Catalog struct {
	Name   string `json:"name"`
	Dishes Dishes `json:"dishes"`
}

type Dishes []*Dish

type Dish struct {
	TotalOrder                int          `json:"total_order"`
	ListingStatus             bool         `json:"listing_status"`
	Rank                      int          `json:"rank"`
	PictureUrlFmt             string       `json:"picture_url_fmt"`
	DisplayOrderUnderCategory int          `json:"display_order_under_category"`
	IsHidden                  bool         `json:"is_hidden"`
	ID                        uint         `json:"id"`
	TotalLike                 int          `json:"total_like"`
	IsGroupDiscountItem       bool         `json:"is_group_discount_item"`
	SaleTimeInfo              SaleTimeInfo `json:"sale_time_info"`
	Description               string       `json:"description"`
	RatingGood                int          `json:"rating_good"`
	Price                     float64      `json:"price"`
	MMSImage                  string       `json:"mms_image"`
	Name                      string       `json:"name"`
	CatalogID                 uint         `json:"catalog_id"`
	CategoryInfo              CatalogInfo  `json:"catalog_info"`
	CategoryID                uint         `json:"category_id"`
}

type CatalogInfo struct {
	Level      int    `json:"level"`
	NameVi     string `json:"name_vi"`
	MMSImageID string `json:"mms_image_id"`
	Rank       int    `json:"rank"`
	NameEn     string `json:"name_en"`
	ID         uint   `json:"id"`
}

type SaleTimeInfo struct {
	IsInSaleTime bool `json:"is_in_sale_time"`
}
