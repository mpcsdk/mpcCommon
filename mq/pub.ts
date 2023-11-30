
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

export function buildRiskCtrlMQ(kind : string, type: string, id: number, contractAddress: string, sceneNo: string):RiskCtrlMQ {
    let data : ContractNotice = {
        type: type,
        id: id,
        contractAddress: contractAddress,
        sceneNo: sceneNo,
    }
    return {
        kind: kind,
        data: data,
    }
}
export { RiskCtrlMQ }