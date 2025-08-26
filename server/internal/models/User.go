package model

import "time"

// type User struct {
// 	_ID           **int64      `json:_id`
// 	Code          *string    `json:code`
// 	Password      *string    `json:password`
// 	Title_ID      *string    `json:title_id`
// 	Name          *string    `json:name`
// 	Surname       *string    `json:surname`
// 	TitleTH_ID    **int64     `json:TitleTH_ID`
// 	NameTH        *string    `json:NameTH`
// 	SurnameTH     *string    `json:SurnameTH`
// 	TitleJP_ID    **int64     `json:TitleJP_ID`
// 	NameJP        *string    `json:NameJP`
// 	SurnameJP     *string    `json:SurnameJP`
// 	Email         *string    `json:Email`
// 	Avatar        *string    `json:Avatar`
// 	SupplierCode  *string    `json:SupplierCode`
// 	UILanguage    *string    `json:UILanguage`
// 	UITheme       *string    `json:UITheme`
// 	UIHeaderBrand *string    `json:UIHeaderBrand`
// 	UIHeader      *string    `json:UIHeader`
// 	UILinkColor   *string    `json:UILinkColor`
// 	UIMenuColor   *string    `json:UIMenuColor`
// 	UIIconColor   *string    `json:UIIconColor`
// 	UIExpandIcon  *string    `json:UIExpandIcon`
// 	UIMenuIcon    *string    `json:UIMenuIcon`
// 	UISideBar     *string    `json:UISideBar`
// 	Token         *string    `json:Token`
// 	ResetToken    *string    `json:ResetToken`
// 	LastLogin     *time.Time `json:LastLogin`
// 	Status        *string    `json:Status`
// 	UpdateAt      *time.Time `json:UpdateAt`
// 	UpdateBy      *string    `json:UpdateBy`
// 	CreateAt      *time.Time `json:CreateAt`
// 	CreateBy      *string    `json:CreateBy`
// 	isDelete      *int       `json:isDelete`
// }

type User struct {
	_ID           *int64     `json:"_id"`
	Code          *string    `json:"code"`
	Password      *string    `json:"password"`
	Title_ID      *string    `json:"title_id"`
	Name          *string    `json:"name"`
	Surname       *string    `json:"surname"`
	TitleTH_ID    *int64     `json:"TitleTH_ID"`
	NameTH        *string    `json:"NameTH"`
	SurnameTH     *string    `json:"SurnameTH"`
	TitleJP_ID    *int64     `json:"TitleJP_ID"`
	NameJP        *string    `json:"NameJP"`
	SurnameJP     *string    `json:"SurnameJP"`
	Email         *string    `json:"Email"`
	Avatar        *string    `json:"Avatar"`
	SupplierCode  *string    `json:"SupplierCode"`
	UILanguage    *string    `json:"UILanguage"`
	UITheme       *string    `json:"UITheme"`
	UIHeaderBrand *string    `json:"UIHeaderBrand"`
	UIHeader      *string    `json:"UIHeader"`
	UILinkColor   *string    `json:"UILinkColor"`
	UIMenuColor   *string    `json:"UIMenuColor"`
	UIIconColor   *string    `json:"UIIconColor"`
	UIExpandIcon  *string    `json:"UIExpandIcon"`
	UIMenuIcon    *string    `json:"UIMenuIcon"`
	UISideBar     *string    `json:"UISideBar"`
	Token         *string    `json:"Token"`
	ResetToken    *string    `json:"ResetToken"`
	LastLogin     *time.Time `json:"LastLogin"`
	Status        *string    `json:"Status"`
	UpdateAt      *time.Time `json:"UpdateAt"`
	UpdateBy      *string    `json:"UpdateBy"`
	CreateAt      *time.Time `json:"CreateAt"`
	CreateBy      *string    `json:"CreateBy"`
	isDelete      *int       `json:"isDelete"`
}

func (User) TableName() string {
	return "erp.user"
}
