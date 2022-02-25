<template>
	<view class="wrap">
		<view class="u-m-t-50 u-flex u-flex-col u-text-center">
			<u-avatar size="150" :src="avatarUrl"></u-avatar>
			<u-button size="medium" shape="circle" class="u-m-t-40"
				style="font-size:34rpx" @click="chooseAvatar">选择头像</u-button>
		</view>
		<u-form class="form" :model="model" :rules="rules" ref="uForm">
			<u-form-item label="昵称" prop="realname" label-width="120">
				<u-input placeholder="请输入用户昵称" v-model="model.realname" type="text"></u-input>
			</u-form-item>
			<u-form-item label="性别" prop="sex" label-width="120">
				<js-radio v-model="model.sex" dict-type="sys_user_sex"></js-radio>
			</u-form-item>
			<u-form-item label="邮箱" prop="email" label-width="120">
				<u-input placeholder="请输入电子邮箱" v-model="model.email" type="text"></u-input>
			</u-form-item>
			<u-form-item label="手机" prop="mobile" label-width="120">
				<u-input placeholder="请输入手机号码" v-model="model.mobile" type="number" maxlength="11"></u-input>
			</u-form-item>
			<u-form-item label="QQ" prop="qq" label-width="120">
				<u-input placeholder="请输入QQ" v-model="model.qq" type="text"></u-input>
			</u-form-item>
			<u-form-item label="地址" prop="address" label-width="120">
				<u-input type="textarea" placeholder="请输入你的地址" v-model="model.address" height="128" />
			</u-form-item>
			<u-form-item label="上次登录时间" label-width="250">
				{{this.$u.date(model.last_time,  'yyyy-mm-dd hh:MM:ss')}}
			</u-form-item>
			<u-form-item label="上次登录地址" label-width="250">
				{{model.last_ip}}
			</u-form-item>
		</u-form>
		<view class="form-footer">
			<u-button class="btn" type="primary" @click="submit">提交</u-button>
			<!-- <u-button class="btn" type="default" @click="cancel">关闭</u-button> -->
		</view>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	data() {
		return {
			model: {
				sex: '1'
			},
			rules: {
				realname: [
					{
						required: true,
						message: '请输入姓名',
						trigger: 'blur' ,
					},
					{
						min: 2,
						max: 32,
						message: '姓名长度在 2 到 32 个字符',
						trigger: ['change', 'blur'],
					},
					// {
					// 	// 此为同步验证，可以直接返回true或者false，如果是异步验证，稍微不同，见下方说明
					// 	validator: (rule, value, callback) => {
					// 		// 调用uView自带的js验证规则，详见：https://www.uviewui.com/js/test.html
					// 		return this.$u.test.chinese(value);
					// 	},
					// 	message: '姓名必须为中文',
					// 	// 触发器可以同时用blur和change，二者之间用英文逗号隔开
					// 	trigger: ['change', 'blur'],
					// },
					// {
					// 	// 异步验证，用途：比如用户注册时输入完账号，后端检查账号是否已存在
					// 	// 异步验证需要通过调用 callback()，并且在里面抛出 new Error()
					// 	// 抛出的内容为需要提示的信息，和其他方式的 message 属性的提示一样
					// 	asyncValidator: (rule, value, callback) => {
					// 		this.$u.post('/ebapi/public_api/index').then(res => {
					// 			if(res.error) {
					// 				// 如果验证出错，需要在callback()抛出new Error('错误提示信息')
					// 				callback(new Error('姓名重复'));
					// 			} else {
					// 				// 如果没有错误，也要执行 callback() 回调
					// 				callback();
					// 			}
					// 		})
					// 	},
					// 	trigger: ['blur'],
					// },
					// {
					// 	// 正则校验示例，此处用正则校验是否中文，此处仅为示例，因为uView有this.$u.test.chinese可以判断是否中文
					// 	pattern: /^[\u4e00-\u9fa5]+$/gi,
					// 	message: '简介只能为中文',
					// 	trigger: 'change',
					// },
				],
				mobile: [
					// {
					// 	required: true,
					// 	message: '请输入手机号',
					// 	trigger: ['change','blur'],
					// },
					{
						validator: (rule, value, callback) => {
							return value === '' || this.$u.test.mobile(value);
						},
						message: '手机号码不正确',
						trigger: ['change','blur'],
					}
				],
				// password: [
				// 	{
				// 		required: true,
				// 		message: '请输入密码',
				// 		trigger: ['change','blur'],
				// 	},
				// 	{
				// 		// 正则不能含有两边的引号
				// 		pattern: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]+\S{5,12}$/,
				// 		message: '需同时含有字母和数字，长度在6-12之间',
				// 		trigger: ['change','blur'],
				// 	}
				// ],
				// rePassword: [
				// 	{
				// 		required: true,
				// 		message: '请重新输入密码',
				// 		trigger: ['change','blur'],
				// 	},
				// 	{
				// 		validator: (rule, value, callback) => {
				// 			return value === this.model.password;
				// 		},
				// 		message: '两次输入的密码不相等',
				// 		trigger: ['change','blur'],
				// 	}
				// ],
			},
			avatarBase64: ''
		};
	},
	onLoad() {
		this.$u.api.memberProfile().then(res => {
			this.model = res.data.user;
		});
		uni.$on('uAvatarCropper', path => {
			this.avatarBase64 = path;
			if (this.avatarBase64 != '' && !this.avatarBase64.startsWith('data:')){
				// #ifdef APP-PLUS
				let self = this, fileUrl = this.avatarBase64;
				plus.io.resolveLocalFileSystemURL(path, function(entry) {
					entry.file(function(file) {
						var fileReader = new plus.io.FileReader()
						fileReader.onload = function(data) {
							// console.log(data.target.result);
							self.avatarBase64 = data.target.result;
						}
						fileReader.onerror = function(error) { }
						fileReader.readAsDataURL(file)
					}, function(error) { })
				}, function(error) { });
				// #endif
				// #ifndef APP-PLUS
				this.avatarBase64 = 'data:image/jpeg;base64,' + uni.getFileSystemManager()
						.readFileSync(this.avatarBase64, "base64");
				// #endif
			}
		})
	},
	computed: {
		avatarUrl() {
			if (this.avatarBase64 != ''){
				return this.avatarBase64;
			}
			let url = this.vuex_user.avatarUrl || '/ctxPath/static/images/user1.jpg';
			url = url.replace('/ctxPath/', this.vuex_config.baseUrl + '/');
			return url;
		}
	},
	onReady() {
		this.$refs.uForm.setRules(this.rules);
	},
	methods: {
		chooseAvatar() {
			this.$u.route({
				url: '/uview-ui/components/u-avatar-cropper/u-avatar-cropper',
				params: {
					destWidth: 800, // 输出图片宽高
					rectWidth: 200, // 裁剪框的宽高
					fileType: 'jpg', // 输出的图片类型，如果'png'类型发现裁剪的图片太大，改成"jpg"即可
				}
			})
		},
		submit() {
			this.$refs.uForm.validate(valid => {
				if (valid) {
					this.model.avatarBase64 = this.avatarBase64;
					this.$u.api.user.infoSaveBase(this.model).then(res => {
						this.$u.api.index(); // 保存后更新用户信息
						uni.showModal({
							title: '提示',
							content: res.message,
							showCancel: false,
							success: function () {
								if (res.result == 'true') {
									uni.navigateBack();
								}
							}
						});
					});
				} else {
					this.$u.toast('您填写的信息有误，请根据提示修正。');
				}
			});
		},
		cancel() {
			uni.navigateBack();
		}
	}
};
</script>
<style lang="scss">

</style>
