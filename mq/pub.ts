
var RiskServerMQ = "RiskServerMQ";
var RiskEngineMQ = "RiskEngineMQ";

interface RiskCtrlMsg {
 subject: string;
 data: any;
}
interface RiskCtrMsqRsp{
    code: number;
    message: string;
}

 var RiskServerMQ_Subj_ContractRule = "ContractRule";
 var RiskServerMQ_Subj_ContractAbi = "ContractAbi";

 var RiskEngineMQ_Subj_RiskCtrlRule = "RiskCtrlRule";

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

 interface RiskEngineRuleStrNotice {
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
 function buildRiskCtrlMQRiskRule(subject: string, type: string,id:number, sceneNo:string, ruleName : string, ruleStr: string):RiskCtrlMsg {
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
////
export {
    RiskCtrMsqRsp,
    RiskCtrlMsg,
    ///
    RiskServerMQ,
    RiskEngineMQ,
    ///
    RiskServerMQ_Subj_ContractRule,
    RiskServerMQ_Subj_ContractAbi,
    //
    RiskEngineMQ_Subj_RiskCtrlRule,
    //
    buildRiskCtrlMQContract,
    buildRiskCtrlMQRiskRule,


    NoticeAdd,
    NoticeDelete,
    NoticeUpdate,
    NoticeVerify,

}