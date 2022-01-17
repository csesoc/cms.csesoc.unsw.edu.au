import { createAction } from "@reduxjs/toolkit";
import { FileEntity, Folder, File, sliceState } from './types';

/**
 * Payload Types
 */

export type RenamePayloadType = {
  id: number;
  newName: string;
}


/**
 * Init Actions
 */
export const initAction = createAction("folders/init");
// this one sets all children of current folder
export const initItemsAction = createAction<FileEntity[]>("folders/initItems");

/**
 * Directory Traversal actions
 */
export const traverseIntoFolder = createAction<number>("folders/traverseIntoFolder");
export const setDirectory = createAction<sliceState>("folders/setDirectory")

/**
 * CRUD actions
 */
export const addFolderItemAction = createAction<Folder>("folders/addFolderItem");
export const addFileItemAction = createAction<File>("folders/addFileItem");

// FileEntity = Folder | File
export const renameFileEntityAction = createAction<RenamePayloadType>("folders/renameFileEntity");

// TODO removeFolderItemAction
// TODO removeFileItemAction