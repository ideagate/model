// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "dashboard/workflow.proto" (package "dashboard", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Workflow } from "../core/endpoint/workflow";
/**
 * @generated from protobuf message dashboard.CreateWorkflowRequest
 */
export interface CreateWorkflowRequest {
    /**
     * @generated from protobuf field: string project_id = 1;
     */
    projectId: string;
    /**
     * @generated from protobuf field: string application_id = 2;
     */
    applicationId: string;
    /**
     * @generated from protobuf field: string entrypoint_id = 3;
     */
    entrypointId: string;
    /**
     * New workflow will be copied from this version.
     * If not specified, then new workflow will be created from scratch.
     *
     * @generated from protobuf field: int64 from_version = 4;
     */
    fromVersion: bigint;
}
/**
 * @generated from protobuf message dashboard.CreateWorkflowResponse
 */
export interface CreateWorkflowResponse {
    /**
     * @generated from protobuf field: int64 version = 1;
     */
    version: bigint;
}
/**
 * @generated from protobuf message dashboard.GetWorkflowsRequest
 */
export interface GetWorkflowsRequest {
    /**
     * @generated from protobuf field: string project_id = 1;
     */
    projectId: string;
    /**
     * @generated from protobuf field: string application_id = 2;
     */
    applicationId: string;
    /**
     * @generated from protobuf field: string entrypoint_id = 3;
     */
    entrypointId: string;
    /**
     * @generated from protobuf field: int64 version = 4;
     */
    version: bigint;
}
/**
 * @generated from protobuf message dashboard.GetWorkflowsResponse
 */
export interface GetWorkflowsResponse {
    /**
     * @generated from protobuf field: repeated endpoint.Workflow workflows = 1;
     */
    workflows: Workflow[];
}
/**
 * @generated from protobuf message dashboard.UpdateWorkflowRequest
 */
export interface UpdateWorkflowRequest {
    /**
     * @generated from protobuf field: endpoint.Workflow workflow = 1;
     */
    workflow?: Workflow;
}
/**
 * @generated from protobuf message dashboard.UpdateWorkflowResponse
 */
export interface UpdateWorkflowResponse {
}
/**
 * @generated from protobuf message dashboard.DeleteWorkflowRequest
 */
export interface DeleteWorkflowRequest {
    /**
     * @generated from protobuf field: string project_id = 1;
     */
    projectId: string;
    /**
     * @generated from protobuf field: string application_id = 2;
     */
    applicationId: string;
    /**
     * @generated from protobuf field: string entrypoint_id = 3;
     */
    entrypointId: string;
    /**
     * @generated from protobuf field: int64 version = 4;
     */
    version: bigint;
}
/**
 * @generated from protobuf message dashboard.DeleteWorkflowResponse
 */
export interface DeleteWorkflowResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class CreateWorkflowRequest$Type extends MessageType<CreateWorkflowRequest> {
    constructor() {
        super("dashboard.CreateWorkflowRequest", [
            { no: 1, name: "project_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "application_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "entrypoint_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "from_version", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<CreateWorkflowRequest>): CreateWorkflowRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.projectId = "";
        message.applicationId = "";
        message.entrypointId = "";
        message.fromVersion = 0n;
        if (value !== undefined)
            reflectionMergePartial<CreateWorkflowRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateWorkflowRequest): CreateWorkflowRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string project_id */ 1:
                    message.projectId = reader.string();
                    break;
                case /* string application_id */ 2:
                    message.applicationId = reader.string();
                    break;
                case /* string entrypoint_id */ 3:
                    message.entrypointId = reader.string();
                    break;
                case /* int64 from_version */ 4:
                    message.fromVersion = reader.int64().toBigInt();
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
    internalBinaryWrite(message: CreateWorkflowRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string project_id = 1; */
        if (message.projectId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.projectId);
        /* string application_id = 2; */
        if (message.applicationId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.applicationId);
        /* string entrypoint_id = 3; */
        if (message.entrypointId !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.entrypointId);
        /* int64 from_version = 4; */
        if (message.fromVersion !== 0n)
            writer.tag(4, WireType.Varint).int64(message.fromVersion);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.CreateWorkflowRequest
 */
export const CreateWorkflowRequest = new CreateWorkflowRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateWorkflowResponse$Type extends MessageType<CreateWorkflowResponse> {
    constructor() {
        super("dashboard.CreateWorkflowResponse", [
            { no: 1, name: "version", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<CreateWorkflowResponse>): CreateWorkflowResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.version = 0n;
        if (value !== undefined)
            reflectionMergePartial<CreateWorkflowResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateWorkflowResponse): CreateWorkflowResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int64 version */ 1:
                    message.version = reader.int64().toBigInt();
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
    internalBinaryWrite(message: CreateWorkflowResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int64 version = 1; */
        if (message.version !== 0n)
            writer.tag(1, WireType.Varint).int64(message.version);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.CreateWorkflowResponse
 */
export const CreateWorkflowResponse = new CreateWorkflowResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetWorkflowsRequest$Type extends MessageType<GetWorkflowsRequest> {
    constructor() {
        super("dashboard.GetWorkflowsRequest", [
            { no: 1, name: "project_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "application_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "entrypoint_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "version", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<GetWorkflowsRequest>): GetWorkflowsRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.projectId = "";
        message.applicationId = "";
        message.entrypointId = "";
        message.version = 0n;
        if (value !== undefined)
            reflectionMergePartial<GetWorkflowsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetWorkflowsRequest): GetWorkflowsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string project_id */ 1:
                    message.projectId = reader.string();
                    break;
                case /* string application_id */ 2:
                    message.applicationId = reader.string();
                    break;
                case /* string entrypoint_id */ 3:
                    message.entrypointId = reader.string();
                    break;
                case /* int64 version */ 4:
                    message.version = reader.int64().toBigInt();
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
    internalBinaryWrite(message: GetWorkflowsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string project_id = 1; */
        if (message.projectId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.projectId);
        /* string application_id = 2; */
        if (message.applicationId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.applicationId);
        /* string entrypoint_id = 3; */
        if (message.entrypointId !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.entrypointId);
        /* int64 version = 4; */
        if (message.version !== 0n)
            writer.tag(4, WireType.Varint).int64(message.version);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.GetWorkflowsRequest
 */
export const GetWorkflowsRequest = new GetWorkflowsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetWorkflowsResponse$Type extends MessageType<GetWorkflowsResponse> {
    constructor() {
        super("dashboard.GetWorkflowsResponse", [
            { no: 1, name: "workflows", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Workflow }
        ]);
    }
    create(value?: PartialMessage<GetWorkflowsResponse>): GetWorkflowsResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.workflows = [];
        if (value !== undefined)
            reflectionMergePartial<GetWorkflowsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetWorkflowsResponse): GetWorkflowsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated endpoint.Workflow workflows */ 1:
                    message.workflows.push(Workflow.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: GetWorkflowsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated endpoint.Workflow workflows = 1; */
        for (let i = 0; i < message.workflows.length; i++)
            Workflow.internalBinaryWrite(message.workflows[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.GetWorkflowsResponse
 */
export const GetWorkflowsResponse = new GetWorkflowsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateWorkflowRequest$Type extends MessageType<UpdateWorkflowRequest> {
    constructor() {
        super("dashboard.UpdateWorkflowRequest", [
            { no: 1, name: "workflow", kind: "message", T: () => Workflow }
        ]);
    }
    create(value?: PartialMessage<UpdateWorkflowRequest>): UpdateWorkflowRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpdateWorkflowRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateWorkflowRequest): UpdateWorkflowRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* endpoint.Workflow workflow */ 1:
                    message.workflow = Workflow.internalBinaryRead(reader, reader.uint32(), options, message.workflow);
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
    internalBinaryWrite(message: UpdateWorkflowRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* endpoint.Workflow workflow = 1; */
        if (message.workflow)
            Workflow.internalBinaryWrite(message.workflow, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.UpdateWorkflowRequest
 */
export const UpdateWorkflowRequest = new UpdateWorkflowRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateWorkflowResponse$Type extends MessageType<UpdateWorkflowResponse> {
    constructor() {
        super("dashboard.UpdateWorkflowResponse", []);
    }
    create(value?: PartialMessage<UpdateWorkflowResponse>): UpdateWorkflowResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpdateWorkflowResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateWorkflowResponse): UpdateWorkflowResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: UpdateWorkflowResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.UpdateWorkflowResponse
 */
export const UpdateWorkflowResponse = new UpdateWorkflowResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteWorkflowRequest$Type extends MessageType<DeleteWorkflowRequest> {
    constructor() {
        super("dashboard.DeleteWorkflowRequest", [
            { no: 1, name: "project_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "application_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "entrypoint_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "version", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<DeleteWorkflowRequest>): DeleteWorkflowRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.projectId = "";
        message.applicationId = "";
        message.entrypointId = "";
        message.version = 0n;
        if (value !== undefined)
            reflectionMergePartial<DeleteWorkflowRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteWorkflowRequest): DeleteWorkflowRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string project_id */ 1:
                    message.projectId = reader.string();
                    break;
                case /* string application_id */ 2:
                    message.applicationId = reader.string();
                    break;
                case /* string entrypoint_id */ 3:
                    message.entrypointId = reader.string();
                    break;
                case /* int64 version */ 4:
                    message.version = reader.int64().toBigInt();
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
    internalBinaryWrite(message: DeleteWorkflowRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string project_id = 1; */
        if (message.projectId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.projectId);
        /* string application_id = 2; */
        if (message.applicationId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.applicationId);
        /* string entrypoint_id = 3; */
        if (message.entrypointId !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.entrypointId);
        /* int64 version = 4; */
        if (message.version !== 0n)
            writer.tag(4, WireType.Varint).int64(message.version);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.DeleteWorkflowRequest
 */
export const DeleteWorkflowRequest = new DeleteWorkflowRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteWorkflowResponse$Type extends MessageType<DeleteWorkflowResponse> {
    constructor() {
        super("dashboard.DeleteWorkflowResponse", []);
    }
    create(value?: PartialMessage<DeleteWorkflowResponse>): DeleteWorkflowResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<DeleteWorkflowResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteWorkflowResponse): DeleteWorkflowResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: DeleteWorkflowResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.DeleteWorkflowResponse
 */
export const DeleteWorkflowResponse = new DeleteWorkflowResponse$Type();
