
const Sub_ChainCfg = 'ChainCfg';
const Sub_ContractAbi = 'ContractCfg';
const Sub_ContractRule = 'ContractRule';
const Sub_RiskRule = 'RiskRule';

const Sub_RiskRuleReply = 'RiskRuleReply';
///
var OptAdd = 'add';
var OptUpdate = 'update';
var OptDelete = 'delete';
var OptCheck = 'check';
export {
  Sub_ChainCfg,
  Sub_ContractAbi,
  Sub_ContractRule,
  Sub_RiskRule,
  Sub_RiskRuleReply,
  OptAdd,
  OptUpdate,
  OptDelete,
  OptCheck,
};

///
export interface Msg {
  sub: string;
  opt: string;
  version: number;
  data: any;
  jsonPatch: string;
}
function buildMsg(sub:string, type:string, version :  number, data: any) : Msg {
  return {
    sub: sub,
    opt: type,
    version: version,
    data: data,
    jsonPatch: JSON.stringify(data),
  };
}
let chainCfgVersion = 0

export function buildChainCfg(type: string, data: any): Msg {
  // let msg: Msg = {
  //   sub: Sub_ChainCfg,
  //   opt: type,
  //   version: chainCfgVersion,
  //   data: data,
  //   jsonPatch: JSON.stringify(data),
  // };
  let msg = buildMsg(Sub_ChainCfg, type, chainCfgVersion, data)
  chainCfgVersion = chainCfgVersion + 1;
  if (chainCfgVersion > 100000000) {
    chainCfgVersion = 0;
  }
  return msg;
}

// ContractAbi
// export interface ContractAbiReq {
//   sub: string;
//   opt: string;
//   id: number;
//   contractAddress: string;
//   chainId: string;
// }
let contractAbiVersion = 0;
export function buildContractAbi(type: string, data: any): Msg {
  // let msg: Msg = {
  //   sub: Sub_ContractAbi,
  //   opt: type,
  //   data: data,
  //   // id: id,
  //   // contractAddress: contractAddress,
  //   // chainId: chainId,
  // };
  let msg = buildMsg(Sub_ContractAbi, type, contractAbiVersion, data)
  contractAbiVersion = contractAbiVersion + 1;
  if (contractAbiVersion > 100000000) {
    contractAbiVersion = 0;
  }
  return msg;
}
// export function isValidContractAbiReq(s: Msg): boolean {
//   if (
//     s.opt === '' ||
//     s.id <= 0 ||
//     s.contractAddress === '' ||
//     s.chainId === ''
//   ) {
//     return false;
//   }
//   return true;
// }

// /ContractRule
// export interface ContractRuleReq {
//   sub: string;
//   // 'add' | 'update' | 'delete'
//   opt: string;
//   id: number;
//   contractAddress: string;
//   chainId: string;
// }

// export function buildContractRule(
//   type: string,
//   data: any
//   // id: number,
//   // chainId: string,
//   // contractAddress: string
// ): Msg {
//   // let msg: Msg = {
//   //   sub: Sub_ContractRule,
//   //   opt: type,
//   //   data: data,
//   //   // chainId: chainId,
//   //   // contractAddress: contractAddress,
//   //   // id: id,
//   // };
//   let msg = buildMsg(Sub_ContractAbi, type, data, contractAbiVersion)
//   contractAbiVersion = contractAbiVersion + 1;
//   if (contractAbiVersion > 100000000) {
//     contractAbiVersion = 0;
//   }
//   return msg;
// }
// export function isValidContractRuleReq(s: ContractRuleReq): boolean {
//   if (
//     s.opt === '' ||
//     s.id <= 0 ||
//     s.contractAddress === '' ||
//     s.chainId === ''
//   ) {
//     return false;
//   }
//   return true;
// }

// /RiskRule
// export interface RiskCtrlRuleReq {
//   sub: string;
//   //up/del/verify
//   opt: string;
//   id: number;
//   isEnable: boolean;
// }
let riskCtrlRuleVersion = 0;
export function buildRiskCtrlRule(
  type: string,
  data: any
): Msg {
  let msg = buildMsg(Sub_RiskRule, type, riskCtrlRuleVersion, data)
  riskCtrlRuleVersion = riskCtrlRuleVersion + 1;
  if (riskCtrlRuleVersion > 100000000) {
    riskCtrlRuleVersion = 0;
  }
  return msg;
}
// export function isValidRiskCtrlRuleReq(s: RiskCtrlRuleReq): boolean {
//   // if (s.opt === '' || s.ruleStr === '' || s.sceneNo === '' || s.id <= 0) {
//   //   return false;
//   // }
//   return true;
// }

// /RiskRuleReply
export interface RiskRuleReplyReq {
  sub: string;
  opt: string;
  chainId: string;
  ruleName: string;
  ruleStr: string;
}
export function buildRiskRuleReply(
  chainId: string,
  ruleStr: string,
  ruleName: string
): RiskRuleReplyReq {
  let data: RiskRuleReplyReq = {
    sub: Sub_RiskRuleReply,
    opt: OptCheck,
    chainId: chainId,
    ruleStr: ruleStr,
    ruleName: ruleName,
  };
  return data;
}
export interface RiskRuleReplyRes {
  Code: number;
  Msg: string;
}
///
///
///
