import { Msg } from './pub';


const Sub_RelayerFeeNotify = "RelayerFeeNotify"
const Sub_RelayerSpecifiedGasNotify = "RelayerSpecifiedGasNotify"
const Sub_RelayerAssignFeeNotify = "RelayerAssignFee"
///

export function buildRelayerAppIdMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerFeeNotify,
    opt: type,
    data: data,
  };
  return msg;
}
export function buildRelayerAssignFeeMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerAssignFeeNotify,
    opt: type,
    data: data,
  };
  return msg;
}
export function buildRelayerSpecifiedGasMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerSpecifiedGasNotify,
    opt: type,
    data: data,
  };
  return msg;
}
