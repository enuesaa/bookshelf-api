// @generated by protoc-gen-connect-web v0.8.5 with parameter "target=ts,import_extension="
// @generated from file v1/board.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddBoardRequest, AddBoardResponse, ArchiveBoardRequest, ArchiveBoardResponse, CheckinRequest, CheckinResponse, DeleteBoardRequest, DeleteBoardResponse, GetBoardRequest, GetBoardResponse, ListBoardsRequest, ListBoardsResponse, ListTimelineRequest, ListTimelineResponse, UnArchiveBoardRequest, UnArchiveBoardResponse, UpdateBoardRequest, UpdateBoardResponse } from "./board_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service v1.Board
 */
export const Board = {
  typeName: "v1.Board",
  methods: {
    /**
     * @generated from rpc v1.Board.ListBoards
     */
    listBoards: {
      name: "ListBoards",
      I: ListBoardsRequest,
      O: ListBoardsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.GetBoard
     */
    getBoard: {
      name: "GetBoard",
      I: GetBoardRequest,
      O: GetBoardResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.AddBoard
     */
    addBoard: {
      name: "AddBoard",
      I: AddBoardRequest,
      O: AddBoardResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.UpdateBoard
     */
    updateBoard: {
      name: "UpdateBoard",
      I: UpdateBoardRequest,
      O: UpdateBoardResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.DeleteBoard
     */
    deleteBoard: {
      name: "DeleteBoard",
      I: DeleteBoardRequest,
      O: DeleteBoardResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.Checkin
     */
    checkin: {
      name: "Checkin",
      I: CheckinRequest,
      O: CheckinResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.ListTimeline
     */
    listTimeline: {
      name: "ListTimeline",
      I: ListTimelineRequest,
      O: ListTimelineResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.ArchiveBoard
     */
    archiveBoard: {
      name: "ArchiveBoard",
      I: ArchiveBoardRequest,
      O: ArchiveBoardResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc v1.Board.UnArchiveBoard
     */
    unArchiveBoard: {
      name: "UnArchiveBoard",
      I: UnArchiveBoardRequest,
      O: UnArchiveBoardResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

