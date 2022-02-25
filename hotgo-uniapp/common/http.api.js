/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
// 此处第二个参数vm，就是我们在页面使用的this，你可以通过vm获取vuex等操作
const install = (Vue, vm) => {

    // 参数配置对象
    const config = vm.vuex_config;

    // 将各个定义的接口名称，统一放进对象挂载到vm.$u.api(因为vm就是this，也即this.$u.api)下
    vm.$u.api = {
        /** 登录 */
        loginCheck: (params = {}) => vm.$u.get(config.adminPath + '/login/check', params),
        login: (params = {}) => vm.$u.post(config.adminPath + '/login/sign', params),
        logout: (params = {}) => vm.$u.get(config.adminPath + '/login/logout', params),

        /** 会员 */
        memberProfile: (params = {}) => vm.$u.get(config.adminPath + '/member/profile', params),

        // 基础服务：登录登出、身份信息、菜单授权、切换系统、字典数据等
        lang: (params = {}) => vm.$u.get('/base/lang', {'l': params.lang}),
        index: (params = {}) => vm.$u.get(config.adminPath + '/index', params),
        authInfo: (params = {}) => vm.$u.get(config.adminPath + '/authInfo', params),
        menuTree: (params = {}) => vm.$u.get(config.adminPath + '/menuTree', params),
        switchSys: (params = {}) => vm.$u.get(config.adminPath + '/switch/' + params.sysCode),
        dictData: (params = {}) => vm.$u.get(config.adminPath + '/dict/attribute', params),

        // 账号服务：验证码接口、忘记密码接口、注册账号接口等
        validCode: (params = {}) => vm.$u.getText('/validCode', params),
        getFpValidCode: (params = {}) => vm.$u.post('/account/getFpValidCode', params),
        savePwdByValidCode: (params = {}) => vm.$u.post('/account/savePwdByValidCode', params),
        getRegValidCode: (params = {}) => vm.$u.post('/account/getRegValidCode', params),
        saveRegByValidCode: (params = {}) => vm.$u.post('/account/saveRegByValidCode', params),

        // APP公共服务
        upgradeCheck: () => vm.$u.post('/app/upgrade/check', {appCode: config.appCode, appVersion: config.appVersion}),
        commentSave: (params = {}) => vm.$u.post('/app/comment/save', params),

        // 个人信息修改
        user: {
            infoSaveBase: (params = {}) => vm.$u.post(config.adminPath + '/sys/user/infoSaveBase', params),
            infoSavePwd: (params = {}) => vm.$u.post(config.adminPath + '/sys/user/infoSavePwd', params),
            infoSavePqa: (params = {}) => vm.$u.post(config.adminPath + '/sys/user/infoSavePqa', params),
        },

        // 员工用户查询
        empUser: {
            listData: (params = {}) => vm.$u.get(config.adminPath + '/sys/empUser/listData', params),
        },

        // 组织机构查询
        office: {
            treeData: (params = {}) => vm.$u.get(config.adminPath + '/sys/office/treeData', params),
        },

        // 增删改查例子
        testData: {
            form: (params = {}) => vm.$u.post(config.adminPath + '/test/testData/form', params),
            list: (params = {}) => vm.$u.get(config.adminPath + '/log/list', params),
            save: (params = {}) => vm.$u.postJson(config.adminPath + '/test/testData/save', params),
            disable: (params = {}) => vm.$u.post(config.adminPath + '/test/testData/disable', params),
            enable: (params = {}) => vm.$u.post(config.adminPath + '/test/testData/enable', params),
            delete: (params = {}) => vm.$u.post(config.adminPath + '/test/testData/delete', params),
        },

    };

}

export default {
    install
}