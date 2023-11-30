type RiskCtrlMQ = {
    kind: string;
    data: any;
};
export declare function buildRiskCtrlMQ(kind: string, type: string, id: number, contractAddress: string, sceneNo: string): RiskCtrlMQ;
export { RiskCtrlMQ };
