package user

type Vehicle struct {
	userId      UserID
	vehicleType int
}

func NewVehicle(userId UserID, vehicleType int) Vehicle {
	return Vehicle{
		userId:      userId,
		vehicleType: vehicleType,
	}
}

func (v *Vehicle) GetUserId() UserID {
	return v.userId
}

func (v *Vehicle) GetVehicleType() int {
	return v.vehicleType
}
