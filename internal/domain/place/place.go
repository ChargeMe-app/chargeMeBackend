package place

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
)

type PlaceID string

func (placeID PlaceID) String() string {
	return string(placeID)
}

type Place struct {
	placeID         PlaceID
	name            string
	score           *float32
	longitude       float32
	latitude        float32
	access          *int
	iconType        *string
	address         string
	description     *string
	cost            *bool
	costDescription *string
	hours           *string
	open247         *bool
	isComingSoon    *bool
	phoneNumber     *string
	domain.Model
}

func NewPlace(
	name string,
	score *float32,
	longitude float32,
	latitude float32,
	access *int,
	iconType *string,
	address string,
	description *string,
	cost *bool,
	costDescription *string,
	hours *string,
	open247 *bool,
	isComingSoon *bool,
	phoneNumber *string,
	model domain.Model,
) Place {
	return Place{
		placeID:         PlaceID(uuid.New().String()),
		name:            name,
		longitude:       longitude,
		latitude:        latitude,
		score:           score,
		access:          access,
		iconType:        iconType,
		address:         address,
		description:     description,
		cost:            cost,
		costDescription: costDescription,
		hours:           hours,
		open247:         open247,
		isComingSoon:    isComingSoon,
		phoneNumber:     phoneNumber,
		Model:           model,
	}
}

func NewPlaceWithID(
	placeID PlaceID,
	name string,
	score *float32,
	longitude float32,
	latitude float32,
	access *int,
	iconType *string,
	address string,
	description *string,
	cost *bool,
	costDescription *string,
	hours *string,
	open247 *bool,
	isComingSoon *bool,
	phoneNumber *string,
	model domain.Model,
) Place {
	return Place{
		placeID:         placeID,
		name:            name,
		longitude:       longitude,
		latitude:        latitude,
		score:           score,
		access:          access,
		iconType:        iconType,
		address:         address,
		description:     description,
		cost:            cost,
		costDescription: costDescription,
		hours:           hours,
		open247:         open247,
		isComingSoon:    isComingSoon,
		phoneNumber:     phoneNumber,
		Model:           model,
	}
}

func (p *Place) GetPlaceID() PlaceID {
	return p.placeID
}

func (p *Place) GetPlaceName() string {
	return p.name
}

func (p *Place) GetPlaceScore() *float32 {
	return p.score
}

func (p *Place) SetPlaceScore(score *float32) {
	p.score = score
}

func (p *Place) GetPlaceLongitude() float32 {
	return p.longitude
}

func (p *Place) GetPlaceLatitude() float32 {
	return p.latitude
}

func (p *Place) GetPlaceAccess() *int {
	return p.access
}

func (p *Place) SetPlaceAccess(access *int) {
	p.access = access
}

func (p *Place) GetPlaceIconType() *string {
	return p.iconType
}

func (p *Place) SetPlaceIconType(iconType *string) {
	p.iconType = iconType
}

func (p *Place) GetPlaceAddress() string {
	return p.address
}

func (p *Place) SetCost(cost *bool) {
	p.cost = cost
}

func (p *Place) GetCost() *bool {
	return p.cost
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

func (p *Place) SetIsComingSoon(flag *bool) {
	p.isComingSoon = flag
}

func (p *Place) IsComingSoon() *bool {
	return p.isComingSoon
}

func (p *Place) SetPhoneNumber(number *string) {
	p.phoneNumber = number
}

func (p *Place) GetPhoneNumber() *string {
	return p.phoneNumber
}

func (p *Place) SetDescription(description *string) {
	p.description = description
}

func (p *Place) GetDescription() *string {
	return p.description
}
