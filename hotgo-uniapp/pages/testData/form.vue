<template>
	<view class="wrap">
		<u-form class="form" :model="model" :rules="rules" ref="uForm" label-position="left">
			<u-form-item label="编号" prop="id" label-width="180" v-if="model.id">
				<u-input placeholder="请输入编号" v-model="model.id" type="text" maxlength="64"></u-input>
			</u-form-item>
			<u-form-item label="单行文本" prop="testInput" label-width="180">
				<u-input placeholder="请输入单行文本" v-model="model.testInput" type="text" maxlength="200"></u-input>
			</u-form-item>
			<u-form-item label="多行文本" prop="testTextarea" label-width="180" label-position="top">
				<u-input type="textarea" placeholder="请输入多行文本" v-model="model.testTextarea" height="100" maxlength="500" />
			</u-form-item>
			<u-form-item label="下拉框" prop="testSelect" label-width="180">
				<js-select v-model="model.testSelect" dict-type="sys_menu_type" placeholder="请选择选项"></js-select>
			</u-form-item>
			<!-- <u-form-item label="下拉框（树结构）" prop="testSelectMultiple" label-width="260">
				<js-select v-model="model.testSelectMultiple" dict-type="app_tree_test" placeholder="请选择选项" :tree="true"
					:label-value="model.testSelectMultipleLabel" @label-input="model.testSelectMultipleLabel = $event"></js-select>
			</u-form-item> -->
			<u-form-item label="单选框" prop="testRadio" label-width="180">
				<js-radio v-model="model.testRadio" dict-type="sys_menu_type"></js-radio>
			</u-form-item>
			<u-form-item label="复选框" prop="testCheckbox" label-width="180">
				<js-checkbox v-model="model.testCheckbox" dict-type="sys_menu_type"></js-checkbox>
			</u-form-item>
			<u-form-item label="机构选择" prop="testOffice" label-width="180">
				<js-select v-model="model.testOffice.officeCode" :items="officeSelectList" placeholder="请选择机构" :tree="true"
					:label-value="model.testOffice.officeName" @label-input="model.testOffice.officeName = $event"></js-select>
			</u-form-item>
			<u-form-item label="人员选择" prop="testUser" label-width="180">
				<js-select v-model="model.testUser.userCode" :items="userSelectList" placeholder="请选择人员" :tree="true"
					:label-value="model.testUser.userName" @label-input="model.testUser.userName = $event"></js-select>
			</u-form-item>
			<u-form-item label="上传图片（选填）" prop="images" label-position="top">
				<js-uploadfile v-model="model.dataMap" :biz-key="model.id" biz-type="testData_image"></js-uploadfile>
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
				testInput: '',
				testTextarea: '',
				testSelect: '',
				testSelectMultiple: '',
				testSelectMultipleLabel: '',
				testRadio: '',
				testCheckbox: '',
				testUser: {
					userCode: '',
					userName: ''
				},
				testOffice: {
					officeCode: '',
					officeName: ''
				}
			},
			rules: {
				testInput: [
					{
						required: true,
						message: '请输入单行文本',
						trigger: ['change','blur'],
					}
				]
			},
			officeSelectList: [],
			userSelectList: [],
		};
	},
	onLoad(params){
		this.$u.api.testData.form(params).then(res => {
			Object.assign(this.model, res.testData);
		});
	},
	onReady() {
		this.$refs.uForm.setRules(this.rules);
		// 机构数据
		this.$u.api.office.treeData().then(res => {
			this.officeSelectList = res;
		});
		// 人员和机构数据
		this.$u.api.office.treeData({isLoadUser: true}).then(res => {
			this.userSelectList = res;
		});
	},
	methods: {
		submit() {
			//console.log(this.model)
			this.$refs.uForm.validate(valid => {
				if (valid) {
					this.$u.api.testData.save(this.model).then(res => {
						uni.showModal({
							title: '提示',
							content: res.message,
							showCancel: false,
							success: function () {
								if (res.result == 'true') {
									uni.setStorageSync('refreshList', true);
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
