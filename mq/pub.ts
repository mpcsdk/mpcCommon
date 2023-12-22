
export const RiskServerMQ = "RiskServerMQ";
export const RiskEngineMQ = "RiskEngineMQ";

export interface RiskCtrlMsg {
 subject: string;
 data: any;
}

export const RiskServerMQ_Subj_ContractRule = "ContractRule";
export const RiskServerMQ_Subj_ContractAbi = "ContractAbi";

export const RiskEngineMQ_Subj_RiskCtrlRule = "RiskCtrlRule";

// //
// //RiskServerMQ

export const NoticeAdd = "add";
export const NoticeUpdate = "update";
export const NoticeDelete = "delete";

export interface ContractNotice {
 type: string;
 id: number;
 contractAddress: string;
 sceneNo: string;
}

function isValidContractNotice(s: ContractNotice): boolean {
 return (
   s.type !== "" &&
   s.id > 0 &&
   s.contractAddress !== "" &&
   s.sceneNo !== ""
 );
}

// //
// //RiskEngineMQ

export const NoticeVerify = "verify";

export interface RiskEngineRuleStrNotice {
 type: string;
 ruleName: string;
 ruleStr: string;
 sceneNo: string;
 id: number;
}

function isValidRiskEngineRuleStrNotice(s: RiskEngineRuleStrNotice): boolean {
 return (
   s.type !== "" &&
   s.ruleStr !== "" &&
   s.sceneNo !== "" 
 );
}
///
///
export function buildRiskCtrlMQContract(subject : string, type: string, id: number, contractAddress: string, sceneNo: string):RiskCtrlMsg {
    let data : ContractNotice = {
        type: type,
        id: id,
        contractAddress: contractAddress,
        sceneNo: sceneNo,
    }
    return {
        subject: subject,
        data: data,
    }
}
export function buildRiskCtrlMQRiskRule(subject: string, type: string,id:number, ruleName : string,  sceneNo: string, ruleStr: string):RiskCtrlMsg {
    let data : RiskEngineRuleStrNotice = {
        type: type,
        sceneNo: sceneNo,
        ruleStr: ruleStr,
        ruleName: ruleName,
        id: id,
    }
    return {
        subject: subject,
        data: data,
    }
}