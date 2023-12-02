package sdk

type UOR interface {
	Create() Statement
	Read() Statement
	Update() Statement
	Delete() Statement
}
