package common

import "math/big"

var (
	Wei       = big.NewInt(1)
	KWei      = new(big.Int).Mul(Wei, big.NewInt(1e3))
	MWei      = new(big.Int).Mul(KWei, big.NewInt(1e3))
	GWei      = new(big.Int).Mul(MWei, big.NewInt(1e3))
	MicroSCOL = new(big.Int).Mul(GWei, big.NewInt(1e3))
	MilliSCOL = new(big.Int).Mul(MicroSCOL, big.NewInt(1e3))
	SCOL      = new(big.Int).Mul(MilliSCOL, big.NewInt(1e3))
	KScol     = new(big.Int).Mul(SCOL, big.NewInt(1e3))
	MScol     = new(big.Int).Mul(KScol, big.NewInt(1e3))
	GScol     = new(big.Int).Mul(MScol, big.NewInt(1e3))
	TScol     = new(big.Int).Mul(GScol, big.NewInt(1e3))
)

var Units = map[string]*big.Int{
	"wei":  Wei,
	"kwei": KWei, "babbage": KWei,
	"mwei": MWei, "lovelace": MWei,
	"gwei": GWei, "shannon": GWei,
	"microscol": MicroSCOL, "micro": MicroSCOL,
	"milliscol": MilliSCOL, "milli": MilliSCOL,
	"scol":  SCOL,
	"kscol": KScol, "grand": KScol,
	"mscol": MScol,
	"gscol": GScol,
	"tscol": TScol,
}
