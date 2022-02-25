<template>
	<view class="wrap">
		<view class="list">
			<view class="list-call">
				<u-icon class="u-icon" size="40" name="account"></u-icon>
				<input class="u-input" type="text" v-model="loginCode" maxlength="32" placeholder="请输入登录账号" />
			</view>
			<view class="list-call">
				<u-icon class="u-icon" size="40" name="chat"></u-icon>
				<input class="u-input" type="text" v-model="userName" maxlength="20" placeholder="请输入用户昵称" />
			</view>
			<!-- <view class="list-call">
				<u-icon class="u-icon" size="40" name="email"></u-icon>
				<input class="u-input" type="text" v-model="email" placeholder="请输入电子邮箱" />
			</view> -->
			<view class="list-call">
				<u-icon class="u-icon" size="40" name="phone"></u-icon>
				<input class="u-input" type="text" v-model="mobile" maxlength="11" placeholder="请输入手机号码" />
			</view>
			<view class="list-call">
				<u-icon class="u-icon" size="40" name="lock"></u-icon>
				<input class="u-input" type="text" v-model="password" maxlength="32" placeholder="请输入登录密码" :password="!showPassword" />
				<image class="u-icon-right" :src="'/static/jeesite/login/eye_' + (showPassword ? 'open' : 'close') + '.png'" @click="showPass()"></image>
			</view>
			<view class="list-call">
				<u-icon class="u-icon" size="40" name="bookmark"></u-icon>
				<input class="u-input" type="text" v-model="validCode" maxlength="4" placeholder="图片验证码" />
				<u-image class="img-valid-code" width="300rpx" height="90rpx" :src="imgValidCodeSrc" @click="refreshImgValidCode()"></u-image>
			</view>
			<view class="list-call">
				<u-icon class="u-icon" size="40" name="coupon"></u-icon>
				<input class="u-input" type="text" v-model="fpValidCode" maxlength="6" placeholder="手机验证码" />
				<u-verification-code ref="uCode" :seconds="seconds" @change="codeChange"></u-verification-code>
				<view class="btn-valid-code" :class="{ 'btn-valid-codes': tips != '获取验证码' && tips != '重新获取' }"
					hover-class="btn-valid-code-hover" @click="getValidCode()">{{tips}}</view>
			</view>
		</view>
		<view class="agreement">
			<u-checkbox v-model="terms">我已阅读并同意</u-checkbox>
			<navigator url="/pages/common/webview?title=软件用户协议&url=https://jeesite.com/docs/support/"
				open-type="navigate">《软件用户协议》</navigator>
		</view>
		<view class="button" hover-class="button-hover" @click="submit()"><text>注册账号</text></view>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	data() {
		return {
			loginCode: '',
			userName: '',
			email: '',
			mobile: '',
			userType: 'member',
			terms: false,
			password: '',
			validCode: '',
			fpValidCode: '',
			showPassword: false,
			imgValidCodeSrc: null,
			tips: '获取验证码',
			seconds: 60
		};
	},
	onLoad() {
		this.refreshImgValidCode();
	},
	methods: {
		showPass() {
			this.showPassword = !this.showPassword;
		},
		refreshImgValidCode(e) {
			if (this.vuex_token == '') {
				this.$u.api.index().then(res => {
					this.imgValidCodeSrc = this.vuex_config.baseUrl + '/validCode?__sid='
						+ res.sessionid + '&t=' + new Date().getTime();
				});
			} else {
				this.imgValidCodeSrc = this.vuex_config.baseUrl + '/validCode?__sid='
						+ this.vuex_token + '&t=' + new Date().getTime();
			}
			this.validCode = '';
		},
		codeChange(text) {
			this.tips = text;
		},
		formValid() {
			if (this.loginCode.length == 0) {
				this.$u.toast('请输入账号');
				return false;
			}
			if (this.userName.length == 0) {
				this.$u.toast('请输入昵称');
				return false;
			}
			if (this.password.length == 0) {
				this.$u.toast('请输入密码');
				return false;
			}
			if (this.validCode.length == 0) {
				this.$u.toast('请输入图片验证码');
				return false;
			}
			return true;
		},
		getValidCode() {
			if (!this.formValid()) {
				return;
			}
			if (this.$refs.uCode.canGetCode) {
				this.$u.api.validCode({
					validCode: this.validCode
				})
				.then(res => {
					if (res !== 'true') {
						this.$u.toast('图片验证码错误');
						return;
					}
					this.$u.api.getRegValidCode({
						loginCode: this.loginCode,
						userName: this.userName,
						email: this.email,
						mobile: this.mobile,
						userType: this.userType,
						validCode: this.validCode,
						validType: 'mobile'
					})
					.then(res => {
						this.$u.toast(res.message, 3000);
						if (res.result == 'false') {
							this.refreshImgValidCode();
						}
					});
					this.$refs.uCode.start();
				});
			}
		},
		submit() {
			if (!this.formValid()) {
				return;
			}
			if (this.terms != true){
				this.$u.toast('请阅读《用户使用协议》');
				return false;
			}
			if (this.fpValidCode.length == 0) {
				this.$u.toast('请输入手机验证码');
				return false;
			}
			this.$u.api.saveRegByValidCode({
				loginCode: this.loginCode,
				userName: this.userName,
				regValidCode: this.fpValidCode,
				password: this.password
			})
			.then(res => {
				uni.showModal({
					title: '提示',
					content: res.message,
					showCancel: false,
					success: function () {
						if (res.result == 'true') {
							uni.reLaunch({
								url: '/pages/sys/login/index'
							});
						}
					}
				});
			});
		}
	}
};
</script>
<style lang="scss">
@import 'index.scss';

.agreement {
	display:flex;
	flex-direction:row;
	align-items:center;
	margin-left:90rpx;
	color:#666;
}
</style>
