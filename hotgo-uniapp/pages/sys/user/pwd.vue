<template>
	<view class="wrap">
		<u-form class="form" :model="model" :rules="rules" ref="uForm">
			<u-form-item label="旧密码" prop="oldPassword" label-width="180">
				<u-input type="password" v-model="model.oldPassword" placeholder="请输入旧密码"></u-input>
			</u-form-item>
			<u-form-item label="新密码" prop="newPassword" label-width="180">
				<u-input type="password" v-model="model.newPassword" placeholder="请输入新密码"></u-input>
			</u-form-item>
			<u-form-item label="确认密码" prop="confirmNewPassword" label-width="180">
				<u-input type="password" v-model="model.confirmNewPassword" placeholder="请确认新密码"></u-input>
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
import base64 from '@/common/base64.js';
export default {
	data() {
		return {
			model: {
				oldPassword: '',
				newPassword: '',
				confirmNewPassword: ''
			},
			rules: {
				oldPassword: [
					{
						required: true,
						message: '请输入旧密码',
						trigger: ['change','blur'],
					}
				],
				newPassword: [
					{
						required: true,
						message: '请输入新密码',
						trigger: ['change','blur'],
					},
					{
						pattern: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]+\S{5,12}$/,
						message: '需同时含有字母和数字，长度在6-12之间',
						trigger: ['change','blur'],
					}
				],
				confirmNewPassword: [
					{
						required: true,
						message: '请重新输入密码',
						trigger: ['change','blur'],
					},
					{
						validator: (rule, value, callback) => {
							return value === this.model.newPassword;
						},
						message: '两次输入的密码不相等',
						trigger: ['change','blur'],
					}
				],
			}
		};
	},
	onReady() {
		this.$refs.uForm.setRules(this.rules);
	},
	methods: {
		submit() {
			this.$refs.uForm.validate(valid => {
				if (valid) {
					this.$u.api.user.infoSavePwd({
						oldPassword: base64.btoa(this.model.oldPassword),
						newPassword: base64.btoa(this.model.newPassword),
						confirmNewPassword: base64.btoa(this.model.confirmNewPassword)
					}).then(res => {
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
