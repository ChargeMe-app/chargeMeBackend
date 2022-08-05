package place

import (
	"github.com/google/uuid"
)

type PlaceID string

func (placeID PlaceID) String() string {
	return string(placeID)
}

type Place struct {
	placeID                      PlaceID
	name                         string
	score                        *float32
	longitude                    float32
	latitude                     float32
	access                       *int
	iconType                     *string
	address                      *string
	accessRestriction            *string
	accessRestrictionDescription *string
	cost                         *bool
	costDescription              *string
	hours                        *string
	open247                      *bool
	isOpenOrActive               *bool
}

func NewPlace(
	name string,
	score *float32,
	longitude float32,
	latitude float32,
	access *int,
	iconType *string,
	address *string,
	accessRestriction *string,
	accessRestrictionDescription *string,
	cost *bool,
	costDescription *string,
	hours *string,
	open247 *bool,
	isOpenOrActive *bool,
) Place {
	p := Place{
		placeID:   PlaceID(uuid.New().String()),
		name:      name,
		longitude: longitude,
		latitude:  latitude,
	}

	p.SetPlaceScore(score)
	p.SetPlaceAccess(access)
	p.SetPlaceIconType(iconType)
	p.SetPlaceAddress(address)
	p.SetAccessRestriction(accessRestriction)
	p.SetAccessRestrictionDescription(accessRestrictionDescription)
	p.SetCost(cost)
	p.SetCostDescription(costDescription)
	p.SetHours(hours)
	p.SetOpen247(open247)
	p.SetIsOpenOrActive(isOpenOrActive)

	return p
}

func NewPlaceWithID(
	placeID PlaceID,
	name string,
	score *float32,
	longitude float32,
	latitude float32,
	access *int,
	iconType *string,
	address *string,
	accessRestriction *string,
	accessRestrictionDescription *string,
	cost *bool,
	costDescription *string,
	hours *string,
	open247 *bool,
	isOpenOrActive *bool,
) Place {
	p := Place{
		placeID:   placeID,
		name:      name,
		longitude: longitude,
		latitude:  latitude,
	}

	p.SetPlaceScore(score)
	p.SetPlaceAccess(access)
	p.SetPlaceIconType(iconType)
	p.SetPlaceAddress(address)
	p.SetAccessRestriction(accessRestriction)
	p.SetAccessRestrictionDescription(accessRestrictionDescription)
	p.SetCost(cost)
	p.SetCostDescription(costDescription)
	p.SetHours(hours)
	p.SetOpen247(open247)
	p.SetIsOpenOrActive(isOpenOrActive)

	return p
}

func (p Place) GetPlaceID() PlaceID {
	return p.placeID
}

func (p Place) GetPlaceName() string {
	return p.name
}

func (p Place) GetPlaceScore() *float32 {
	return p.score
}

func (p *Place) SetPlaceScore(score *float32) {
	p.score = score
}

func (p Place) GetPlaceLongitude() float32 {
	return p.longitude
}

func (p Place) GetPlaceLatitude() float32 {
	return p.latitude
}

func (p Place) GetPlaceAccess() *int {
	return p.access
}

func (p *Place) SetPlaceAccess(access *int) {
	p.access = access
}

func (p Place) GetPlaceIconType() *string {
	return p.iconType
}

func (p *Place) SetPlaceIconType(iconType *string) {
	p.iconType = iconType
}

func (p Place) GetPlaceAddress() *string {
	return p.address
}

func (p *Place) SetPlaceAddress(address *string) {
	p.address = address
}

func (p *Place) SetAccessRestriction(accessRestriction *string) {
	p.accessRestriction = accessRestriction
}

func (p *Place) GetAccessRestriction() *string {
	return p.accessRestriction
}

func (p *Place) SetCost(cost *bool) {
	p.cost = cost
}

func (p *Place) GetCost() *bool {
	return p.cost
}

func (p *Place) SetAccessRestrictionDescription(accessRestrictionDescription *string) {
	p.accessRestrictionDescription = accessRestrictionDescription
}

func (p *Place) GetAccessRestrictionDescription() *string {
	return p.accessRestrictionDescription
}

func (p *Place) SetCostDescription(costDescription *string) {
	p.costDescription = costDescription
}

func (p *Place) GetCostDescription() *string {
	return p.costDescription
}

func (p *Place) SetHours(hours *string) {
	p.hours = hours
}

func (p *Place) GetHours() *string {
	return p.hours
}

func (p *Place) SetOpen247(flag *bool) {
	p.open247 = flag
}

func (p *Place) GetOpen247() *bool {
	return p.open247
}

func (p *Place) SetIsOpenOrActive(flag *bool) {
	p.isOpenOrActive = flag
}

func (p *Place) GetIsOpenOrActive() *bool {
	return p.isOpenOrActive
}
