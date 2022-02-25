<template>
	<view class="u-flex-1">
		<u-upload width="160" height="160" ref="uUpload"
			:action="options.action"
			:header="options.header"
			:form-data="options.formData"
			:name="options.name"
			:max-count="maxCount"
			:auto-upload="true"
			:deletable="deletable"
			:before-upload="beforeUpload"
			@on-success="uploadSuccess"
			@on-uploaded="uploadUploaded"
			:before-remove="beforeRemove"
			@on-remove="uploadRemove"
			></u-upload>
	</view>
</template>
<script>
import SparkMD5 from '@/common/spark-md5.js';
/**
 * 文件上传组件组件
 * @property {Object} value 使用 v-model="this.model" 指定表单的 model 对象（存放文件上传编号信息）
 * @property {String} bizKey 业务表的主键值（与附件关联的业务数据）【可选】如果不设置，则获取 value.id 作为主键
 * @property {String} bizType 业务表的上传类型（全网唯一，推荐格式：实体名_上传类型，例如，意见反馈图片：appComment_image）
 * @property {String} uploadType 上传文件类型：image，目前移动端仅支持上传图片
 * @property {String} imageMaxWidth 图片压缩，最大宽度（uploadType为image生效），设置-1代表不做任何处理
 * @property {String} imageMaxHeight 图片压缩，最大宽度（uploadType为image生效），设置-1代表不做任何处理
 * @property {String} maxCount 最大上传个数，默认 52 个，如果设置为 0 可以当做【只读模式】使用
 * @example <js-uploadfile v-model="model.otherData" :biz-key="model.id" biz-type="testData_image"></js-uploadfile>
 * @description Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 * @author ThinkGem
 * @version 2021-3-11
 */
export default {
	props: {
		value: {
			type: Object,
			default() {
				return {}
			}
		},
		bizKey: {
			type: String,
			default: ''
		},
		bizType: {
			type: String,
			default: 'images'
		},
		uploadType: {
			type: String,
			default: 'image'
		},
		imageMaxWidth: {
			type: [String, Number],
			default: 1024
		},
		imageMaxHeight: {
			type: [String, Number],
			default: 768
		},
		maxCount: {
			type: [String, Number],
			default: 52
		}
	},
	data() {
		return {
			options: {
				value: {},
				action: '',
				header: {},
				formData: {
					fileMd5: '',
					fileName: '',
					bizKey: this.bizKey || (this.value && this.value.id) || '',
					bizType: this.bizType,
					uploadType: this.uploadType,
					imageMaxWidth: this.imageMaxWidth,
					imageMaxHeight: this.imageMaxHeight
				},
				name: 'file',
				// 文件上传的 id 数组
				fileUploadIds: [],
				// 文件删除的 id 数组
				fileUploadDelIds: []
			},
			deletable: true
		};
	},
	watch: {
		value(val, oldVal) {
			this.options.value = val;
		},
		maxCount(val, oldVal) {
			this.refreshStatus();
		},
		bizKey(val, oldVal) {
			this.options.formData.bizKey = val;
			this.loadData();
		}
	},
	created() {
		this.refreshStatus();
		this.options.action = this.vuex_config.baseUrl + this.vuex_config.adminPath + '/file/upload';
		this.options.formData = Object.assign(this.options.formData, this.formData);
		this.loadData();
	},
	methods: {
		// 刷新是否只读状态
		refreshStatus(){
			if (this.maxCount <= 0){
				this.deletable = false;
			}
		},
		// 已上传的文件回显到上传组件
		loadData(){
			if (this.options.formData.bizKey != ''){
				let baseUrl = this.vuex_config.baseUrl;
				let adminPath = this.vuex_config.adminPath;
				this.$u.post(adminPath + '/file/fileList', {
					bizKey: this.options.formData.bizKey,
					bizType: this.options.formData.bizType,
				}).then(res => {
					let lists = [];
					if (!(typeof res === 'object' && (res.result === 'login' || res.result === 'false'))){
						for (let i in res){
							let f = res[i];
							lists.push({
								url: baseUrl + f.fileUrl,
								fileUploadId: f.id,
								progress: 100,
								error: false
							});
						}
					}
					// console.log(lists)
					this.$refs.uUpload.lists = lists;
					this.uploadRefreshIds(lists);
				});
			}
		},
		// 上传之前，验证秒传、是否继续上传等
		beforeUpload(index, lists) {
			let self = this;
			let item = lists[index];
			let upload = this.upload;
			let formData = this.options.formData;
			let baseUrl = this.vuex_config.baseUrl;
			let adminPath = this.vuex_config.adminPath;
			self.$u.http.interceptor.request(this.options);
			return new Promise((resolve, reject) => {
				try{
					function uploadFile(arrayBuffer){
						let buffer = arrayBuffer;
						let size = 10 * 1024 * 1024;
						let spark = new SparkMD5.ArrayBuffer();
						if (buffer.byteLength > size){
							spark.append(buffer.slice(0, size));;
						}else{
							spark.append(buffer);
						}
						formData.fileEntityId = '';
						formData.fileUploadId = '';
						formData.fileMd5 = spark.end();
						formData.fileName = item.file.name;
						// console.log('formData' + JSON.stringify(formData));
						self.$u.post(adminPath + '/file/upload', formData).then(res => {
							// console.log(res)
							// 文件已经上传，启用秒传
							if (res.result == 'true' && res.fileUpload){
								item.fileUploadId = res.fileUpload.id;
								item.progress = 100;
								item.error = false;
								reject(res);
							}
							// 文件未上传过，继续上传文件
							else if (res.fileUploadId && res.fileEntityId){
								formData.fileUploadId = res.fileUploadId;
								formData.fileEntityId = res.fileEntityId;
								item.fileUploadId = res.fileUploadId;
								resolve();
							}
							// 未知错误，提示服务端返回的信息
							else {
								uni.showModal({title: '提示', content: res.message });
								reject(res);
							}
						}).catch(err => {
							console.error(err);
							reject(err);
						})
					}
					// #ifdef APP-PLUS
					plus.io.requestFileSystem(plus.io.PRIVATE_WWW, function(fs){
						fs.root.getFile(item.url, {create: false}, function(fileEntry){
							fileEntry.file(function(file){
								// console.log("getFile:" + JSON.stringify(file))
								item.file.name = file.name;
								var fileReader = new plus.io.FileReader();
								fileReader.readAsText(file, 'utf-8');
								fileReader.onloadend = function(evt) {
									uploadFile(evt.target.result);
								}
								fileReader.onerror = function(error) {
									reject(error);
								}
							}, reject);
						}, reject);
					} );
					// #endif
					// #ifndef APP-PLUS
					uni.request({
						url: item.url,
						responseType: 'arraybuffer',
						complete: res => {
							// console.log(res)
							if (res.statusCode == 200) {
								uploadFile(res.data);
							}else{
								reject(res);
							}
						}
					})
					// #endif
				}catch(err){
					console.error(err);
					reject(err);
				}
			})
		},
		// 上传成功一个，就写进 fileUploadIds
		uploadSuccess(data, index, lists, name){
			let item = lists[index];
			this.options.fileUploadIds.push(item.fileUploadId);
		},
		// 全部上传后，刷新 fileUploadIds、fileUploadDelIds
		uploadUploaded(lists, name) {
			this.uploadRefreshIds(lists);
		},
		// 移除之前获取删除的 fileUploadId，写进 fileUploadDelIds
		beforeRemove(index, lists){
			let item = lists[index];
			if (item.fileUploadId){
				this.options.fileUploadDelIds.push(item.fileUploadId);
			}
			return true;
		},
		// 移除之后，刷新 fileUploadIds、fileUploadDelIds
		uploadRemove(index, lists){
			this.uploadRefreshIds(lists);
		},
		// 刷新 fileUploadIds、fileUploadDelIds
		uploadRefreshIds(lists, name) {
			let fileUploadIds = [];
			lists.forEach(item => {
				if (item.fileUploadId && item.progress == 100){
					fileUploadIds.push(item.fileUploadId);
				}
			});
			this.options.fileUploadIds = fileUploadIds;
			// console.log('fileUploadIds', this.options.fileUploadIds)
			// console.log('fileUploadDelIds', this.options.fileUploadDelIds)
			// 将上传和删除的 id 回传给 model
			let formData = this.options.formData;
			let fileParams = this.options.value || {};
			fileParams[formData.bizType] = this.options.fileUploadIds.join(',');
			fileParams[formData.bizType+'__del'] = this.options.fileUploadDelIds.join(',');
			this.options.value = fileParams;
			this.$emit('input', Object.assign(this.options.value, this.value));
		},
	}
}
</script>
<style lang="scss" scoped>
	
</style>
