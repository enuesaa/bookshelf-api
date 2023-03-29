// @generated by protoc-gen-connect-web v0.8.4 with parameter "target=ts,import_extension="
// @generated from file v1/tag.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddTagItemRequest, AddTagItemResponse, ListTagItemsRequest, ListTagItemsResponse, ListTagsRequest, ListTagsResponse, RemoveTagItemRequest, RemoveTagItemResponse } from "./tag_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service v1.Tag
 */
export const Tag = {
  typeName: "v1.Tag",
  methods: {
    /**
     * @generated from rpc v1.Tag.ListTags
     */
    listTags: {
      name: "ListTags",
      I: ListTagsRequest,
      O: ListTagsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Tag.AddTagItem
     */
    addTagItem: {
      name: "AddTagItem",
      I: AddTagItemRequest,
      O: AddTagItemResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Tag.ListTagItems
     */
    listTagItems: {
      name: "ListTagItems",
      I: ListTagItemsRequest,
      O: ListTagItemsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Tag.RemoveTagItem
     */
    removeTagItem: {
      name: "RemoveTagItem",
      I: RemoveTagItemRequest,
      O: RemoveTagItemResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
