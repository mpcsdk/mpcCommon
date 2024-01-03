
var RiskEngineMQ = "RiskEngineMQ"
var RiskEngineQueueMQ = "RiskEngineQueueMQ"
interface RiskCtrlMsg {
 subject: string;
 data: any;
}
interface RiskCtrMsqRsp{
    code: number;
    message: string;
}

var RiskEngineMQ_Subj_ContractRule = "ContractRule"
var RiskEngineMQ_Subj_ContractAbi  = "ContractAbi"
var RiskEngineMQ_Subj_RiskRule     = "RiskRule"
var RiskEngineQueueMQ_Subj_RiskRule = "RiskRule"

// //
// //RiskServerMQ

 var NoticeAdd = "add";
 var NoticeUpdate = "update";
 var NoticeDelete = "delete";

interface ContractNotice {
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

 var NoticeVerify = "verify";

 interface RiskCtrlRulesNotice {
 type: string;
 salience: number;
 ruleName: string;
 ruleStr: string;
 sceneNo: string;
 id: number;
}
function isValidRiskCtrlRulesNotice(s: RiskCtrlRulesNotice): boolean {
 return (
   s.type !== "" &&
   s.ruleStr !== "" &&
   s.sceneNo !== "" 
 );
}
///
///
 function buildRiskCtrlMQContract(subject : string, type: string, id: number, contractAddress: string, sceneNo: string):RiskCtrlMsg {
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
 function buildRiskCtrlMQRiskRule(
    subject: string, 
    type: string,
    id:number, 
    sceneNo:string, 
    ruleName : string, 
    ruleStr: string,
    salience: number):RiskCtrlMsg {
    let data : RiskCtrlRulesNotice = {
        type: type,
        salience:salience,
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
////
export {
    RiskCtrlMsg,
    RiskCtrMsqRsp,
    ///
    RiskEngineQueueMQ,
    RiskEngineMQ,
    ///
    RiskEngineMQ_Subj_ContractRule,
    RiskEngineMQ_Subj_ContractAbi,
    RiskEngineMQ_Subj_RiskRule,
    //
    RiskEngineQueueMQ_Subj_RiskRule,
    //
    buildRiskCtrlMQContract,
    buildRiskCtrlMQRiskRule,


    NoticeAdd,
    NoticeDelete,
    NoticeUpdate,
    NoticeVerify,

}