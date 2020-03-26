package cli

type (
	Command struct {
		OptionString   string  `cli:"short:s type:option default:"`
		OptionInt      int     `cli:"short:i type:option default:0"`
		OptionInt8     int8    `cli:"short:8 type:option default:0"`
		OptionFloat32  float32 `cli:"short:f type:option default:0"`
		ArgumentString string  `cli:"type:argument"`
		ArgumentInt    int     `cli:"type:argument"`
		ArgumentUint   uint    `cli:"type:argument"`
	}
)