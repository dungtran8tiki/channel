// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgOpenChannel } from "./types/channel/tx";
import { MsgCommitment } from "./types/channel/tx";
import { MsgFund } from "./types/channel/tx";
import { MsgWithdrawTimelock } from "./types/channel/tx";
import { MsgWithdrawHashlock } from "./types/channel/tx";
import { MsgCloseChannel } from "./types/channel/tx";


const types = [
  ["/channel.channel.MsgOpenChannel", MsgOpenChannel],
  ["/channel.channel.MsgCommitment", MsgCommitment],
  ["/channel.channel.MsgFund", MsgFund],
  ["/channel.channel.MsgWithdrawTimelock", MsgWithdrawTimelock],
  ["/channel.channel.MsgWithdrawHashlock", MsgWithdrawHashlock],
  ["/channel.channel.MsgCloseChannel", MsgCloseChannel],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgOpenChannel: (data: MsgOpenChannel): EncodeObject => ({ typeUrl: "/channel.channel.MsgOpenChannel", value: MsgOpenChannel.fromPartial( data ) }),
    msgCommitment: (data: MsgCommitment): EncodeObject => ({ typeUrl: "/channel.channel.MsgCommitment", value: MsgCommitment.fromPartial( data ) }),
    msgFund: (data: MsgFund): EncodeObject => ({ typeUrl: "/channel.channel.MsgFund", value: MsgFund.fromPartial( data ) }),
    msgWithdrawTimelock: (data: MsgWithdrawTimelock): EncodeObject => ({ typeUrl: "/channel.channel.MsgWithdrawTimelock", value: MsgWithdrawTimelock.fromPartial( data ) }),
    msgWithdrawHashlock: (data: MsgWithdrawHashlock): EncodeObject => ({ typeUrl: "/channel.channel.MsgWithdrawHashlock", value: MsgWithdrawHashlock.fromPartial( data ) }),
    msgCloseChannel: (data: MsgCloseChannel): EncodeObject => ({ typeUrl: "/channel.channel.MsgCloseChannel", value: MsgCloseChannel.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
