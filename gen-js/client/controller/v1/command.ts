// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "client/controller/v1/command.proto" (package "v1", syntax proto3)
// tslint:disable
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Pod } from "./pod";
/**
 * @generated from protobuf message v1.GetListPodRequest
 */
export interface GetListPodRequest {
}
/**
 * @generated from protobuf message v1.GetListPodResponse
 */
export interface GetListPodResponse {
    /**
     * @generated from protobuf field: repeated v1.Pod pods = 1;
     */
    pods: Pod[];
}
// @generated message type with reflection information, may provide speed optimized methods
class GetListPodRequest$Type extends MessageType<GetListPodRequest> {
    constructor() {
        super("v1.GetListPodRequest", []);
    }
    create(value?: PartialMessage<GetListPodRequest>): GetListPodRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<GetListPodRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetListPodRequest): GetListPodRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetListPodRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetListPodRequest
 */
export const GetListPodRequest = new GetListPodRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetListPodResponse$Type extends MessageType<GetListPodResponse> {
    constructor() {
        super("v1.GetListPodResponse", [
            { no: 1, name: "pods", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Pod }
        ]);
    }
    create(value?: PartialMessage<GetListPodResponse>): GetListPodResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.pods = [];
        if (value !== undefined)
            reflectionMergePartial<GetListPodResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetListPodResponse): GetListPodResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated v1.Pod pods */ 1:
                    message.pods.push(Pod.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetListPodResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated v1.Pod pods = 1; */
        for (let i = 0; i < message.pods.length; i++)
            Pod.internalBinaryWrite(message.pods[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetListPodResponse
 */
export const GetListPodResponse = new GetListPodResponse$Type();
