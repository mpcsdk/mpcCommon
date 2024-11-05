
const Sub_ChainCfg = 'ChainCgf';
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
  data: any;
}
export function buildChainCfg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_ChainCfg,
    opt: type,
    data: data,
  };
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
export function buildContractAbi(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_ContractAbi,
    opt: type,
    data: data,
    // id: id,
    // contractAddress: contractAddress,
    // chainId: chainId,
  };
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

export function buildContractRule(
  type: string,
  data: any
  // id: number,
  // chainId: string,
  // contractAddress: string
): Msg {
  let msg: Msg = {
    sub: Sub_ContractRule,
    opt: type,
    data: data,
    // chainId: chainId,
    // contractAddress: contractAddress,
    // id: id,
  };
  return msg;
}
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
export function buildRiskCtrlRule(
  type: string,
  data: any
  // id: number,
  // isEnable: boolean
): Msg {
  let msg: Msg = {
    sub: Sub_RiskRule,
    opt: type,
    data: data,
    // id: id,
    // isEnable: isEnable,
  };
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
