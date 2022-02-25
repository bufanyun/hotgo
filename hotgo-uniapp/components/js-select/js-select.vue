<template>
	<view class="u-flex-1">
		<u-input :type="disabled?'input':'select'"
			v-model="options.label"
			:disabled="disabled"
			:placeholder="placeholder"
			:select-open="options.open"
			@click="inputClick"
			></u-input>
		<u-select :mode="options.mode"
			v-model="options.open"
			:list="options.items"
			:label-name="options.itemLabel"
			:value-name="options.itemValue"
			:default-value="options.currentIndex"
			@confirm="selectConfirm"
			style="width: 100%"
			></u-select>
	</view>
</template>
<script>
/**
 * 下拉选择组件
 * @property {Object} value 用于双向绑定选择框的值，返回选择框的 Value
 * @property {Boolean} disabled 是否禁用模式，是否只读模式 
 * @property {String} tree 是否为树结构（默认 false）
 * @property {String} placeholder 选择框的占位符，提示文字
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
		labelValue: {
			type: String,
			default: ''
		},
		placeholder: {
			type: String,
			default: '请选择选项'
		},
		disabled: {
			type: Boolean,
			default: false
		},
		tree: {
			type: Boolean,
			default: false
		},
		dictType: {
			type: String,
			default: ''
		},
		multiple: {
			type: Boolean,
			default: false
		},
		items: {
			type: Array,
			default() {
				return [];
			}
		},
		itemLabel: {
			type: String,
			default: ''
		},
		itemValue: {
			type: String,
			default: ''
		},
		returnFullName: {
			type: Boolean,
			default: false
		},
		returnFullNameSplit: {
			type: String,
			default: '/'
		},
	},
	data() {
		return {
			options: {
				value: this.value,
				label: this.labelValue,
				open: false,
				mode: 'single-column',
				items: this.items,
				itemLabel: this.itemLabel || 'name',
				itemValue: this.itemValue || (this.tree ? 'id' : 'value'),
				currentIndex: [],
				indexMap: {}
			}
		};
	},
	watch: {
		value(val, oldVal) {
			this.options.value = val;
		},
		labelValue(val, oldVal) {
			this.options.label = val;
		},
		items(val, oldVal){
			this.setItems(val);
		}
	},
	created() {
		this.loadData();
	},
	methods: {
		loadData() {
			if (this.dictType != ''){
				this.$u.api.dictData({type: this.dictType}).then(res => {
					if (typeof res === 'object' && res.result === 'login'){
						return;
					}
					this.setItems(res.data);
				});
			}else{
				this.setItems(this.items);
			}
		},
		setItems(res){
			if (this.tree){
				this.options.mode = 'mutil-column-auto';
				res = this.convertTree(res);
			}
			this.options.items = res;
			this.selectValue();
		},
		selectValue() {
			// 微信小程序，需要延迟下，否则获取不 value 导致无法回显数据
			this.$nextTick(() => {
				if (!this.options.value) {
					return;
				}
				for (let i in this.options.items){
					let item = this.options.items[i];
					this.options.indexMap[item[this.options.itemValue]] = Number(i);
					if (item[this.options.itemValue] == this.options.value){
						// this.options.value = item[this.options.itemValue];
						this.options.label = item[this.options.itemLabel];
						if (!this.tree){
							this.options.currentIndex = [this.options.indexMap[this.options.value]];
						}
					}
				}
			});
		},
		convertTree(data) {
			let i, l, key = "id", parentKey = "pId", childKey = "children";
			if (Object.prototype.toString.apply(data) === "[object Array]") {
				let treeData = [], map = [];
				for (i=0, l=data.length; i<l; i++) {
					map[data[i][key]] = data[i];
				}
				for (i=0, l=data.length; i<l; i++) {
					if (map[data[i][parentKey]] && data[i][key] != data[i][parentKey]) {
						if (!map[data[i][parentKey]][childKey]){
							map[data[i][parentKey]][childKey] = [];
						}
						map[data[i][parentKey]][childKey].push(data[i]);
					} else {
						treeData.push(data[i]);
					}
				}
				return treeData;
			}else {
				return [data];
			}
		},
		inputClick() {
			if (!this.disabled){
				this.options.open = true;
			}
		},
		selectConfirm(items) {
			let values = [], labels = [], currentIndexes = [];
			for (let i in items){
				let item = items[i];
				values.push(String(item.value).replace(/^u_/g,''));
				labels.push(String(item.label));
				if (!this.tree){
					currentIndexes.push(this.options.indexMap[item.value])
				}
			}
			this.options.value = values.length > 0 ? values[values.length-1] : '';
			if (this.returnFullName){
				this.options.label = labels.join(this.returnFullNameSplit);
			}else{
				this.options.label = labels.length > 0 ? labels[labels.length-1] : '';
			}
			if (!this.tree){
				this.options.currentIndex = currentIndexes;
			}
			//console.log(this.options.value, this.options.label)
			this.$emit('input', this.options.value);
			this.$emit('label-input', this.options.label);
			this.$emit('confirm', this.options.value, this.options.label);
		}
	}
}
</script>
<style lang="scss" scoped>
	
</style>
