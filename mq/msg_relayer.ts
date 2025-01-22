import { Msg } from './pub';


const Sub_RelayerAppIdNotify = "RelayerAppIdNotify"
const Sub_RelayerSpecifiedGasNotify = "RelayerSpecifiedGasNotify"
const Sub_RelayerAssignFeeNotify = "RelayerAssignFee"
///
let RelayerAppIVersion = 0
export function buildRelayerAppIdMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerAppIdNotify,
    opt: type,
    data: data,
    version:RelayerAppIVersion,
    jsonPatch: JSON.stringify(data),
  };
  RelayerAppIVersion = RelayerAppIVersion + 1;
  if (RelayerAppIVersion > 100000000) {
    RelayerAppIVersion = 0;
  }
  return msg;
}

let RelayerAssignFeeVersion = 0
export function buildRelayerAssignFeeMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerAssignFeeNotify,
    opt: type,
    data: data,
    version: RelayerAssignFeeVersion,
    jsonPatch: JSON.stringify(data),
  };

  RelayerAssignFeeVersion = RelayerAssignFeeVersion + 1;
  if (RelayerAssignFeeVersion > 100000000) {
    RelayerAssignFeeVersion = 0;
  }
  return msg;
}

let RelayerSpecifiedGasVersion = 0
export function buildRelayerSpecifiedGasMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerSpecifiedGasNotify,
    opt: type,
    data: data,
    version: RelayerSpecifiedGasVersion,
    jsonPatch: JSON.stringify(data),
  };
  RelayerSpecifiedGasVersion = RelayerSpecifiedGasVersion + 1;
  if (RelayerSpecifiedGasVersion > 100000000) {
    RelayerSpecifiedGasVersion = 0;
  }
  return msg;
}
