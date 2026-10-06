package openbank

import (
	"github.com/techpartners-asia/golomt-api-go/openbank/model"
)

// /v1/utility сервисүүдийн хариу нууцлагдаагүй (SPEC 3. Системийн аюулгүй байдал)
// тул postPlain ашиглана. client_id/state/scope query шаардлагагүй.

// 10.1.	Хот, аймагийн жагсаалт авах
func (o *openbank) StateListInq(body model.StateListReq) ([]model.StateListResp, error) {
	return postPlain[[]model.StateListResp](o, "STATEINQ", "/v1/utility/state/inq", body)
}

// 10.2.	Сум, дүүргийн жагсаалт авах
func (o *openbank) DistrictListInq(body model.DistrictListReq) ([]model.DistrictListResp, error) {
	return postPlain[[]model.DistrictListResp](o, "CITYINQ", "/v1/utility/city/inq", body)
}

// 10.3.	Категори төрлөөр сонголтын жагсаалт авах
func (o *openbank) CategoryListInq(body model.CategoryReq) ([]model.CategoryResp, error) {
	return postPlain[[]model.CategoryResp](o, "CATINQ", "/v1/utility/category/inq", body)
}

// 10.4.	Ханшны мэдээлэл авах
func (o *openbank) RateInq(body model.RateReq) (*model.RateResp, error) {
	return postPlain[*model.RateResp](o, "RATEINQ", "/v1/utility/rate/inq", body)
}

// 10.5.	Салбарын жагсаалт авах
func (o *openbank) BranchListInq(body model.BranchListReq) ([]model.BranchListResp, error) {
	return postPlain[[]model.BranchListResp](o, "SOLINQ", "/v1/utility/sol/inq", body)
}

// 10.6.	Бүтээгдэхүүн лавлах
func (o *openbank) ProductListInq(body model.ProductListReq) ([]model.ProductData, error) {
	return postPlain[[]model.ProductData](o, "PROCATINQ", "/v1/utility/product/category/inq", body)
}

// 10.7. RateCode лавлах
func (o *openbank) RateCodeInq(body model.RateCodeReq) (*model.RateCodeResp, error) {
	return postPlain[*model.RateCodeResp](o, "RATECODEINQ", "/v1/utility/rate/code/inq", body)
}

// 10.8. Exchange Rate лавлах
func (o *openbank) ExchangeRateInq(body model.ExchangeRateReq) (*model.ExchangeRateResp, error) {
	// SPEC 1.5.8 дээр X-Golomt-Service нь 10.7-той ижил RATECODEINQ гэж заасан
	return postPlain[*model.ExchangeRateResp](o, "RATECODEINQ", "/v1/utility/exchange/rate/inq", body)
}
