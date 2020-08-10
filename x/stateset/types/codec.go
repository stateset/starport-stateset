package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding
		cdc.RegisterConcrete(MsgCreateCase{}, "stateset/CreateCase", nil)
		cdc.RegisterConcrete(MsgCreateLoan{}, "stateset/CreateLoan", nil)
		cdc.RegisterConcrete(MsgCreateContact{}, "stateset/CreateContact", nil)
		cdc.RegisterConcrete(MsgCreateInvoice{}, "stateset/CreateInvoice", nil)
		cdc.RegisterConcrete(MsgCreatePurchaseorder{}, "stateset/CreatePurchaseorder", nil)
		cdc.RegisterConcrete(MsgCreateAgreement{}, "stateset/CreateAgreement", nil)
		cdc.RegisterConcrete(MsgCreateProposal{}, "stateset/CreateProposal", nil)
		cdc.RegisterConcrete(MsgCreateUser{}, "stateset/CreateUser", nil)
		cdc.RegisterConcrete(MsgCreateNumber{}, "stateset/CreateNumber", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
