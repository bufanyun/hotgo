<template>
    <view class="wrap">
        <js-lang title="login.title" :showBtn="true"></js-lang>
        <view class="logo">
            <image src="/static/jeesite/logo200.png"></image>
        </view>
        <view class="list">
            <view class="list-call">
                <u-icon class="u-icon" size="40" name="account"></u-icon>
                <input class="u-input" type="text" v-model="username" maxlength="32"
                       :placeholder="$t('login.placeholderAccount')"/>
                <u-checkbox v-model="remember" active-color="#69cbff">{{$t('login.autoLogin')}}</u-checkbox>
            </view>
            <view class="list-call">
                <u-icon class="u-icon" size="40" name="lock"></u-icon>
                <input class="u-input" type="text" v-model="password" maxlength="32"
                       :placeholder="$t('login.placeholderPassword')" :password="!showPassword"/>
                <image class="u-icon-right"
                       :src="'/static/jeesite/login/eye_' + (showPassword ? 'open' : 'close') + '.png'"
                       @click="showPass()"></image>
            </view>
            <view class="list-call" v-if="isValidCodeLogin">
                <u-icon class="u-icon" size="40" name="coupon"></u-icon>
                <input class="u-input" type="text" v-model="validCode" maxlength="4" placeholder="验证码"/>
                <u-image class="img-valid-code" width="300rpx" height="90rpx" :src="imgValidCodeSrc"
                         @click="refreshImgValidCode()"></u-image>
            </view>
        </view>
        <view class="button" hover-class="button-hover" @click="submit()">
            <text>{{$t('login.loginButton')}}</text>
        </view>
        <view class="footer">
            <navigator url="forget" open-type="navigate">{{$t('login.forget')}}</navigator>
            <text>|</text>
            <navigator url="reg" open-type="navigate">{{$t('login.reg')}}</navigator>
        </view>
        <view class="oauth2">
            <u-icon class="u-icon" size="120" color="#00d969" name="weixin-circle-fill" @click="wxLogin"></u-icon>
            <u-icon class="u-icon" size="120" color="#4fa1e8" name="qq-circle-fill" @click="qqLogin"></u-icon>
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
                username: 'admin',
                password: '123456',
                showPassword: false,
                remember: true,
                isValidCodeLogin: false,
                validCode: '',
                imgValidCodeSrc: null,
                baseUrl: ''
            };
        },
        onLoad() {
            this.$u.api.loginCheck().then(res => {
                if (typeof res === 'object' && res.data.result !== 'login') {
                    uni.reLaunch({
                        url: '/pages/sys/home/index'
                    });
                }
            });
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
            submit() {
                if (this.username.length == 0) {
                    this.$u.toast('请输入账号');
                    return;
                }
                if (this.password.length == 0) {
                    this.$u.toast('请输入密码');
                    return;
                }
                this.$u.api.login({
                    username: this.username,
                    password: this.password,
                    cid: 111,
                    code: this.validCode === '' ? '1234' : this.validCode,
                    device: 'api',
                    param_remember: this.remember
                })
                    .then(res => {
                        this.$u.toast(res.message || '未连接服务器');
                        if (res.code === 0) {
                            setTimeout(() => {
                                uni.reLaunch({
                                    url: '/pages/sys/home/index'
                                });
                            }, 500);
                        }
                        if (res.isValidCodeLogin) {
                            this.isValidCodeLogin = true;
                            this.refreshImgValidCode();
                        }
                    });
            },
            wxLogin(res) {
                this.$u.toast('微信登录');
            },
            qqLogin() {
                this.$u.toast('QQ 登录');
            },
            updateBaseUrl() {
                this.vuex_config.baseUrl = this.baseUrl;
                this.$u.vuex('vuex_config', this.vuex_config);
                this.$u.http.setConfig({
                    baseUrl: this.baseUrl
                });
                this.$u.toast('切换成功！');
            }
        }
    };
</script>
<style lang="scss">
    @import 'index.scss';

    .logo {
        width: 260 rpx;
        height: 260 rpx;
        background: rgba(59, 121, 235, 1);
        box-shadow: 0 rpx 5 rpx 20 rpx 5 rpx rgba(45, 127, 235, 0.5);
        border-radius: 50%;
        margin: 70 rpx auto 10 rpx auto;
    }

    .logo image {
        width: 260 rpx;
        height: 260 rpx;
        border-radius: 50%;
    }

    .base-url js-select {
        width: 100%;
    }

    .button {
        margin: 30 rpx auto 0;
    }

    .footer {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        text-align: center;
        color: #46afff;
        height: 40 rpx;
        line-height: 40 rpx;
        font-size: 35 rpx;
        margin-top: 60 rpx;
    }

    .footer text {
        font-size: 30 rpx;
        margin-left: 25 rpx;
        margin-right: 25 rpx;
    }

    .oauth2 {
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        margin: 55 rpx 100 rpx;

        image {
            height: 100 rpx;
            width: 100 rpx;
        }
    }
</style>
