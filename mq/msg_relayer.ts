import { Msg } from './pub';


const Sub_RelayerFeeNotify = "RelayerFeeNotify"
const Sub_RelayerChannelNotify = "RelayerChannelNotify"
///

export function buildRelayerFeeMsg(type: string, data: any): Msg {
  let msg: Msg = {
    sub: Sub_RelayerFeeNotify,
    opt: type,
    data: data,
  };
  return msg;
}
