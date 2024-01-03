package mpcmodel

import (
	"math/big"
)

type ContractRule struct {
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`

	MethodName       string `json:"methodName"`
	MethodSig        string `json:"methodSig"`
	MethodFromField  string `json:"methodFromField"`
	MethodToField    string `json:"methodToField"`
	MethodValueField string `json:"methodValueField"`

	EventName       string `json:"eventName"`
	EventSig        string `json:"eventSig"`
	EventTopic      string `json:"eventTopic"`
	EventFromField  string `json:"eventFromField"`
	EventToField    string `json:"eventToField"`
	EventValueField string `json:"eventValueField"`

	WhiteAddrList []string `json:"skipToAddr"`
	Threshold     *big.Int `json:"threshold"`
	ThresholdNft  int64    `jons:"thresholdNft"`
	///
}
