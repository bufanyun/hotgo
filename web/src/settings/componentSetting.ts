export default {
  table: {
    apiSetting: {
      // 当前页的字段名
      pageField: 'page',
      // 每页数量字段名
      sizeField: 'pageSize',
      // 接口返回的数据字段名
      listField: 'list',
      // 接口返回总页数字段名
      totalField: 'pageCount',
      // 接口返回总行数字段名
      itemCountField: 'totalCount',
    },
    //默认分页数量
    defaultPageSize: 10,
    //可切换每页数量集合
    pageSizes: [10, 15, 20, 30, 50, 100],
  },
  upload: {
    //考虑接口规范不同
    apiSetting: {
      // 集合字段名
      infoField: 'data',
      // 图片地址字段名
      imgField: 'fileUrl',
    },
    //最大上传图片大小
    maxSize: 10,
    //图片上传类型
    imageType: ['image/png', 'image/jpg', 'image/jpeg', 'image/gif', 'image/svg+xml', 'image/webp'],
    //文件上传类型
    fileType: [
      // 图片
      'image/png',
      'image/jpg',
      'image/jpeg',
      'image/gif',
      'image/svg+xml',
      'image/webp',
      // 文档
      'application/msword',
      'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
      'application/vnd.ms-excel',
      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
      'application/vnd.ms-powerpoint',
      'application/vnd.openxmlformats-officedocument.presentationml.presentation',
      // 音频
      'audio/mpeg',
      'audio/midi',
      // 视频
      'audio/mp4',
      'video/webm',
      'video/x-flv',
    ],
  },
};
