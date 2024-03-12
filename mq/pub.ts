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
// /Sub_ChainCfg
export class ChainCfgReq {
  sub: string;
  opt: string;
  id: number;
  chainId: number;
  coin: string;
  rpc: string;

  constructor(
    opt: string,
    id: number,
    chainId: number,
    coin: string,
    rpc: string
  ) {
    this.opt = opt;
    this.id = id;
    this.chainId = chainId;
    this.coin = coin;
    this.rpc = rpc;
  }
}
export function buildChainCfg(
  type: string,
  id: number,
  chainId: number,
  coin: string,
  rpc: string
): ChainCfgReq {
  let data: ChainCfgReq = {
    sub: Sub_ChainCfg,
    opt: type,
    id: id,
    chainId: chainId,
    coin: coin,
    rpc: rpc,
  };
  return data;
}
// ContractAbi
export interface ContractAbiReq {
  sub: string;
  opt: string;
  id: number;
  contractAddress: string;
  chainId: string;
}
export function buildContractAbi(
  type: string,
  id: number,
  contractAddress: string,
  chainId: string
): ContractAbiReq {
  let data: ContractAbiReq = {
    sub: Sub_ContractAbi,
    opt: type,
    id: id,
    contractAddress: contractAddress,
    chainId: chainId,
  };
  return data;
}
export function isValidContractAbiReq(s: ContractAbiReq): boolean {
  if (
    s.opt === '' ||
    s.id <= 0 ||
    s.contractAddress === '' ||
    s.chainId === ''
  ) {
    return false;
  }
  return true;
}

// /ContractRule
export interface ContractRuleReq {
  sub: string;
  // 'add' | 'update' | 'delete'
  opt: string;
  id: number;
  contractAddress: string;
  chainId: string;
}

export function buildContractRule(
  type: string,
  id: number,
  chainId: string,
  contractAddress: string
): ContractRuleReq {
  let data: ContractRuleReq = {
    sub: Sub_ContractRule,
    opt: type,
    chainId: chainId,
    contractAddress: contractAddress,
    id: id,
  };
  return data;
}
export function isValidContractRuleReq(s: ContractRuleReq): boolean {
  if (
    s.opt === '' ||
    s.id <= 0 ||
    s.contractAddress === '' ||
    s.chainId === ''
  ) {
    return false;
  }
  return true;
}

// /RiskRule
export interface RiskCtrlRuleReq {
  sub: string;
  //up/del/verify
  opt: string;
  id: number;
  isEnable: boolean;
}
export function buildRiskCtrlRule(
  type: string,
  id: number,
  isEnable: boolean
): RiskCtrlRuleReq {
  let data: RiskCtrlRuleReq = {
    sub: Sub_RiskRule,
    opt: type,
    id: id,
    isEnable: isEnable,
  };
  return data;
}
export function isValidRiskCtrlRuleReq(s: RiskCtrlRuleReq): boolean {
  // if (s.opt === '' || s.ruleStr === '' || s.sceneNo === '' || s.id <= 0) {
  //   return false;
  // }
  return true;
}

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

////
