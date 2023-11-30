
type RiskCtrlMQ = {
    //ContractRule/ContractAbi
    kind: string,
    data: any,
}
type ContractNotice = {
    // 'add' | 'update' | 'delete'
    type: string,
    id: number,
    contractAddress: string,
    sceneNo: string,
}

export function buildContractRuleMQ(kind : string, type: string, id: number, contractAddress: string, sceneNo: string):RiskCtrlMQ {
    return {
        kind: "kind",
        data : {
            type: type,
            id: id,
            contractAddress: contractAddress,
            sceneNo: sceneNo,
        }
    }
}
export { RiskCtrlMQ }