package models

import "time"

type Product struct {
	Id 				string			`json:"id"`
	Category 		string			`json:"category"`
	Name 			string			`json:"name"`
	Images 			[]string		`json:"images"`
	Price 			string			`json:"price"`
	PromoPrice 		string			`json:"promoPrice"`
	PreOrder 		int				`json:"preOrder"` // In Days
	Description 	string			`json:"description"`
	Digital 		bool			`json:"digital"`
	Slug 			string			`json:"slug"`
	Status 			string			`json:"status"` // Published, Archived, Draft
	Variants 		[]ProductVariant	`json:"variants"`
	CreatedAt 		time.Time		`json:"createdAt"`
}

type ProductVariant struct {
	Id 				string			`json:"id"`
	ProductId 		string			`json:"productId"`
	Name			string			`json:"name"`
	Price 			string			`json:"price"`
	Status 			bool			`json:"status"`
	SKU 			string			`json:"sku"`
}

type PrintfulProductResult struct {
	SyncProduct		PrintfulProduct		`json:"sync_product"`
	SyncVariants 	[]PrintfulVariant	`json:"sync_variants"`
}

type PrintfulProduct struct {
	Id 				int				`json:"id"`
	Name 			string			`json:"name"`
	Thumbnail 		string			`json:"thumbnail"`
	IsIgnored 		bool			`json:"is_ignored"`
}

type PrintfulVariant struct {
	Id 				int				`json:"id"`
	ProductId		int				`json:"product_id"`
	Name 			string			`json:"name"`
	RetailPrice		string			`json:"retail_price"`
	Sku				string			`json:"sku"`
	Currency 		string			`json:"currency"`
	VariantId 		int				`json:"variant_id"`
	IsIgnored 		bool			`json:"is_ignored"`
}

type PrintfulResult struct {
	Result 			PrintfulProductResult		`json:"result"`
}