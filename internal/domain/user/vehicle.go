package user

type Vehicle struct {
	userId      UserId
	vehicleType string
}

func NewVehicle(userId UserId, vehicleType string) Vehicle {
	return Vehicle{
		userId:      userId,
		vehicleType: vehicleType,
	}
}

func (v *Vehicle) GetUserId() UserId {
	return v.userId
}

func (v *Vehicle) GetVehicleType() string {
	return v.vehicleType
}
