/* eslint-disable */
import { TokenPairType } from "../dex/token_pair_type";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface PairMap {
  pairId: string;
  tokenPair: TokenPairType | undefined;
}

const basePairMap: object = { pairId: "" };

export const PairMap = {
  encode(message: PairMap, writer: Writer = Writer.create()): Writer {
    if (message.pairId !== "") {
      writer.uint32(10).string(message.pairId);
    }
    if (message.tokenPair !== undefined) {
      TokenPairType.encode(
        message.tokenPair,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PairMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePairMap } as PairMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairId = reader.string();
          break;
        case 2:
          message.tokenPair = TokenPairType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PairMap {
    const message = { ...basePairMap } as PairMap;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = String(object.pairId);
    } else {
      message.pairId = "";
    }
    if (object.tokenPair !== undefined && object.tokenPair !== null) {
      message.tokenPair = TokenPairType.fromJSON(object.tokenPair);
    } else {
      message.tokenPair = undefined;
    }
    return message;
  },

  toJSON(message: PairMap): unknown {
    const obj: any = {};
    message.pairId !== undefined && (obj.pairId = message.pairId);
    message.tokenPair !== undefined &&
      (obj.tokenPair = message.tokenPair
        ? TokenPairType.toJSON(message.tokenPair)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<PairMap>): PairMap {
    const message = { ...basePairMap } as PairMap;
    if (object.pairId !== undefined && object.pairId !== null) {
      message.pairId = object.pairId;
    } else {
      message.pairId = "";
    }
    if (object.tokenPair !== undefined && object.tokenPair !== null) {
      message.tokenPair = TokenPairType.fromPartial(object.tokenPair);
    } else {
      message.tokenPair = undefined;
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
