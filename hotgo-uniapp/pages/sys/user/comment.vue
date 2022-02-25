<template>
	<view class="wrap">
		<u-form class="form" :model="model" :rules="rules" ref="uForm" label-position="top">
			<u-form-item label="分类：" prop="category" label-position="left" label-width="100">
				<js-select v-model="model.category" dict-type="app_comment_category"
						placeholder="请选择分类"></js-select>
			</u-form-item>
			<u-form-item label="问题和意见" prop="content">
				<u-input type="textarea" placeholder="请填写10个字以上的问题描述以便我们提供更好的帮助"
						v-model="model.content" height="200" maxlength="500" />
			</u-form-item>
			<u-form-item label="上传图片（选填，提供问题截图）" prop="images">
				<js-uploadfile v-model="model" biz-type="appComment_image"></js-uploadfile>
			</u-form-item>
			<u-form-item label="联系方式（手机、邮箱、QQ号码）" prop="contact">
				<u-input placeholder="选填，便于我们与你联系，进一步沟通"
						v-model="model.contact" type="text" maxlength="200"></u-input>
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
				id: '',
				category: '',
				content: '',
				contact: '',
				deviceInfo: ''
			},
			rules: {
				category: [
					{
						required: true,
						message: '请选择问题和意见的分类',
						trigger: ['change','blur'],
					}
				],
				content: [
					{
						required: true,
						min: 10, max: 500,
						message: '问题和意见在 10 到 500 个字符之间',
						trigger: ['change','blur'],
					}
				],
			}
		};
	},
	onReady() {
		this.$refs.uForm.setRules(this.rules);
		// 获取设备信息
		uni.getSystemInfo({
			success: res => {
				this.model.deviceInfo = JSON.stringify(res);
			}
		});
	},
	methods: {
		submit() {
			// console.log(this.model)
			this.$refs.uForm.validate(valid => {
				if (valid) {
					this.$u.api.commentSave(this.model).then(res => {
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
