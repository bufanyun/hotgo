<template>
	<view class="u-flex-1">
		<u-checkbox-group :disabled="disabled" @change="change">
			<u-checkbox v-model="item.checked" v-for="(item, index) in options.items" :key="index"
				:name="item.value">{{item.label}}</u-checkbox>
		</u-checkbox-group>
	</view>
</template>
<script>
/**
 * 复选框组件
 * @property {Object} value 用于双向绑定选择框的值，返回选择框的 Value
 * @property {Boolean} disabled 是否禁用模式，是否只读模式 
 * @property {String} dictType 字典类型，从字典里获取，自动设置 items、itemLabel、itemValue
 * @property {String} items 列表数据，可接受对象集合，如：[{name: '是', value: '否'}]
 * @property {String} itemLabel 指定列表数据中的什么属性名作为 option 的标签名，如 name
 * @property {String} itemValue 指定列表数据中的什么属性名作为 option 的 value 值，如 value
 * @example <js-select v-model="model.type" dict-type="sys_yes_no"></js-select>
 * @description Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 * @author ThinkGem
 * @version 2021-3-11
 */
export default {
	props: {
		value: {
			type: String,
			default: ''
		},
		disabled: {
			type: Boolean,
			default: false
		},
		dictType: {
			type: String,
			default: ''
		},
		items: {
			type: Array,
			default() {
				return [];
			}
		},
		itemLabel: {
			type: String,
			default: 'name'
		},
		itemValue: {
			type: String,
			default: 'value'
		}
	},
	data() {
		return {
			options: {
				value: (this.value && this.value.split(',')) || [],
				items: this.items
			}
		};
	},
	watch: {
		value(val, oldVal) {
			this.options.value = (val && val.split(',')) || [];
		},
		items(val, oldVal){
			this.options.items = val;
		}
	},
	created() {
		this.loadData();
	},
	methods: {
		loadData(){
			if (this.dictType != ''){
				this.$u.api.dictData({type: this.dictType}).then(res => {
					// this.options.items = res;
					this.selectValue(res.data);
				});
			}else{
				// this.options.items = this.items;
				this.selectValue(this.items);
			}
		},
		selectValue(items){
			// 微信小程序，需要延迟下，否则获取不 value 导致无法回显数据
			this.$nextTick(() => {
				let vals = this.options.value;
				for (let i in items){
					let item = items[i];
					item.checked = vals.includes(item[this.itemValue]);
				}
				this.options.items = items;
			});
		},
		change(val, name){
			this.$emit('input', (val && val.join(',')) || '');
		}
	}
}
</script>
<style lang="scss" scoped>
	
</style>
