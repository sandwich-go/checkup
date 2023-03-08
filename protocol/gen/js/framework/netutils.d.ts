// Code generated by ProtoKitGo. DO NOT EDIT.

import {Metadata} from "./metadata"
import {RawAny} from "../netutils/packet_pb";
import * as jspb from "google-protobuf";

export class NetMessage {
    setMessage(msg: jspb.Message): void
    getMessage(): jspb.Message
    setUri(uri: string): string
    getUri(): string
    setPassThrough(passThrough: string): void
    getPassThrough(): string
}

export class NetPacket {
    pushMessage(msg: jspb.Message): void
    getMessages(): Array<jspb.Message>
    pushNetMessage(msg: NetMessage): void
    getNetMessages(): Array<NetMessage>
    setMetadata(md: Metadata): void
    getMetadata(): Metadata
}

export class Netutils {
    static nextSequenceID(): number
    static nextPassThrough(): string
    static marshalNetPacketWithMessage(md: Metadata,...msgs: Array<NetMessage>): Uint8Array
    static marshalNetPacket(p: NetPacket): Uint8Array
    static encodeMessage(msg: NetMessage): RawAny
    static decodeMessage(uri: string, bytes: Uint8Array): NetMessage
    static unmarshalRawPacket(bytes: Uint8Array): NetPacket
}