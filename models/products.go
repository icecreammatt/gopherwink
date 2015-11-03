package models

type ProductType int64

const (
	Unknown               ProductType = 0x0000
	GELinkBulb                        = 0xce3d
	GoControlSwitch                   = 0x0102
	GoControlMotionSensor             = 0x0203
	GoControlSiren                    = 0x0503
)
