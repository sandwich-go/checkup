// Code generated by ProtoKitGo. DO NOT EDIT.
// package: internal_command
// file: internal_command/internal.proto

import * as jspb from "google-protobuf";
// Code generated by protokitgo. DO NOT EDIT.
export class CmdStream extends jspb.Message {
	getAddr(): string;
  	setAddr(value: string): void;

	getToken(): string;
  	setToken(value: string): void;

	getMetaMap(): jspb.Map<string, string>;
  	clearMetaMap(): void;

	serializeBinary(): Uint8Array;
  	toObject(includeInstance?: boolean): CmdStream.AsObject;
  	static toObject(includeInstance: boolean, msg: CmdStream): CmdStream.AsObject;
  	static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  	static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  	static serializeBinaryToWriter(message: CmdStream, writer: jspb.BinaryWriter): void;
  	static deserializeBinary(bytes: Uint8Array): CmdStream;
  	static deserializeBinaryFromReader(message: CmdStream, reader: jspb.BinaryReader): CmdStream;
}

export namespace CmdStream {
  	export type AsObject = {
    	addr: string,
    	token: string,
		metaMap: Array<[string, string]>,
   	}
}

