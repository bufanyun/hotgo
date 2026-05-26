import { VNode } from '@vue/runtime-core';

export type Attachment = {
  id: number;
  appId: string;
  memberId: number;
  cateId: number;
  drive: string;
  name: string;
  kind: string;
  metaType: string;
  naiveType: string;
  path: string;
  fileUrl: string;
  size: number;
  ext: string;
  md5: string;
  status: number;
  createdAt: string;
  updatedAt: string;
  sizeFormat: string;
};

export interface KindOption {
  key: string;
  label: any;
  icon: VNode;
  extra: any;
  disabled: boolean;
}

export type FileType = 'image' | 'doc' | 'audio' | 'video' | 'zip' | 'other' | 'default';

export function getFileType(fileType: string): string {
  switch (fileType) {
    case 'image':
      return '图片';
    case 'doc':
      return '文档';
    case 'audio':
      return '音频';
    case 'video':
      return '视频';
    case 'zip':
      return '压缩包';
    case 'other':
    case 'default':
      return '文件';
  }
  return '文件';
}
